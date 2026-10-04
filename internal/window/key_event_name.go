package window

import "github.com/hajimehoshi/ebiten/v2"

// keyEventName is the name the screen sees. The caret keys carry the
// shift+, ctrl+, and alt+ prefixes the page understands, so the window can
// extend a selection or jump by word. Every other key keeps its plain name.
func keyEventName(key ebiten.Key, mods modifiers) string {
	name := keyName(key)
	if name == "" {
		return ""
	}

	switch name {
	case "arrowleft", "arrowright", "home", "end":
	default:
		return name
	}

	if mods.Shift {
		name = "shift+" + name
	}

	if mods.command() {
		name = "ctrl+" + name
	} else if mods.Alt {
		name = "alt+" + name
	}

	return name
}
