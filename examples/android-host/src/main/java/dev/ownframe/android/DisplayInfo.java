package dev.ownframe.android;

/** Current display density and Android-scaled base text size. */
public final class DisplayInfo {
    public final float density;
    public final float fontScale;
    public final float baseTextSizeDp;

    DisplayInfo(float density, float fontScale, float baseTextSizeDp) {
        this.density = density;
        this.fontScale = fontScale;
        this.baseTextSizeDp = baseTextSizeDp;
    }
}
