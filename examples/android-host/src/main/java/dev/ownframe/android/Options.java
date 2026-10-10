package dev.ownframe.android;

/** Per-activity Android host behavior. Values are immutable after build. */
public final class Options {
    public final boolean edgeToEdge;
    public final InsetMode insetMode;
    public final float baseTextSizeSp;

    private Options(Builder builder) {
        edgeToEdge = builder.edgeToEdge;
        insetMode = builder.insetMode;
        baseTextSizeSp = builder.baseTextSizeSp;
    }

    public static Builder builder() { return new Builder(); }

    public static final class Builder {
        private boolean edgeToEdge = true;
        private InsetMode insetMode = InsetMode.HOST_PADDING;
        private float baseTextSizeSp = 16f;

        public Builder edgeToEdge(boolean value) {
            edgeToEdge = value;
            return this;
        }

        public Builder insetMode(InsetMode value) {
            if (value == null) throw new IllegalArgumentException("insetMode is required");
            insetMode = value;
            return this;
        }

        public Builder baseTextSizeSp(float value) {
            if (!(value > 0f)) throw new IllegalArgumentException("text size must be positive");
            baseTextSizeSp = value;
            return this;
        }

        public Options build() { return new Options(this); }
    }
}
