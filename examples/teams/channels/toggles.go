package channels

// toggleTeam flips one team's expanded state.
func toggleTeam(d *Data, id string) {
	for i := range d.Teams {
		if d.Teams[i].ID == id {
			d.Teams[i].Expanded = !d.Teams[i].Expanded

			return
		}
	}
}

// toggleLike flips one post's like and keeps the count in step.
func toggleLike(d *Data, id string) {
	for i := range d.Posts {
		if d.Posts[i].ID != id {
			continue
		}

		if d.Posts[i].Liked {
			d.Posts[i].Liked = false
			if d.Posts[i].Likes > 0 {
				d.Posts[i].Likes--
			}
		} else {
			d.Posts[i].Liked = true
			d.Posts[i].Likes++
		}

		return
	}
}

// togglePin flips one post's pinned flag.
func togglePin(d *Data, id string) {
	for i := range d.Posts {
		if d.Posts[i].ID == id {
			d.Posts[i].Pinned = !d.Posts[i].Pinned

			return
		}
	}
}

// toggleStar flips one file's starred flag.
func toggleStar(d *Data, id string) {
	for i := range d.Files {
		if d.Files[i].ID == id {
			d.Files[i].Starred = !d.Files[i].Starred

			return
		}
	}
}
