package calendar

// week6 is the sample week of October 26, 2026.
var week6 = week{
	label: "October 26 – October 30",
	today: -1,
	days: []Day{
		{"Mon 10/26", false},
		{"Tue 10/27", false},
		{"Wed 10/28", false},
		{"Thu 10/29", false},
		{"Fri 10/30", false},
	},
	events: []Event{
		{"e40", "Avengers all-hands", "9:00 AM", "45 min", "violet", "Avengers Compound", 0, false},
		{"e41", "Stark budget review - FY27", "11:00 AM", "60 min", "red", "Stark Industries R&D", 0, false},
		{"e42", "Avengers daily standup", "9:00 AM", "15 min", "blue", "Avengers Tower 42", 1, false},
		{"e43", "Web shooter workshop with Peter Parker", "10:00 AM", "90 min", "orange", "Queens", 2, false},
		{"e44", "1:1 with Carol Danvers", "2:00 PM", "30 min", "green", "Avengers Tower 42", 2, false},
		{"e45", "Sanctum consult with Wong", "2:00 PM", "45 min", "orange", "Sanctum Sanctorum", 3, false},
		{"e46", "Mark 85 flight demo", "3:30 PM", "45 min", "blue", "Stark Industries R&D", 4, false},
	},
}
