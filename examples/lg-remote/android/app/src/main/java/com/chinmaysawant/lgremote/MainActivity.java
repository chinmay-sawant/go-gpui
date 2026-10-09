package com.chinmaysawant.lgremote;

import android.Manifest;
import android.app.Activity;
import android.content.pm.PackageManager;
import android.net.wifi.WifiManager;
import android.os.Build;
import android.graphics.Insets;
import android.view.View;
import android.view.WindowInsets;
import android.view.WindowManager;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import com.chinmaysawant.lgremote.mobile.EbitenView;
import com.chinmaysawant.lgremote.mobile.Mobile;
import com.chinmaysawant.lgremote.mobile.CommandListener;
import java.util.concurrent.atomic.AtomicBoolean;

// MainActivity fills the screen with the remote and polls Bluetooth commands.
public class MainActivity extends Activity {
    private static final int MATCH = ViewGroup.LayoutParams.MATCH_PARENT;

    private EbitenView view;
    private Hid hid;
    private RemoteAccessibility accessibility;
    private WifiManager.MulticastLock lock;
    private final Handler poller = new Handler(Looper.getMainLooper());
    private final AtomicBoolean dispatchQueued = new AtomicBoolean();
    private final CommandListener listener = new CommandListener() {
        @Override public void ready() {
            if (!dispatchQueued.compareAndSet(false, true)) { return; }
            poller.post(() -> {
                dispatchQueued.set(false);
                String cmd;
                while (hid != null && !(cmd = Mobile.takeCommand()).isEmpty()) { hid.command(cmd); }
            });
        }
    };
    private final Runnable poll = new Runnable() {
        @Override
        public void run() {
            if (accessibility != null) { accessibility.refresh(); }
            poller.postDelayed(this, 40);
        }
    };

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        if (Build.VERSION.SDK_INT >= 30) {
            getWindow().setDecorFitsSystemWindows(false);
        } else {
            getWindow().getDecorView().setSystemUiVisibility(
                View.SYSTEM_UI_FLAG_LAYOUT_STABLE
                    | View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN
                    | View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION);
        }
        getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_ADJUST_RESIZE);
        holdMulticast();

        // Load the Go runtime and set the key directory before the game loop
        // starts. The first frame reads that directory.
        Mobile.setStoreDir(getFilesDir().getAbsolutePath());

        FrameLayout layout = new FrameLayout(this);
        layout.setLayoutParams(new ViewGroup.LayoutParams(MATCH, MATCH));
        view = new EbitenView(this);
        layout.addView(view, new FrameLayout.LayoutParams(MATCH, MATCH));
        view.setImportantForAccessibility(View.IMPORTANT_FOR_ACCESSIBILITY_NO_HIDE_DESCENDANTS);
        accessibility = new RemoteAccessibility(this);
        layout.addView(accessibility, new FrameLayout.LayoutParams(MATCH, MATCH));
        layout.setOnApplyWindowInsetsListener((v, insets) -> {
            if (Build.VERSION.SDK_INT >= 30) {
                Insets safe = insets.getInsets(WindowInsets.Type.systemBars()
                    | WindowInsets.Type.displayCutout() | WindowInsets.Type.ime());
                layout.setPadding(safe.left, safe.top, safe.right, safe.bottom);
            } else {
                int left = insets.getSystemWindowInsetLeft();
                int top = insets.getSystemWindowInsetTop();
                int right = insets.getSystemWindowInsetRight();
                int bottom = insets.getSystemWindowInsetBottom();
                if (insets.getDisplayCutout() != null) {
                    left = Math.max(left, insets.getDisplayCutout().getSafeInsetLeft());
                    top = Math.max(top, insets.getDisplayCutout().getSafeInsetTop());
                    right = Math.max(right, insets.getDisplayCutout().getSafeInsetRight());
                    bottom = Math.max(bottom, insets.getDisplayCutout().getSafeInsetBottom());
                }
                layout.setPadding(left, top, right, bottom);
            }
            if (Build.VERSION.SDK_INT >= 30) {
                return WindowInsets.CONSUMED;
            }
            return insets.consumeSystemWindowInsets().consumeDisplayCutout();
        });
        setContentView(layout);
        if (Build.VERSION.SDK_INT >= 33) {
            getOnBackInvokedDispatcher().registerOnBackInvokedCallback(
                android.window.OnBackInvokedDispatcher.PRIORITY_DEFAULT, this::onBackPressed);
        }
        layout.requestApplyInsets();
        updateTextSize();
        hid = new Hid(this);
    }

    @Override
    protected void onResume() {
        super.onResume();
        updateTextSize();
        if (hid != null) { hid.resume(); }
        Mobile.setCommandListener(listener);
        view.resumeGame();
        poller.removeCallbacks(poll);
        poller.post(poll);
    }

    @Override
    protected void onPause() {
        if (accessibility != null) { accessibility.cancelTouch(); }
        Mobile.setCommandListener(null);
        if (hid != null) { hid.pause(); }
        poller.removeCallbacks(poll);
        view.suspendGame();
        super.onPause();
    }

    @Override
    protected void onDestroy() {
        if (lock != null && lock.isHeld()) {
            lock.release();
        }
        Mobile.setCommandListener(null);
        poller.removeCallbacksAndMessages(null);
        if (hid != null) { hid.close(); }
        super.onDestroy();
    }

    @Override
    public void onBackPressed() {
        WindowInsets insets = getWindow().getDecorView().getRootWindowInsets();
        if (Build.VERSION.SDK_INT >= 30 && insets != null
            && insets.isVisible(WindowInsets.Type.ime())) {
            Mobile.queueAction("blur:");
            ((android.view.inputmethod.InputMethodManager) getSystemService(INPUT_METHOD_SERVICE))
                .hideSoftInputFromWindow(view.getWindowToken(), 0);
            return;
        }
        try {
            org.json.JSONObject state = new org.json.JSONObject(Mobile.accessibility());
            if (!"remote".equals(state.optString("Panel", "remote"))) {
                Mobile.queueAction("panel:remote");
                return;
            }
        } catch (org.json.JSONException ex) { }
        finish();
    }

    private void updateTextSize() {
        float scale = getResources().getConfiguration().fontScale;
        Mobile.setFontSize(Math.round(16 * scale));
    }

    @Override
    public void onRequestPermissionsResult(int requestCode, String[] permissions, int[] results) {
        super.onRequestPermissionsResult(requestCode, permissions, results);
        if (requestCode != 1) { return; }
        boolean allowed = results.length == 2;
        for (int result : results) { allowed &= result == PackageManager.PERMISSION_GRANTED; }
        if (allowed && hid != null) {
            hid.command("start");
        } else {
            Mobile.setBluetooth("Bluetooth permission denied. Wi-Fi controls still work.");
        }
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
