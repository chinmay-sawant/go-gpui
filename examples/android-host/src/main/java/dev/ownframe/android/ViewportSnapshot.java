package dev.ownframe.android;

/** Immutable measured viewport and Android display state, in dp unless named otherwise. */
public final class ViewportSnapshot {
    public final float contentWidthDp;
    public final float contentHeightDp;
    public final float safeLeftDp;
    public final float safeTopDp;
    public final float safeRightDp;
    public final float safeBottomDp;
    public final float imeBottomDp;
    public final boolean imeVisible;
    public final float density;
    public final float fontScale;
    public final float baseTextSizeDp;
    public final long revision;

    ViewportSnapshot(float width, float height, Insets insets, DisplayInfo display,
                     long revision) {
        contentWidthDp = width;
        contentHeightDp = height;
        safeLeftDp = insets.leftDp;
        safeTopDp = insets.topDp;
        safeRightDp = insets.rightDp;
        safeBottomDp = insets.bottomDp;
        imeBottomDp = insets.imeBottomDp;
        imeVisible = insets.imeVisible;
        density = display.density;
        fontScale = display.fontScale;
        baseTextSizeDp = display.baseTextSizeDp;
        this.revision = revision;
    }

    boolean hasSameValues(ViewportSnapshot other) {
        return Float.compare(contentWidthDp, other.contentWidthDp) == 0
                && Float.compare(contentHeightDp, other.contentHeightDp) == 0
                && Float.compare(safeLeftDp, other.safeLeftDp) == 0
                && Float.compare(safeTopDp, other.safeTopDp) == 0
                && Float.compare(safeRightDp, other.safeRightDp) == 0
                && Float.compare(safeBottomDp, other.safeBottomDp) == 0
                && Float.compare(imeBottomDp, other.imeBottomDp) == 0
                && imeVisible == other.imeVisible
                && Float.compare(density, other.density) == 0
                && Float.compare(fontScale, other.fontScale) == 0
                && Float.compare(baseTextSizeDp, other.baseTextSizeDp) == 0;
    }
}
