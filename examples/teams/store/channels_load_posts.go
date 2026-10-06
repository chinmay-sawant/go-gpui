package store

import (
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/teams/channels"
)

// loadPosts reads every post, grouped by channel id.
func loadPosts(db *sql.DB) (map[string][]channels.Post, error) {
	type row struct {
		channel string
		post    channels.Post
	}

	var rows []row

	err := readInto(db, `SELECT id, channel_id, author, initials, color, time, subject, text, pinned, likes, liked FROM posts ORDER BY channel_id, position`, &rows, func(r *sql.Rows, v *row) error {
		return r.Scan(&v.post.ID, &v.channel, &v.post.Author, &v.post.Initials, &v.post.Color, &v.post.Time, &v.post.Subject, &v.post.Text, &v.post.Pinned, &v.post.Likes, &v.post.Liked)
	})
	if err != nil {
		return nil, err
	}

	posts := map[string][]channels.Post{}
	for _, v := range rows {
		posts[v.channel] = append(posts[v.channel], v.post)
	}

	return posts, nil
}

// postIndex maps every post id to its slot in posts.
func postIndex(posts map[string][]channels.Post) map[string]*channels.Post {
	index := map[string]*channels.Post{}
	for _, list := range posts {
		for i := range list {
			index[list[i].ID] = &list[i]
		}
	}

	return index
}

// loadChannelFiles reads the file rows in position order.
func loadChannelFiles(db *sql.DB) ([]channels.FileItem, error) {
	var files []channels.FileItem

	err := readInto(db, `SELECT id, name, badge, kind, modified, modified_by, size, starred FROM channel_files ORDER BY position`, &files, func(r *sql.Rows, f *channels.FileItem) error {
		return r.Scan(&f.ID, &f.Name, &f.Badge, &f.Kind, &f.Modified, &f.ModifiedBy, &f.Size, &f.Starred)
	})

	return files, err
}
