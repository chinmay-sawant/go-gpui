package store

import (
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/teams/channels"
)

// loadReplies appends each stored reply to its post.
func loadReplies(db *sql.DB, index map[string]*channels.Post) error {
	type row struct {
		post  string
		reply channels.Reply
	}

	var rows []row

	err := readInto(db, `SELECT id, post_id, author, initials, color, time, text, own FROM replies ORDER BY post_id, position`, &rows, func(r *sql.Rows, v *row) error {
		return r.Scan(&v.reply.ID, &v.post, &v.reply.Author, &v.reply.Initials, &v.reply.Color, &v.reply.Time, &v.reply.Text, &v.reply.Own)
	})
	if err != nil {
		return err
	}

	for _, v := range rows {
		if p := index[v.post]; p != nil {
			p.Replies = append(p.Replies, v.reply)
		}
	}

	return nil
}

// loadReactions appends each stored reaction to its post.
func loadReactions(db *sql.DB, index map[string]*channels.Post) error {
	type row struct {
		post string
		re   channels.Reaction
	}

	var rows []row

	err := readInto(db, `SELECT post_id, emoji, count, mine FROM reactions ORDER BY post_id, position`, &rows, func(r *sql.Rows, v *row) error {
		return r.Scan(&v.post, &v.re.Emoji, &v.re.Count, &v.re.Mine)
	})
	if err != nil {
		return err
	}

	for _, v := range rows {
		if p := index[v.post]; p != nil {
			p.Reactions = append(p.Reactions, v.re)
		}
	}

	return nil
}
