package remote

import "strings"

const remoteRows = `power wake
volup up chup
left ok right
voldn down chdn
mute back input
home exit guide
info menu
rew play pause stop ff
red green yellow blue
hotstar prime youtube`

const numberRows = `n1 n2 n3
n4 n5 n6
n7 n8 n9
star n0 dash
list livetv qmenu
hdmi1 hdmi2 hdmi3
cc sap ad
record aspect threed
apps recent zoom
text magnify program
screenoff screenon`

func rows(spec string, keys map[string]Key) []Row {
	var out []Row
	for _, line := range strings.Split(spec, "\n") {
		row := Row{}
		for _, id := range strings.Fields(line) {
			key := keys[id]
			key.Class = strings.ReplaceAll(key.Class, "wide", "")
			row.Keys = append(row.Keys, key)
		}
		if strings.HasPrefix(line, "left") || strings.HasPrefix(line, "rew") {
			row.Class = "midrow"
		}
		out = append(out, row)
	}
	return out
}
