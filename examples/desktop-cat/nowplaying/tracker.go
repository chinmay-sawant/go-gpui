package nowplaying

type tracker struct{ candidate, last string }

func (t *tracker) next(track Track) string {
	if !track.Playing || track.Title == "" {
		t.candidate = ""
		return ""
	}
	key := track.App + "\x00" + track.Title + "\x00" + track.Artist
	ready := key == t.candidate && key != t.last
	t.candidate = key
	if ready {
		return key
	}
	return ""
}
