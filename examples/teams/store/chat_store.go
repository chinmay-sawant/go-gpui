package store

import (
	"database/sql"

	"github.com/chinmay-sawant/go-gpui/examples/teams/chat"
)

// saveChat replaces the chat tables with d.
func saveChat(tx *sql.Tx, d chat.Data) error {
	if _, err := tx.Exec(`DELETE FROM chats`); err != nil {
		return err
	}

	if _, err := tx.Exec(`DELETE FROM chat_messages`); err != nil {
		return err
	}

	for i, c := range d.AllChats() {
		_, err := tx.Exec(`INSERT INTO chats (id, position, name, initials, color,
			preview, time, unread, pinned, muted, is_group)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			c.ID, i, c.Name, c.Initials, c.Color, c.Preview, c.Time, c.Unread,
			c.Pinned, c.Muted, c.Group)
		if err != nil {
			return err
		}
	}

	for id, msgs := range d.AllThreads() {
		for i, m := range msgs {
			_, err := tx.Exec(`INSERT INTO chat_messages (id, chat_id, position,
				author, initials, color, time, text, own)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				m.ID, id, i, m.Author, m.Initials, m.Color, m.Time, m.Text, m.Own)
			if err != nil {
				return err
			}
		}
	}

	_, err := tx.Exec(`INSERT OR REPLACE INTO chat_state (id, active, filter, query)
		VALUES (1, ?, ?, ?)`, d.Active, d.Filter, d.Query)

	return err
}

// loadChat reads every chat and thread, then rebuilds the view state.
func loadChat(db *sql.DB) (chat.Data, error) {
	chats, err := loadChats(db)
	if err != nil {
		return chat.Data{}, err
	}

	threads, err := loadThreads(db)
	if err != nil {
		return chat.Data{}, err
	}

	var active, filter, query string

	err = db.QueryRow(`SELECT active, filter, query FROM chat_state WHERE id = 1`).
		Scan(&active, &filter, &query)
	if err != nil && err != sql.ErrNoRows {
		return chat.Data{}, err
	}

	return chat.FromDB(chats, threads, filter, query, active), nil
}
