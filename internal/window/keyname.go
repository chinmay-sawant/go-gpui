package window

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// keyName is the name a key handler sees: the Ebiten key name lowercased,
// with the Digit prefix removed, so "Space" is "space", "ArrowUp" is
// "arrowup", "A" is "a", and "Digit1" is "1". An unknown key is "".
func keyName(key ebiten.Key) string {
	name := strings.ToLower(key.String())

	return strings.TrimPrefix(name, "digit")
}
