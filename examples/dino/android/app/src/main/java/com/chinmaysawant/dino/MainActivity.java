package com.chinmaysawant.dino;

import android.app.Activity;
import android.os.Bundle;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import com.chinmaysawant.dino.mobile.EbitenView;

// MainActivity fills the screen with the bound go-gpui dino page.
public class MainActivity extends Activity {
    private static final int MATCH = ViewGroup.LayoutParams.MATCH_PARENT;

    private EbitenView view;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        FrameLayout layout = new FrameLayout(this);
        layout.setLayoutParams(new ViewGroup.LayoutParams(MATCH, MATCH));

        view = new EbitenView(this);
        layout.addView(view, new FrameLayout.LayoutParams(MATCH, MATCH));
        setContentView(layout);
    }

    @Override
    protected void onResume() {
        super.onResume();
        view.resumeGame();
    }

    @Override
    protected void onPause() {
        view.suspendGame();
        super.onPause();
    }
}
