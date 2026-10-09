package com.chinmaysawant.lgremote;

import android.app.Activity;
import android.app.Instrumentation;
import android.os.Bundle;
import android.view.View;
import android.view.accessibility.AccessibilityEvent;
import android.view.accessibility.AccessibilityManager;
import android.view.accessibility.AccessibilityNodeProvider;
import android.widget.FrameLayout;
import java.lang.reflect.Method;

// Runs on a device with accessibility services disabled, matching the startup crash.
public final class AccessibilityRegression extends Instrumentation {
    @Override public void onCreate(Bundle arguments) {
        super.onCreate(arguments);
        start();
    }

    @Override public void onStart() {
        Bundle result = new Bundle();
        try {
            final Throwable[] failure = {null};
            runOnMainSync(() -> {
                try { checkDisabledEvents(); }
                catch (Throwable error) { failure[0] = error; }
            });
            if (failure[0] != null) { throw failure[0]; }
            result.putString("stream", "PASS: disabled accessibility did not dispatch an event\n");
            finish(Activity.RESULT_OK, result);
        } catch (Throwable error) {
            result.putString("stream", "FAIL: " + error + "\n");
            finish(Activity.RESULT_CANCELED, result);
        }
    }

    private void checkDisabledEvents() {
        AccessibilityManager manager = getTargetContext().getSystemService(AccessibilityManager.class);
        if (manager.isEnabled()) {
            throw new AssertionError("Run this test with accessibility services disabled");
        }
        FrameLayout parent = new FrameLayout(getTargetContext()) {
            @Override public boolean requestSendAccessibilityEvent(View child, AccessibilityEvent event) {
                throw new AssertionError("Disabled accessibility received an event");
            }
        };
        RemoteAccessibility view = new RemoteAccessibility(getTargetContext());
        parent.addView(view);
        try {
            Method event = RemoteAccessibility.class.getDeclaredMethod("event", int.class, int.class);
            event.setAccessible(true);
            event.invoke(view, AccessibilityNodeProvider.HOST_VIEW_ID,
                AccessibilityEvent.TYPE_WINDOW_CONTENT_CHANGED);
        } catch (ReflectiveOperationException error) {
            throw new AssertionError(error);
        }
    }
}
