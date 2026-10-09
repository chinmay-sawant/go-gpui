package com.chinmaysawant.lgremote;

import android.content.Context;
import android.graphics.Rect;
import android.os.Bundle;
import android.view.MotionEvent;
import android.view.View;
import android.view.accessibility.AccessibilityEvent;
import android.view.accessibility.AccessibilityManager;
import android.view.accessibility.AccessibilityNodeInfo;
import android.view.accessibility.AccessibilityNodeProvider;
import com.chinmaysawant.lgremote.mobile.Mobile;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
import java.util.LinkedHashMap;
import java.util.Map;

// Exposes the Go canvas controls as native virtual accessibility nodes.
final class RemoteAccessibility extends View {
    private final Map<Integer, JSONObject> nodes = new LinkedHashMap<>();
    private final Map<Integer, Rect> bounds = new LinkedHashMap<>();
    private String snapshot = "";
    private int focus = Integer.MIN_VALUE;
    private int hovered = Integer.MIN_VALUE;
    private boolean scrollable;
    private final ButtonTouch touch = new ButtonTouch();
    private final AccessibilityNodeProvider provider = new Provider();

    RemoteAccessibility(Context context) {
        super(context);
        setImportantForAccessibility(IMPORTANT_FOR_ACCESSIBILITY_YES);
        setFocusable(false);
    }

    void refresh() {
        String data = Mobile.accessibility();
        if (data == null || data.isEmpty() || data.equals(snapshot) || getWidth() == 0) { return; }
        try {
            JSONObject state = new JSONObject(data);
            scrollable = state.optBoolean("Scrollable");
            float sx = getWidth() / (float) state.getInt("Width");
            float sy = getHeight() / (float) state.getInt("Height");
            nodes.clear();
            bounds.clear();
            JSONArray list = state.getJSONArray("Nodes");
            for (int i = 0; i < list.length(); i++) {
                JSONObject node = list.getJSONObject(i);
                int id = node.getString("ID").hashCode() & 0x7fffffff;
                int x = Math.round((float) node.getDouble("X") * sx);
                int y = Math.round((float) node.getDouble("Y") * sy);
                int w = Math.round((float) node.getDouble("W") * sx);
                int h = Math.round((float) node.getDouble("H") * sy);
                Rect rect = new Rect(x, y, x+w, y+h);
                if (rect.intersect(0, 0, getWidth(), getHeight())) {
                    nodes.put(id, node);
                    bounds.put(id, rect);
                }
            }
            if (!nodes.containsKey(focus)) { focus = Integer.MIN_VALUE; }
            snapshot = data;
            event(HOST_ID, AccessibilityEvent.TYPE_WINDOW_CONTENT_CHANGED);
        } catch (JSONException ex) {
            // Keep the previous snapshot until the next valid game-loop frame.
        }
    }

    @Override public AccessibilityNodeProvider getAccessibilityNodeProvider() { return provider; }

    void cancelTouch() { touch.cancel(); }

    @Override public boolean onTouchEvent(MotionEvent event) {
        if (event.getActionMasked() == MotionEvent.ACTION_DOWN) {
            touch.cancel();
            for (Map.Entry<Integer, Rect> item : bounds.entrySet()) {
                if (!item.getValue().contains((int) event.getX(), (int) event.getY())) { continue; }
                JSONObject node = nodes.get(item.getKey());
                String id = node.optString("ID");
                if (!node.optBoolean("Button") || "pad".equals(id)) { continue; }
                return touch.start(id, item.getValue(), event, node.optBoolean("Enabled"));
            }
            return false;
        }
        return touch.handle(event);
    }

    @Override public boolean dispatchHoverEvent(MotionEvent event) {
        android.view.accessibility.AccessibilityManager manager =
            (android.view.accessibility.AccessibilityManager) getContext()
                .getSystemService(Context.ACCESSIBILITY_SERVICE);
        if (manager == null || !manager.isTouchExplorationEnabled()) { return false; }
        int next = Integer.MIN_VALUE;
        if (event.getAction() != MotionEvent.ACTION_HOVER_EXIT) {
            for (Map.Entry<Integer, Rect> item : bounds.entrySet()) {
                if (item.getValue().contains((int) event.getX(), (int) event.getY())) {
                    next = item.getKey();
                    break;
                }
            }
        }
        if (next != hovered) {
            if (next != Integer.MIN_VALUE) { event(next, AccessibilityEvent.TYPE_VIEW_HOVER_ENTER); }
            if (hovered != Integer.MIN_VALUE) { event(hovered, AccessibilityEvent.TYPE_VIEW_HOVER_EXIT); }
            hovered = next;
        }
        return true;
    }

    private static final int HOST_ID = AccessibilityNodeProvider.HOST_VIEW_ID;

    private void event(int id, int type) {
        AccessibilityManager manager = getContext().getSystemService(AccessibilityManager.class);
        if (manager == null || !manager.isEnabled() || getParent() == null) { return; }
        AccessibilityEvent e = AccessibilityEvent.obtain(type);
        e.setPackageName(getContext().getPackageName());
        e.setSource(this, id);
        JSONObject node = nodes.get(id);
        e.setClassName(node != null && node.optBoolean("Editable") ? "android.widget.EditText" : "android.widget.Button");
        if (node != null) { e.setContentDescription(node.optString("Name")); }
        getParent().requestSendAccessibilityEvent(this, e);
    }

    private final class Provider extends AccessibilityNodeProvider {
        @Override public AccessibilityNodeInfo createAccessibilityNodeInfo(int id) {
            if (id == HOST_ID) {
                AccessibilityNodeInfo root = AccessibilityNodeInfo.obtain(RemoteAccessibility.this);
                onInitializeAccessibilityNodeInfo(root);
                root.setClassName("android.widget.ScrollView");
                root.setScrollable(scrollable);
                if (scrollable) {
                    root.addAction(AccessibilityNodeInfo.ACTION_SCROLL_FORWARD);
                    root.addAction(AccessibilityNodeInfo.ACTION_SCROLL_BACKWARD);
                }
                for (int child : nodes.keySet()) { root.addChild(RemoteAccessibility.this, child); }
                return root;
            }
            JSONObject data = nodes.get(id);
            if (data == null) { return null; }
            AccessibilityNodeInfo n = AccessibilityNodeInfo.obtain();
            n.setSource(RemoteAccessibility.this, id);
            n.setParent(RemoteAccessibility.this);
            n.setPackageName(getContext().getPackageName());
            boolean editable = data.optBoolean("Editable");
            boolean slider = data.optBoolean("Slider");
            boolean button = data.optBoolean("Button");
            boolean enabled = data.optBoolean("Enabled");
            n.setClassName(slider ? "android.widget.SeekBar" : editable ? "android.widget.EditText" : button ? "android.widget.Button" : "android.widget.TextView");
            n.setContentDescription(data.optString("Name"));
            n.setEnabled(enabled);
            n.setFocusable(true);
            n.setVisibleToUser(true);
            n.setSelected(data.optBoolean("Selected"));
            n.setAccessibilityFocused(focus == id);
            n.addAction(focus == id ? AccessibilityNodeInfo.ACTION_CLEAR_ACCESSIBILITY_FOCUS : AccessibilityNodeInfo.ACTION_ACCESSIBILITY_FOCUS);
            n.setClickable(enabled && (button || editable));
            if (enabled && (button || editable)) { n.addAction(AccessibilityNodeInfo.ACTION_CLICK); }
            n.setEditable(editable);
            if (slider) {
                n.setContentDescription(data.optString("Name") + ", " + data.optString("Value"));
                n.setRangeInfo(AccessibilityNodeInfo.RangeInfo.obtain(AccessibilityNodeInfo.RangeInfo.RANGE_TYPE_FLOAT,
                    .25f, 3f, data.optInt("Progress", 100)/100f));
                n.addAction(AccessibilityNodeInfo.AccessibilityAction.ACTION_SET_PROGRESS);
                n.addAction(AccessibilityNodeInfo.ACTION_SCROLL_FORWARD);
                n.addAction(AccessibilityNodeInfo.ACTION_SCROLL_BACKWARD);
            }
            if (editable) {
                n.setText(data.optString("Value"));
                n.addAction(AccessibilityNodeInfo.ACTION_SET_TEXT);
            }
            Rect rect = new Rect(bounds.get(id));
            n.setBoundsInParent(rect);
            int[] location = new int[2];
            getLocationOnScreen(location);
            rect.offset(location[0], location[1]);
            n.setBoundsInScreen(rect);
            return n;
        }

        @Override public AccessibilityNodeInfo findFocus(int kind) {
            return kind == AccessibilityNodeInfo.FOCUS_ACCESSIBILITY && nodes.containsKey(focus)
                ? createAccessibilityNodeInfo(focus) : null;
        }

        @Override public boolean performAction(int id, int action, Bundle args) {
            if (id == HOST_ID) {
                if (!scrollable) { return false; }
                if (action == AccessibilityNodeInfo.ACTION_SCROLL_FORWARD) { return Mobile.queueAction("scroll:forward"); }
                if (action == AccessibilityNodeInfo.ACTION_SCROLL_BACKWARD) { return Mobile.queueAction("scroll:back"); }
                return false;
            }
            JSONObject node = nodes.get(id);
            if (node == null) { return false; }
            if (action == AccessibilityNodeInfo.ACTION_ACCESSIBILITY_FOCUS) {
                if (focus != Integer.MIN_VALUE && focus != id) { event(focus, AccessibilityEvent.TYPE_VIEW_ACCESSIBILITY_FOCUS_CLEARED); }
                focus = id;
                event(id, AccessibilityEvent.TYPE_VIEW_ACCESSIBILITY_FOCUSED);
                return true;
            }
            if (action == AccessibilityNodeInfo.ACTION_CLEAR_ACCESSIBILITY_FOCUS) {
                if (focus != id) { return false; }
                focus = Integer.MIN_VALUE;
                event(id, AccessibilityEvent.TYPE_VIEW_ACCESSIBILITY_FOCUS_CLEARED);
                return true;
            }
            if (!node.optBoolean("Enabled")) { return false; }
            if (node.optBoolean("Slider")) {
                int value = node.optInt("Progress", 100);
                if (action == AccessibilityNodeInfo.ACTION_SCROLL_FORWARD) { value += 5; }
                else if (action == AccessibilityNodeInfo.ACTION_SCROLL_BACKWARD) { value -= 5; }
                else if (action == AccessibilityNodeInfo.AccessibilityAction.ACTION_SET_PROGRESS.getId() && args != null) {
                    value = Math.round(args.getFloat(AccessibilityNodeInfo.ACTION_ARGUMENT_PROGRESS_VALUE)*100);
                } else { return false; }
                return Mobile.queueAction("sensitivity:" + Math.max(25, Math.min(300, value)));
            }
            if (action == AccessibilityNodeInfo.ACTION_CLICK) {
                boolean queued = Mobile.queueAction("click:"+node.optString("ID"));
                if (queued) { event(id, AccessibilityEvent.TYPE_VIEW_CLICKED); }
                return queued;
            }
            if (action == AccessibilityNodeInfo.ACTION_SET_TEXT && node.optBoolean("Editable") && args != null) {
                CharSequence text = args.getCharSequence(AccessibilityNodeInfo.ACTION_ARGUMENT_SET_TEXT_CHARSEQUENCE);
                return text != null && Mobile.queueAction("text:"+text);
            }
            return false;
        }
    }
}
