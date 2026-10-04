package host

// Shape is a window cursor shape.
type Shape int

const (
	// ShapeDefault is the plain arrow.
	ShapeDefault Shape = iota
	// ShapeText is the I-beam over a text field.
	ShapeText
	// ShapePointer is the hand over a link or control.
	ShapePointer
	// ShapeResizeEW is the horizontal resize over a scrollbar thumb.
	ShapeResizeEW
	// ShapeResizeNS is the vertical resize over a scrollbar thumb.
	ShapeResizeNS
)

// CursorShape is a screen that names the cursor shape for the element the
// pointer is over.
type CursorShape interface {
	CursorShape() Shape
}
