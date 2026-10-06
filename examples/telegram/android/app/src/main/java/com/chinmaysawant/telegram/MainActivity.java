package com.chinmaysawant.telegram;

import android.app.Activity;
import android.os.Bundle;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import com.chinmaysawant.telegram.mobile.EbitenView;

// MainActivity shows the bound go-gpui page full screen. EbitenView is the
// class the ebitenmobile-generated AAR defines for the mobile package.
public class MainActivity extends Activity {
    private EbitenView view;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        FrameLayout layout = new FrameLayout(this);
        layout.setLayoutParams(new ViewGroup.LayoutParams(
            ViewGroup.LayoutParams.MATCH_PARENT,
            ViewGroup.LayoutParams.MATCH_PARENT));

        view = new EbitenView(this);
        layout.addView(view, new FrameLayout.LayoutParams(
            ViewGroup.LayoutParams.MATCH_PARENT,
            ViewGroup.LayoutParams.MATCH_PARENT));

        setContentView(layout);
    }

    // The game stops while the app is not visible and resumes on return.
    @Override
    protected void onPause() {
        view.suspendGame();
        super.onPause();
    }

    @Override
    protected void onResume() {
        super.onResume();
        view.resumeGame();
    }
}
