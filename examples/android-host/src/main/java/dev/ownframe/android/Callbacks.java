package dev.ownframe.android;

/** Optional app hooks. Implement only the behavior the application owns. */
public interface Callbacks {
    default void onInsetsChanged(Insets insets) {}
    default void onConfigurationChanged(DisplayInfo display) {}
    default void onViewportChanged(ViewportSnapshot viewport) {}
    default void onImeBackPressed() {}
    default boolean onBackPressed() { return false; }
    default void onResume() {}
    default void onPause() {}
}
