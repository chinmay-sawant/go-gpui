package remote

func actionOf(id string) (string, bool) {
	if spec, ok := faceAct[id]; ok {
		return spec, true
	}

	spec, ok := moreAct[id]

	return spec, ok
}

var faceAct = map[string]string{
	"power":   "power:",
	"voldn":   "vol:Down",
	"volup":   "vol:Up",
	"mute":    "mute:",
	"chdn":    "ch:Down",
	"chup":    "ch:Up",
	"up":      "button:UP",
	"left":    "button:LEFT",
	"ok":      "button:ENTER",
	"right":   "button:RIGHT",
	"down":    "button:DOWN",
	"back":    "button:BACK",
	"home":    "button:HOME",
	"exit":    "button:EXIT",
	"input":   "hub:",
	"n1":      "button:1",
	"n2":      "button:2",
	"n3":      "button:3",
	"n4":      "button:4",
	"n5":      "button:5",
	"n6":      "button:6",
	"n7":      "button:7",
	"n8":      "button:8",
	"n9":      "button:9",
	"n0":      "button:0",
	"play":    "media:PLAY",
	"pause":   "media:PAUSE",
	"stop":    "media:STOP",
	"rew":     "media:REWIND",
	"ff":      "media:FASTFORWARD",
	"red":     "button:RED",
	"green":   "button:GREEN",
	"yellow":  "button:YELLOW",
	"blue":    "button:BLUE",
	"hotstar": "hotstar:",
	"prime":   "app:AMAZON|amazon",
	"youtube": "launch:youtube.leanback.v4",
	"guide":   "button:GUIDE",
	"info":    "button:INFO",
	"menu":    "button:MENU",
}
