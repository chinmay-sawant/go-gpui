package window

// scrollbarThumb returns the thumb position and length along a track.
func scrollbarThumb(track, content, viewport, offset int) (float32, float32) {
	length := float32(track) * float32(viewport) / float32(content)
	if length < scrollbarMinThumb {
		length = scrollbarMinThumb
	}

	if length > float32(track) {
		length = float32(track)
	}

	maxOffset := content - viewport
	pos := float32(0)

	if maxOffset > 0 && float32(track) > length {
		pos = (float32(track) - length) * float32(offset) / float32(maxOffset)
	}

	return pos, length
}

// scrollbarOffset maps a thumb position back to a scroll offset.
func scrollbarOffset(pos float64, track, content, viewport int) int {
	_, length := scrollbarThumb(track, content, viewport, 0)
	span := float64(track) - float64(length)
	maxOffset := content - viewport

	if maxOffset <= 0 || span <= 0 {
		return 0
	}

	if pos < 0 {
		pos = 0
	}

	if pos > span {
		pos = span
	}

	return int(pos * float64(maxOffset) / span)
}
