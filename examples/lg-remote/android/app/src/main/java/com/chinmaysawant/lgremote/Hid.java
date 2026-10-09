package com.chinmaysawant.lgremote;

import android.app.Activity;
import android.bluetooth.BluetoothAdapter;
import android.bluetooth.BluetoothDevice;
import android.bluetooth.BluetoothHidDevice;
import android.bluetooth.BluetoothHidDeviceAppSdpSettings;
import android.bluetooth.BluetoothProfile;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.os.Build;
import android.os.Handler;
import android.os.Looper;
import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.IntentFilter;
import java.util.List;
import java.util.Locale;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.RejectedExecutionException;

import com.chinmaysawant.lgremote.mobile.Mobile;

import java.util.concurrent.ExecutorService;

// Hid turns the phone into a Bluetooth keyboard and consumer remote.
// The TV pairs with the phone from its Bluetooth device list.
public final class Hid {
    private static final byte[] DESCRIPTOR = new byte[] {
        0x05, 0x01, 0x09, 0x06, (byte) 0xa1, 0x01,
        (byte) 0x85, 0x01,
        0x05, 0x07, 0x19, (byte) 0xe0, 0x29, (byte) 0xe7,
        0x15, 0x00, 0x25, 0x01, 0x75, 0x01, (byte) 0x95, 0x08,
        (byte) 0x81, 0x02,
        (byte) 0x95, 0x01, 0x75, 0x08, (byte) 0x81, 0x01,
        (byte) 0x95, 0x05, 0x75, 0x01, 0x05, 0x08, 0x19, 0x01, 0x29, 0x05,
        (byte) 0x91, 0x02,
        (byte) 0x95, 0x01, 0x75, 0x03, (byte) 0x91, 0x01,
        (byte) 0x95, 0x06, 0x75, 0x08, 0x15, 0x00, 0x25, 0x65,
        0x05, 0x07, 0x19, 0x00, 0x29, 0x65, (byte) 0x81, 0x00,
        (byte) 0xc0,
        0x05, 0x0c, 0x09, 0x01, (byte) 0xa1, 0x01,
        (byte) 0x85, 0x02,
        0x15, 0x00, 0x26, (byte) 0xff, 0x03,
        0x19, 0x00, 0x2a, (byte) 0xff, 0x03,
        0x75, 0x10, (byte) 0x95, 0x01, (byte) 0x81, 0x00,
        (byte) 0xc0
    };

    private final Activity activity;
    private final ExecutorService exec = new ThreadPoolExecutor(1, 1, 0,
        TimeUnit.MILLISECONDS, new ArrayBlockingQueue<Runnable>(32));
    private final Handler reconnect = new Handler(Looper.getMainLooper());
    private volatile boolean wanted, paused, closed, registering, connecting;
    private int retries;
    private final Runnable retry = new Runnable() {
        @Override public void run() {
            if (closed || paused || !wanted || device != null) { return; }
            connecting = false;
            try { ensure(); if (registered) { connectKnown(); } }
            catch (SecurityException ex) { permissionLost(); return; }
            if (device == null && ++retries < 8) {
                reconnect.postDelayed(this, Math.min(5000, 500L << Math.min(retries, 3)));
            } else if (device == null) {
                Mobile.setBluetooth("TV unavailable. Tap Connect to retry Bluetooth.");
            }
        }
    };
    private final BroadcastReceiver radio = new BroadcastReceiver() {
        @Override public void onReceive(Context context, Intent intent) {
            int state = intent.getIntExtra(BluetoothAdapter.EXTRA_STATE, -1);
            if (state == BluetoothAdapter.STATE_ON && wanted && !paused) { resume(); }
            if (state == BluetoothAdapter.STATE_OFF) {
                device = null; registered = false; registering = false; connecting = false;
                reconnect.removeCallbacks(retry);
                if (wanted) { Mobile.setBluetooth("Bluetooth is off. Turn it on or use Wi-Fi."); }
            }
        }
    };
    private volatile BluetoothHidDevice hid;
    private volatile BluetoothDevice device;
    private volatile boolean registered;
    private volatile boolean proxyAsked;
    private boolean asked;

    Hid(Activity activity) {
        this.activity = activity;
        IntentFilter filter = new IntentFilter(BluetoothAdapter.ACTION_STATE_CHANGED);
        if (Build.VERSION.SDK_INT >= 33) {
            activity.registerReceiver(radio, filter, Context.RECEIVER_EXPORTED);
        } else { activity.registerReceiver(radio, filter); }
    }

    void command(String cmd) {
        if (closed) { return; }
        if ("stop".equals(cmd)) { wanted = false; reconnect.removeCallbacks(retry); return; }
        if ("start".equals(cmd)) {
            wanted = true;
            retries = 0;
            try { ensure(); if (registered) { connectKnown(); } scheduleReconnect(); }
            catch (SecurityException ex) {
                Mobile.setBluetooth("Bluetooth permission was revoked. Use Wi-Fi or allow Bluetooth.");
            }
            return;
        }

        try { exec.execute(() -> send(cmd)); }
        catch (RejectedExecutionException ex) { Mobile.setBluetooth("Bluetooth is busy. Try again."); }
    }

    private boolean allowed() {
        if (Build.VERSION.SDK_INT < 31) {
            return true;
        }

        return activity.checkSelfPermission(android.Manifest.permission.BLUETOOTH_CONNECT)
                == PackageManager.PERMISSION_GRANTED
            && activity.checkSelfPermission(android.Manifest.permission.BLUETOOTH_ADVERTISE)
                == PackageManager.PERMISSION_GRANTED;
    }

    private void ensure() {
        if (!allowed()) {
            activity.requestPermissions(new String[] {
                android.Manifest.permission.BLUETOOTH_CONNECT,
                android.Manifest.permission.BLUETOOTH_ADVERTISE
            }, 1);
            Mobile.setBluetooth("Allow Bluetooth to pair this phone. Wi-Fi still works.");
            return;
        }

        BluetoothAdapter radioAdapter = BluetoothAdapter.getDefaultAdapter();
        if (radioAdapter != null && !radioAdapter.isEnabled()) {
            activity.startActivity(new Intent(BluetoothAdapter.ACTION_REQUEST_ENABLE));
            Mobile.setBluetooth("Turn on Bluetooth to connect. Wi-Fi still works.");
            return;
        }
        if (hid != null) {
            if (!registered) {
                register();
            }
            return;
        }

        if (proxyAsked) {
            return;
        }

        BluetoothAdapter adapter = BluetoothAdapter.getDefaultAdapter();
        if (adapter == null) {
            Mobile.setBluetooth("This phone has no Bluetooth radio.");
            return;
        }

        proxyAsked = true;
        boolean ok = adapter.getProfileProxy(activity, new BluetoothProfile.ServiceListener() {
            @Override
            public void onServiceConnected(int profile, BluetoothProfile proxy) {
                if (closed) { BluetoothAdapter.getDefaultAdapter().closeProfileProxy(profile, proxy); return; }
                hid = (BluetoothHidDevice) proxy;
                register();
            }

            @Override
            public void onServiceDisconnected(int profile) {
                hid = null;
                device = null;
                registered = false;
                registering = false;
                proxyAsked = false;
                scheduleReconnect();
            }
        }, BluetoothProfile.HID_DEVICE);

        if (!ok) {
            proxyAsked = false;
            Mobile.setBluetooth("Bluetooth HID is unavailable.");
        }
    }

    private void register() {
        BluetoothHidDevice sender = hid;
        if (sender == null || registered || registering || closed || paused || !wanted) {
            return;
        }

        BluetoothHidDeviceAppSdpSettings sdp = new BluetoothHidDeviceAppSdpSettings(
            "ownframe remote",
            "LG TV remote",
            "ownframe",
            BluetoothHidDevice.SUBCLASS1_COMBO,
            DESCRIPTOR);
        registering = true;
        boolean ok = sender.registerApp(sdp, null, null, exec, new BluetoothHidDevice.Callback() {
            @Override
            public void onAppStatusChanged(BluetoothDevice plugged, boolean ready) {
                if (closed) { return; }
                registered = ready;
                registering = false;
                if (!ready) {
                    device = null;
                    Mobile.setBluetooth("Bluetooth paused. Tap Connect to resume.");
                    scheduleReconnect();
                    return;
                }

                Mobile.setBluetooth("Bluetooth waiting for the TV");
                connectKnown();
                if (device == null) { scheduleReconnect(); }
            }

            @Override
            public void onConnectionStateChanged(BluetoothDevice dev, int state) {
                if (state == BluetoothProfile.STATE_CONNECTED) {
                    connected(dev);
                    return;
                }

                if (closed) { return; }
                if (state == BluetoothProfile.STATE_DISCONNECTED && (device == null || dev.equals(device))) {
                    device = null;
                    connecting = false;
                    Mobile.setBluetooth("Bluetooth disconnected. Reconnecting to the TV.");
                    scheduleReconnect();
                }
            }
        });

        if (!ok) {
            registering = false;
            Mobile.setBluetooth("Bluetooth HID did not register.");
        }
    }

    private void askDiscoverable() {
        if (asked) {
            return;
        }

        asked = true;
        Intent intent = new Intent(BluetoothAdapter.ACTION_REQUEST_DISCOVERABLE);
        intent.putExtra(BluetoothAdapter.EXTRA_DISCOVERABLE_DURATION, 300);
        activity.startActivity(intent);
    }

    private void connected(BluetoothDevice dev) {
        if (closed) { return; }
        device = dev;
        connecting = false;
        reconnect.removeCallbacks(retry);
        activity.runOnUiThread(() -> retries = 0);
        activity.getPreferences(Context.MODE_PRIVATE).edit().putString("tv-address", dev.getAddress()).apply();
        Mobile.setBluetooth("Bluetooth connected");
    }

    private void connectKnown() {
        try { connectKnownAllowed(); }
        catch (SecurityException ex) { permissionLost(); }
    }

    private void connectKnownAllowed() {
        BluetoothAdapter adapter = BluetoothAdapter.getDefaultAdapter();
        BluetoothHidDevice sender = hid;
        if (adapter == null || sender == null || !registered || connecting || paused || !wanted || closed) { return; }
        List<BluetoothDevice> live = sender.getConnectedDevices();
        if (!live.isEmpty()) { connected(live.get(0)); return; }
        String saved = activity.getPreferences(Context.MODE_PRIVATE).getString("tv-address", "");
        BluetoothDevice target = null;
        for (BluetoothDevice dev : adapter.getBondedDevices()) {
            if (saved.equals(dev.getAddress())) { target = dev; break; }
            String name = dev.getName();
            if (saved.isEmpty() && target == null && name != null) {
                String low = name.toLowerCase(Locale.ROOT);
                if (low.contains("lg") || low.contains("webos") || low.contains("[tv]")) { target = dev; }
            }
        }
        if (target != null) {
            connecting = sender.connect(target);
            Mobile.setBluetooth(connecting ? "Reconnecting to the paired TV" : "Bluetooth reconnect failed. Retrying.");
        } else { activity.runOnUiThread(this::askDiscoverable); }
    }

    private void scheduleReconnect() {
        activity.runOnUiThread(() -> {
            reconnect.removeCallbacks(retry);
            BluetoothAdapter adapter = BluetoothAdapter.getDefaultAdapter();
            if (!closed && !paused && wanted && device == null && allowed() && adapter != null && adapter.isEnabled()) {
                reconnect.postDelayed(retry, 500);
            }
        });
    }

    private void permissionLost() {
        Mobile.setBluetooth("Bluetooth permission was revoked. Use Wi-Fi or allow Bluetooth.");
    }

    void resume() {
        paused = false;
        if (!wanted || closed) { return; }
        retries = 0;
        try { ensure(); if (registered) { connectKnown(); } scheduleReconnect(); }
        catch (SecurityException ex) { permissionLost(); }
    }

    void pause() { paused = true; reconnect.removeCallbacks(retry); }

    void close() {
        closed = true;
        reconnect.removeCallbacksAndMessages(null);
        activity.unregisterReceiver(radio);
        exec.shutdown();
        BluetoothHidDevice sender = hid;
        hid = null; device = null;
        if (sender == null || !allowed()) { return; }
        try {
            sender.unregisterApp();
            BluetoothAdapter adapter = BluetoothAdapter.getDefaultAdapter();
            if (adapter != null) { adapter.closeProfileProxy(BluetoothProfile.HID_DEVICE, sender); }
        } catch (SecurityException ex) { permissionLost(); }
    }

    private void reportResult(boolean sent) {
        Mobile.setBluetooth(sent ? "Bluetooth report sent" : "Bluetooth send failed. Tap Connect to retry.");
        if (!sent) { device = null; scheduleReconnect(); }
    }

    private void send(String cmd) {
        BluetoothHidDevice sender = hid;
        BluetoothDevice target = device;
        if (!wanted || paused || closed) { return; }
        if (sender == null || target == null) {
            Mobile.setBluetooth("TV is disconnected. Reconnecting; tap again once connected.");
            scheduleReconnect();
            return;
        }

        try {
            if (cmd.startsWith("key:")) {
                int code = Integer.parseInt(cmd.substring(4));
                boolean sent = sender.sendReport(target, 1,
                    new byte[] {0, 0, (byte) code, 0, 0, 0, 0, 0});
                Thread.sleep(12);
                sent &= sender.sendReport(target, 1, new byte[8]);
                reportResult(sent);
                return;
            }

            if (cmd.startsWith("con:")) {
                int code = Integer.parseInt(cmd.substring(4));
                boolean sent = sender.sendReport(target, 2, new byte[] {
                    (byte) (code & 0xff), (byte) ((code >> 8) & 0xff)
                });
                Thread.sleep(12);
                sent &= sender.sendReport(target, 2, new byte[] {0, 0});
                reportResult(sent);
            }
        } catch (SecurityException ex) {
            Mobile.setBluetooth("Bluetooth permission was revoked. Use Wi-Fi or allow Bluetooth.");
        } catch (InterruptedException ex) {
            Thread.currentThread().interrupt();
        }
    }
}
