package calendar

// week3 is the sample week of October 5, 2026.
var week3 = week{
	label: "October 5 – October 9",
	today: 0,
	days: []Day{
		{"Mon 10/5", true},
		{"Tue 10/6", false},
		{"Wed 10/7", false},
		{"Thu 10/8", false},
		{"Fri 10/9", false},
	},
	events: []Event{
		{"e1", "Avengers daily standup", "9:00 AM", "15 min", "blue", "Avengers Tower 42", 0, false},
		{"e2", "Suit diagnostics with Shuri", "10:30 AM", "45 min", "violet", "Wakanda Lab 3", 0, false},
		{"e3", "1:1 with Sam Wilson", "1:00 PM", "30 min", "green", "Avengers Compound", 1, false},
		{"e4", "Sanctum consult with Doctor Strange", "9:30 AM", "30 min", "blue", "Sanctum Sanctorum", 2, false},
		{"e5", "Asgard observatory briefing", "2:30 PM", "45 min", "orange", "Asgard observatory", 2, false},
		{"e6", "S.H.I.E.L.D. mission briefing", "11:00 AM", "60 min", "violet", "S.H.I.E.L.D. Helicarrier", 3, false},
		{"e7", "Web shooter bug triage", "3:00 PM", "30 min", "red", "Queens", 4, false},
		{"e8", "Guardians uplink check", "4:00 PM", "30 min", "orange", "Knowhere uplink", 4, false},
	},
}
