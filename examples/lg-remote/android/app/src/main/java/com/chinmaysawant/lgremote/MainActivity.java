package com.chinmaysawant.lgremote;

import android.Manifest;
import android.app.Activity;
import android.content.pm.PackageManager;
import android.net.wifi.WifiManager;
import android.os.Build;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import com.chinmaysawant.lgremote.mobile.EbitenView;
import com.chinmaysawant.lgremote.mobile.Mobile;

// MainActivity fills the screen with the remote and polls Bluetooth commands.
public class MainActivity extends Activity {
    private static final int MATCH = ViewGroup.LayoutParams.MATCH_PARENT;

    private EbitenView view;
    private Hid hid;
    private WifiManager.MulticastLock lock;
    private final Handler poller = new Handler(Looper.getMainLooper());
    private final Runnable poll = new Runnable() {
        @Override
        public void run() {
            String cmd = Mobile.takeCommand();
            if (cmd != null && cmd.length() > 0 && hid != null) {
                hid.command(cmd);
            }
            poller.postDelayed(this, 40);
        }
    };

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        askRadio();
        holdMulticast();

        // Load the Go runtime and set the key directory before the game loop
        // starts. The first frame reads that directory.
        Mobile.setStoreDir(getFilesDir().getAbsolutePath());

        FrameLayout layout = new FrameLayout(this);
        layout.setLayoutParams(new ViewGroup.LayoutParams(MATCH, MATCH));
        view = new EbitenView(this);
        layout.addView(view, new FrameLayout.LayoutParams(MATCH, MATCH));
        setContentView(layout);
        hid = new Hid(this);
    }

    @Override
    protected void onResume() {
        super.onResume();
        view.resumeGame();
        poller.removeCallbacks(poll);
        poller.post(poll);
    }

    @Override
    protected void onPause() {
        poller.removeCallbacks(poll);
        view.suspendGame();
        super.onPause();
    }

    @Override
    protected void onDestroy() {
        if (lock != null && lock.isHeld()) {
            lock.release();
        }
        super.onDestroy();
    }

    private void askRadio() {
        if (Build.VERSION.SDK_INT < 31) {
            return;
        }

        if (checkSelfPermission(Manifest.permission.BLUETOOTH_CONNECT)
                == PackageManager.PERMISSION_GRANTED
            && checkSelfPermission(Manifest.permission.BLUETOOTH_ADVERTISE)
                == PackageManager.PERMISSION_GRANTED) {
            return;
        }

        requestPermissions(new String[] {
            Manifest.permission.BLUETOOTH_CONNECT,
            Manifest.permission.BLUETOOTH_ADVERTISE
        }, 1);
    }

    private void holdMulticast() {
        WifiManager wifi = (WifiManager) getApplicationContext()
            .getSystemService(WIFI_SERVICE);
        if (wifi == null) {
            return;
        }

        lock = wifi.createMulticastLock("lg-remote");
        lock.setReferenceCounted(false);
        lock.acquire();
    }
}
