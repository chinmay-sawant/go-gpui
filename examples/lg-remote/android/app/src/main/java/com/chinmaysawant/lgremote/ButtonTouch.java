package com.chinmaysawant.lgremote;

import android.graphics.Rect;
import android.view.MotionEvent;
import com.chinmaysawant.lgremote.mobile.Mobile;

// Queues both edges so taps shorter than a game-loop frame are preserved.
final class ButtonTouch {
    private String held = "";
    private Rect area;
    private int pointer;
    private boolean tracking;

    boolean start(String id, Rect bounds, MotionEvent event, boolean enabled) {
        tracking = true;
        pointer = event.getPointerId(0);
        area = new Rect(bounds);
        held = enabled && Mobile.queueAction("down:" + id) ? id : "";
        return true;
    }

    boolean handle(MotionEvent event) {
        if (!tracking) { return false; }
        switch (event.getActionMasked()) {
        case MotionEvent.ACTION_MOVE:
            int index = event.findPointerIndex(pointer);
            if (index < 0 || !area.contains((int) event.getX(index), (int) event.getY(index))) {
                cancelHeld();
            }
            break;
        case MotionEvent.ACTION_POINTER_DOWN:
        case MotionEvent.ACTION_CANCEL:
            cancelHeld();
            break;
        case MotionEvent.ACTION_UP:
            int last = event.findPointerIndex(pointer);
            if (last < 0 || !area.contains((int) event.getX(last), (int) event.getY(last))) {
                cancelHeld();
            }
            if (!held.isEmpty()) { Mobile.queueAction("up:" + held); }
            held = "";
            tracking = false;
            break;
        }
        if (event.getActionMasked() == MotionEvent.ACTION_CANCEL) { tracking = false; }
        return true;
    }

    private void cancelHeld() {
        if (!held.isEmpty()) { Mobile.queueAction("cancel:"); }
        held = "";
    }

    void cancel() {
        cancelHeld();
        tracking = false;
    }
}
