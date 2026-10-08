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

import com.chinmaysawant.lgremote.mobile.Mobile;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

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
    private final ExecutorService exec = Executors.newSingleThreadExecutor();
    private BluetoothHidDevice hid;
    private BluetoothDevice device;
    private boolean registered;
    private boolean proxyAsked;
    private boolean asked;

    Hid(Activity activity) {
        this.activity = activity;
    }

    void command(String cmd) {
        if ("start".equals(cmd)) {
            ensure();
            return;
        }

        exec.execute(() -> send(cmd));
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
            Mobile.setBluetooth("Allow Bluetooth for this app, then tap Connect.");
            return;
        }

        if (hid != null) {
            if (!registered) {
                register();
            }
            activity.runOnUiThread(this::askDiscoverable);
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
                hid = (BluetoothHidDevice) proxy;
                register();
            }

            @Override
            public void onServiceDisconnected(int profile) {
                hid = null;
                registered = false;
                proxyAsked = false;
            }
        }, BluetoothProfile.HID_DEVICE);

        if (!ok) {
            proxyAsked = false;
            Mobile.setBluetooth("Bluetooth HID is unavailable.");
        }
    }

    private void register() {
        if (hid == null || registered) {
            return;
        }

        BluetoothHidDeviceAppSdpSettings sdp = new BluetoothHidDeviceAppSdpSettings(
            "ownframe remote",
            "LG TV remote",
            "ownframe",
            BluetoothHidDevice.SUBCLASS1_COMBO,
            DESCRIPTOR);
        boolean ok = hid.registerApp(sdp, null, null, exec, new BluetoothHidDevice.Callback() {
            @Override
            public void onAppStatusChanged(BluetoothDevice plugged, boolean ready) {
                registered = ready;
                if (!ready) {
                    return;
                }

                Mobile.setBluetooth("Bluetooth waiting for the TV");
                activity.runOnUiThread(() -> askDiscoverable());
                connectKnown();
            }

            @Override
            public void onConnectionStateChanged(BluetoothDevice dev, int state) {
                if (state == BluetoothProfile.STATE_CONNECTED) {
                    device = dev;
                    Mobile.setBluetooth("Bluetooth connected");
                    return;
                }

                if (state == BluetoothProfile.STATE_DISCONNECTED
                    && device != null && dev.equals(device)) {
                    device = null;
                    Mobile.setBluetooth("Bluetooth waiting for the TV");
                }
            }
        });

        if (!ok) {
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

    private void connectKnown() {
        BluetoothAdapter adapter = BluetoothAdapter.getDefaultAdapter();
        if (adapter == null || hid == null) {
            return;
        }

        for (BluetoothDevice dev : adapter.getBondedDevices()) {
            String name = dev.getName();
            if (name == null) {
                continue;
            }

            String low = name.toLowerCase();
            if (low.contains("lg") || low.contains("webos") || low.contains("[tv]")) {
                hid.connect(dev);
            }
        }
    }

    private void send(String cmd) {
        if (hid == null || device == null) {
            Mobile.setBluetooth("Pair this phone in the TV Bluetooth list.");
            return;
        }

        try {
            if (cmd.startsWith("key:")) {
                int code = Integer.parseInt(cmd.substring(4));
                hid.sendReport(device, 1, new byte[] {0, 0, (byte) code, 0, 0, 0, 0, 0});
                Thread.sleep(40);
                hid.sendReport(device, 1, new byte[8]);
                return;
            }

            if (cmd.startsWith("con:")) {
                int code = Integer.parseInt(cmd.substring(4));
                hid.sendReport(device, 2, new byte[] {
                    (byte) (code & 0xff), (byte) ((code >> 8) & 0xff)
                });
                Thread.sleep(40);
                hid.sendReport(device, 2, new byte[] {0, 0});
            }
        } catch (InterruptedException ex) {
            Thread.currentThread().interrupt();
        }
    }
}
