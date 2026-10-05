package calendar

// week5 is the sample week of October 19, 2026.
var week5 = week{
	label: "October 19 – October 23",
	today: -1,
	days: []Day{
		{"Mon 10/19", false},
		{"Tue 10/20", false},
		{"Wed 10/21", false},
		{"Thu 10/22", false},
		{"Fri 10/23", false},
	},
	events: []Event{
		{"e34", "Avengers daily standup", "9:00 AM", "15 min", "blue", "Avengers Tower 42", 0, false},
		{"e35", "Wakanda stealth suit review", "11:00 AM", "45 min", "violet", "Wakanda Lab 3", 0, false},
		{"e36", "1:1 with Peter Parker", "9:30 AM", "30 min", "green", "Queens", 1, false},
		{"e37", "Westview anomaly follow-up", "1:00 PM", "45 min", "orange", "Sanctum Sanctorum", 2, false},
		{"e38", "Operation planning with Maria Hill", "11:00 AM", "60 min", "violet", "S.H.I.E.L.D. Helicarrier", 3, false},
		{"e39", "Shawarma Friday", "3:00 PM", "45 min", "red", "Avengers Tower 42", 4, false},
	},
}
