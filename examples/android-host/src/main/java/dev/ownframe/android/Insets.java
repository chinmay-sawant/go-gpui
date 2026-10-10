package dev.ownframe.android;

/** Safe system insets and keyboard space, all measured in dp. */
public final class Insets {
    public final float leftDp;
    public final float topDp;
    public final float rightDp;
    public final float bottomDp;
    public final float imeBottomDp;
    public final boolean imeVisible;

    Insets(float left, float top, float right, float bottom,
           float imeBottom, boolean visible) {
        leftDp = left;
        topDp = top;
        rightDp = right;
        bottomDp = bottom;
        imeBottomDp = imeBottom;
        imeVisible = visible;
    }
}
