package channels

import (
	"strconv"
	"strings"
)

// toggleReaction flips one chip's own state and count.
func toggleReaction(d *Data, spec string) {
	id, n, ok := splitIndex(spec)
	if !ok {
		return
	}

	for i := range d.Posts {
		p := &d.Posts[i]
		if p.ID != id || n < 0 || n >= len(p.Reactions) {
			continue
		}

		r := &p.Reactions[n]
		r.Mine = !r.Mine
		if r.Mine {
			r.Count++
		} else if r.Count > 0 {
			r.Count--
		}

		return
	}
}

// addReaction adds picker emoji i to one post, or increments it.
func addReaction(d *Data, spec string) {
	id, n, ok := splitIndex(spec)
	if !ok || n < 0 || n >= len(d.Emojis) {
		return
	}

	emoji := d.Emojis[n]
	for i := range d.Posts {
		p := &d.Posts[i]
		if p.ID != id {
			continue
		}

		p.Picker = false
		for j := range p.Reactions {
			if p.Reactions[j].Emoji == emoji {
				p.Reactions[j].Count++
				p.Reactions[j].Mine = true

				return
			}
		}

		p.Reactions = append(p.Reactions, Reaction{Emoji: emoji, Count: 1, Mine: true})

		return
	}
}

// togglePicker flips one post's emoji picker.
func togglePicker(d *Data, id string) {
	for i := range d.Posts {
		if d.Posts[i].ID == id {
			d.Posts[i].Picker = !d.Posts[i].Picker

			return
		}
	}
}

// splitIndex splits "post-3" into "post" and 3.
func splitIndex(s string) (string, int, bool) {
	i := strings.LastIndex(s, "-")
	if i < 0 {
		return s, 0, false
	}

	n, err := strconv.Atoi(s[i+1:])
	if err != nil {
		return s, 0, false
	}

	return s[:i], n, true
}
