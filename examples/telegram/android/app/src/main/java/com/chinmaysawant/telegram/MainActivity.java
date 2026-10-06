package com.chinmaysawant.telegram;

import android.Manifest;
import android.app.Activity;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.graphics.Bitmap;
import android.graphics.BitmapFactory;
import android.graphics.Color;
import android.graphics.Insets;
import android.net.Uri;
import android.os.Build;
import android.os.Bundle;
import android.os.Handler;
import android.os.Looper;
import android.provider.MediaStore;
import android.view.KeyEvent;
import android.view.View;
import android.view.ViewGroup;
import android.view.WindowInsets;
import android.view.WindowInsetsController;
import android.widget.FrameLayout;

import com.chinmaysawant.telegram.mobile.EbitenView;
import com.chinmaysawant.telegram.mobile.Mobile;

import java.io.ByteArrayOutputStream;

// MainActivity fills the screen with the bound go-gpui page. It forwards
// the back key, a camera capture, and a gallery pick to the page, and
// reports the system bar insets so the page keeps clear of them.
public class MainActivity extends Activity {
    private static final int REQ_PHOTO = 1;
    private static final int REQ_PERMISSION = 2;
    private static final int MATCH = ViewGroup.LayoutParams.MATCH_PARENT;

    private EbitenView view;
    private boolean waitingPermission;

    private final Handler poller = new Handler(Looper.getMainLooper());
    private final Runnable poll = new Runnable() {
        @Override
        public void run() {
            String attach = Mobile.attachRequest();
            if ("camera".equals(attach)) {
                openCamera();
            } else if ("gallery".equals(attach)) {
                openGallery();
            }
            setBarIcons(Mobile.darkTheme());
            poller.postDelayed(this, 250);
        }
    };

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        edgeToEdge();
        setContentView(buildLayout());
    }

    // buildLayout holds the Ebiten view and listens for window insets.
    private FrameLayout buildLayout() {
        FrameLayout layout = new FrameLayout(this);
        layout.setLayoutParams(new ViewGroup.LayoutParams(MATCH, MATCH));

        view = new EbitenView(this);
        layout.addView(view, new FrameLayout.LayoutParams(MATCH, MATCH));

        layout.setOnApplyWindowInsetsListener((v, insets) -> {
            int top, bottom;
            if (Build.VERSION.SDK_INT >= 30) {
                // The IME inset lifts the composer above the keyboard, so
                // Ebiten does not pan the canvas to reveal the caret.
                Insets bars = insets.getInsets(
                    WindowInsets.Type.systemBars()
                        | WindowInsets.Type.displayCutout()
                        | WindowInsets.Type.ime());
                top = bars.top;
                bottom = bars.bottom;
            } else {
                top = insets.getSystemWindowInsetTop();
                bottom = insets.getSystemWindowInsetBottom();
            }
            float density = getResources().getDisplayMetrics().density;
            Mobile.setInsets(Math.round(top / density), Math.round(bottom / density));
            return insets;
        });

        return layout;
    }

    // edgeToEdge draws the page under the system bars; the page pads itself.
    private void edgeToEdge() {
        if (Build.VERSION.SDK_INT >= 30) {
            getWindow().setDecorFitsSystemWindows(false);
        } else {
            getWindow().getDecorView().setSystemUiVisibility(
                View.SYSTEM_UI_FLAG_LAYOUT_STABLE
                    | View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN
                    | View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION);
        }
        getWindow().setStatusBarColor(Color.TRANSPARENT);
        getWindow().setNavigationBarColor(Color.TRANSPARENT);
    }

    // setBarIcons keeps the status bar icons readable in both themes.
    private void setBarIcons(boolean dark) {
        if (Build.VERSION.SDK_INT < 30) {
            return;
        }
        WindowInsetsController c = getWindow().getInsetsController();
        if (c == null) {
            return;
        }
        int mask = WindowInsetsController.APPEARANCE_LIGHT_STATUS_BARS
            | WindowInsetsController.APPEARANCE_LIGHT_NAVIGATION_BARS;
        c.setSystemBarsAppearance(dark ? 0 : mask, mask);
    }

    // keyboardVisible reports whether the soft keyboard is up, in which
    // case back belongs to the keyboard.
    private boolean keyboardVisible() {
        if (Build.VERSION.SDK_INT < 30) {
            return false;
        }
        WindowInsets insets = getWindow().getDecorView().getRootWindowInsets();
        return insets != null && insets.isVisible(WindowInsets.Type.ime());
    }

    // dispatchKeyEvent sends the system back press to the page first, and
    // closes the activity when the page has nowhere to go back to.
    @Override
    public boolean dispatchKeyEvent(KeyEvent event) {
        if (event.getKeyCode() != KeyEvent.KEYCODE_BACK) {
            return super.dispatchKeyEvent(event);
        }
        if (keyboardVisible()) {
            return super.dispatchKeyEvent(event);
        }
        if (event.getAction() == KeyEvent.ACTION_UP && !Mobile.back()) {
            finish();
        }
        return true;
    }

    // openCamera asks for the camera permission once, then captures.
    private void openCamera() {
        if (Build.VERSION.SDK_INT >= 23
            && checkSelfPermission(Manifest.permission.CAMERA) != PackageManager.PERMISSION_GRANTED) {
            waitingPermission = true;
            requestPermissions(new String[]{Manifest.permission.CAMERA}, REQ_PERMISSION);
            return;
        }
        Intent intent = new Intent(MediaStore.ACTION_IMAGE_CAPTURE);
        if (intent.resolveActivity(getPackageManager()) != null) {
            startActivityForResult(intent, REQ_PHOTO);
        }
    }

    // openGallery opens the system photo picker on Android 13 and newer, and
    // the document picker on older releases.
    private void openGallery() {
        Intent intent;
        if (Build.VERSION.SDK_INT >= 33) {
            intent = new Intent(MediaStore.ACTION_PICK_IMAGES);
        } else {
            intent = new Intent(Intent.ACTION_GET_CONTENT);
            intent.addCategory(Intent.CATEGORY_OPENABLE);
        }
        intent.setType("image/*");
        if (intent.resolveActivity(getPackageManager()) == null) {
            intent = new Intent(Intent.ACTION_GET_CONTENT);
            intent.addCategory(Intent.CATEGORY_OPENABLE);
            intent.setType("image/*");
        }
        startActivityForResult(intent, REQ_PHOTO);
    }

    @Override
    public void onRequestPermissionsResult(int requestCode, String[] permissions, int[] results) {
        super.onRequestPermissionsResult(requestCode, permissions, results);
        if (requestCode != REQ_PERMISSION) {
            return;
        }
        boolean granted = results.length > 0 && results[0] == PackageManager.PERMISSION_GRANTED;
        if (granted && waitingPermission) {
            openCamera();
        }
        waitingPermission = false;
    }

    // onActivityResult hands the captured thumbnail or the picked image to
    // the page.
    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        if (requestCode != REQ_PHOTO || resultCode != RESULT_OK || data == null) {
            return;
        }
        Bitmap bitmap = photo(data);
        if (bitmap == null) {
            return;
        }
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        bitmap.compress(Bitmap.CompressFormat.PNG, 100, out);
        Mobile.setPhoto(out.toByteArray());
    }

    // photo returns the image the gallery picker chose, or the thumbnail a
    // camera app attached to its result.
    private Bitmap photo(Intent data) {
        Uri uri = data.getData();
        if (uri != null) {
            try {
                return BitmapFactory.decodeStream(getContentResolver().openInputStream(uri));
            } catch (Exception e) {
                return null;
            }
        }
        if (data.getExtras() == null) {
            return null;
        }
        Object extra = data.getExtras().get("data");
        return extra instanceof Bitmap ? (Bitmap) extra : null;
    }

    @Override
    protected void onResume() {
        super.onResume();
        view.resumeGame();
        poller.post(poll);
    }

    @Override
    protected void onPause() {
        poller.removeCallbacks(poll);
        view.suspendGame();
        super.onPause();
    }
}
