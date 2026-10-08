package remote

import "strconv"

// btCommand is the HID report the phone should send.
// key is a keyboard usage. con is a consumer-control usage.
// Controls that have no HID usage return false and stay on Wi-Fi.
func btCommand(id string) (string, bool) {
	if code, ok := btKey[id]; ok {
		return "key:" + strconv.Itoa(code), true
	}

	if code, ok := btCon[id]; ok {
		return "con:" + strconv.Itoa(code), true
	}

	return "", false
}

var btKey = map[string]int{
	"up":    0x52,
	"down":  0x51,
	"left":  0x50,
	"right": 0x4F,
	"ok":    0x28,
	"back":  0x29,
	"exit":  0x29,
	"n1":    0x1E,
	"n2":    0x1F,
	"n3":    0x20,
	"n4":    0x21,
	"n5":    0x22,
	"n6":    0x23,
	"n7":    0x24,
	"n8":    0x25,
	"n9":    0x26,
	"n0":    0x27,
}

var btCon = map[string]int{
	"volup":  0xE9,
	"voldn":  0xEA,
	"mute":   0xE2,
	"chup":   0x9C,
	"chdn":   0x9D,
	"play":   0xB0,
	"pause":  0xB1,
	"stop":   0xB7,
	"rew":    0xB4,
	"ff":     0xB3,
	"power":  0x30,
	"home":   0x223,
	"menu":   0x40,
	"record": 0xB2,
}
