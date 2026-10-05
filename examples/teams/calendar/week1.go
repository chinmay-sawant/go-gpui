package calendar

// week1 is the sample week of September 21, 2026.
var week1 = week{
	label: "September 21 – September 25",
	today: -1,
	days: []Day{
		{"Mon 9/21", false},
		{"Tue 9/22", false},
		{"Wed 9/23", false},
		{"Thu 9/24", false},
		{"Fri 9/25", false},
	},
	events: []Event{
		{"e22", "Avengers daily standup", "9:00 AM", "15 min", "blue", "Avengers Tower 42", 0, false},
		{"e23", "Suit test with Shuri", "10:00 AM", "90 min", "orange", "Wakanda Lab 3", 1, false},
		{"e24", "1:1 with Pepper Potts", "2:00 PM", "30 min", "green", "Avengers Tower 42", 1, false},
		{"e25", "Guardians uplink window", "11:00 AM", "60 min", "blue", "Knowhere uplink", 2, false},
		{"e26", "Stark budget review - Q3", "8:30 AM", "60 min", "red", "Stark Industries R&D", 3, false},
		{"e27", "Interview - Riri Williams", "10:00 AM", "30 min", "violet", "Avengers Compound", 4, false},
	},
}
