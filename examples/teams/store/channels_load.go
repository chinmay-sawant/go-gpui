package store

import (
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/teams/channels"
)

// loadChannels reads every table back and rebuilds the view state.
func loadChannels(db *sql.DB) (channels.Data, error) {
	teams, err := loadTeams(db)
	if err != nil {
		return channels.Data{}, err
	}

	posts, err := loadPosts(db)
	if err != nil {
		return channels.Data{}, err
	}

	index := postIndex(posts)

	if err := loadReplies(db, index); err != nil {
		return channels.Data{}, err
	}

	if err := loadReactions(db, index); err != nil {
		return channels.Data{}, err
	}

	files, err := loadChannelFiles(db)
	if err != nil {
		return channels.Data{}, err
	}

	var team, channel, tab string

	err = db.QueryRow(`SELECT team, channel, tab FROM channel_state WHERE id = 1`).Scan(&team, &channel, &tab)
	if err != nil && err != sql.ErrNoRows {
		return channels.Data{}, err
	}

	return channels.FromDB(teams, posts, files, team, channel, tab), nil
}

// loadTeams reads the teams, then attaches their channels in order.
func loadTeams(db *sql.DB) ([]channels.Team, error) {
	var teams []channels.Team

	err := readInto(db, `SELECT id, name, initials, color, expanded FROM teams ORDER BY position`, &teams, func(r *sql.Rows, t *channels.Team) error {
		return r.Scan(&t.ID, &t.Name, &t.Initials, &t.Color, &t.Expanded)
	})
	if err != nil {
		return nil, err
	}

	type row struct {
		team string
		ch   channels.Channel
	}

	var rows []row

	err = readInto(db, `SELECT c.id, c.team_id, c.name, c.unread FROM channels c JOIN teams t ON t.id = c.team_id ORDER BY t.position, c.position`, &rows, func(r *sql.Rows, v *row) error {
		return r.Scan(&v.ch.ID, &v.team, &v.ch.Name, &v.ch.Unread)
	})
	if err != nil {
		return nil, err
	}

	byID := map[string]int{}
	for i, t := range teams {
		byID[t.ID] = i
	}

	for _, v := range rows {
		if i, ok := byID[v.team]; ok {
			teams[i].Channels = append(teams[i].Channels, v.ch)
		}
	}

	return teams, nil
}
