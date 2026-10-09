package bridge

var fontSize = 16

// SetFontSize records Android's system-scaled text size.
func SetFontSize(size int) {
	if size < 16 {
		size = 16
	}
	if size > 48 {
		size = 48
	}
	mu.Lock()
	fontSize = size
	mu.Unlock()
}

// FontSize is the current base text size in CSS pixels.
func FontSize() int {
	mu.Lock()
	defer mu.Unlock()
	return fontSize
}
