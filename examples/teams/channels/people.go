package channels

import "strconv"

// repPeople maps a reply author to avatar initials and color.
var repPeople = map[string][2]string{
	"Robert Downey Jr.": {"RDJ", "red"},
	"Steve Rogers":      {"SR", "blue"},
	"Natasha Romanoff":  {"NR", "red"},
	"Thor":              {"TH", "gold"},
	"Bruce Banner":      {"BB", "green"},
	"Peter Parker":      {"PP", "red"},
	"Stephen Strange":   {"SS", "gray"},
	"Wanda Maximoff":    {"WM", "pink"},
	"Sam Wilson":        {"SW", "teal"},
	"Clint Barton":      {"CB", "purple"},
	"Shuri":             {"SH", "purple"},
	"T'Challa":          {"TC", "violet"},
	"Okoye":             {"OK", "orange"},
	"Nick Fury":         {"NF", "gray"},
	"Maria Hill":        {"MH", "blue"},
	"Carol Danvers":     {"CD", "gold"},
	"Scott Lang":        {"SL", "red"},
	"James Rhodes":      {"JR", "gray"},
	"Loki":              {"LK", "green"},
	"Vision":            {"VI", "violet"},
	"Happy Hogan":       {"HH", "orange"},
	"Pepper Potts":      {"PP", "pink"},
}

// rep builds one reply by person name.
func rep(author, time, text string) Reply {
	p := repPeople[author]

	return Reply{Author: author, Initials: p[0], Color: p[1], Time: time, Text: text}
}

// thread numbers the replies of one post.
func thread(id string, rows []Reply) []Reply {
	for i := range rows {
		rows[i].ID = id + "-r" + strconv.Itoa(i+1)
	}

	return rows
}
