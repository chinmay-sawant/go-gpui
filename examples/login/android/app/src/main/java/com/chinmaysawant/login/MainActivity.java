package com.chinmaysawant.login;

import android.app.Activity;
import android.os.Bundle;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import dev.ownframe.android.AndroidHost;
import dev.ownframe.android.Callbacks;
import dev.ownframe.android.InsetMode;
import dev.ownframe.android.Options;
import com.chinmaysawant.login.mobile.EbitenView;

public class MainActivity extends Activity {
    private static final int MATCH = ViewGroup.LayoutParams.MATCH_PARENT;
    private EbitenView view;
    private AndroidHost host;

    @Override public void onCreate(Bundle state) {
        super.onCreate(state);
        FrameLayout layout = new FrameLayout(this);
        view = new EbitenView(this);
        layout.addView(view, new FrameLayout.LayoutParams(MATCH, MATCH));
        setContentView(layout);
        host = AndroidHost.attach(this, layout,
            Options.builder().edgeToEdge(true).insetMode(InsetMode.HOST_PADDING).build(),
            new Callbacks() {
                @Override public void onResume() { view.resumeGame(); }
                @Override public void onPause() { view.suspendGame(); }
            });
    }

    @Override protected void onDestroy() {
        if (host != null) host.close();
        super.onDestroy();
    }
}
