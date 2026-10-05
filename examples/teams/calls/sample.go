package calls

// sampleHistory is the read-only call history every Default clones.
var sampleHistory = []Call{
	{"c1", "Shuri", "SH", "purple", "Outgoing call", "10:12 AM", "4m 32s", false, false},
	{"c2", "Pepper Potts", "PP", "pink", "Incoming call", "9:48 AM", "12m 05s", false, true},
	{"c3", "Nick Fury", "NF", "gray", "Missed call", "9:20 AM", "", true, false},
	{"c4", "Steve Rogers", "SR", "blue", "Outgoing call", "Yesterday", "1m 18s", false, false},
	{"c5", "Natasha Romanoff", "NR", "red", "Incoming call", "Yesterday", "8m 44s", false, true},
	{"c6", "Thor", "TH", "gold", "Missed call", "Yesterday", "", true, false},
	{"c7", "Bruce Banner", "BB", "green", "Outgoing call", "Monday", "2m 07s", false, false},
	{"c8", "Sam Wilson", "SW", "teal", "Missed call", "Monday", "", true, false},
}

// sampleVoicemails is the read-only voicemail every Default clones.
var sampleVoicemails = []Voicemail{
	{"v1", "Happy Hogan", "HH", "orange", "8:02 AM", "0:42", "The jet is fueled and ready. Wheels up at nine tomorrow."},
	{"v2", "Peter Parker", "PP", "red", "Yesterday", "0:18", "The field test went great. The new web shooters held at full swing."},
	{"v3", "Maria Hill", "MH", "violet", "Monday", "1:05", "Coordinates for the briefing: 40.7484 N, 73.9857 W. Rendezvous at oh-six-hundred."},
}
