package scene

import (
	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
	"github.com/chinmay-sawant/ownframe/examples/tetris/store"
)

// themeDark is the stored value for the dark theme.
const themeDark = "dark"

// pageOf slices one page out of a ranked list and assigns ranks.
func pageOf(all []game.Result, page, limit int) ScorePage {
	start := page * PageSize
	if start > len(all) {
		start = len(all)
	}

	end := min(start+PageSize, len(all))
	out := ScorePage{
		Index:   page,
		Total:   len(all),
		More:    len(all) >= limit,
		Entries: make([]ScoreEntry, 0, end-start),
	}

	for i := start; i < end; i++ {
		r := all[i]
		out.Entries = append(out.Entries, ScoreEntry{
			Rank:  i + 1,
			ID:    r.ID,
			Score: r.Score,
			Lines: r.Lines,
			Level: r.Level,
			Dummy: r.Dummy,
		})
	}

	return out
}

// sceneSettings maps the core settings onto the scene view.
func sceneSettings(set store.Settings) Settings {
	return Settings{Dark: set.Theme == themeDark, Control: set}
}

// themeName is the stored theme string.
func themeName(dark bool) string {
	if dark {
		return themeDark
	}

	return "light"
}
