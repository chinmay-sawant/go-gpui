package tv

import "testing"

func TestTagText(t *testing.T) {
	body := "<friendlyName>[LG] webOS TV UP7750PTZ</friendlyName>"
	if got := tagText(body, "friendlyName"); got != "[LG] webOS TV UP7750PTZ" {
		t.Fatal(got)
	}
}

func TestChoosePrefersTheSet(t *testing.T) {
	hits := []Hit{
		{IP: "192.168.0.1", Name: "router"},
		{IP: "192.168.0.101", Name: "[LG] webOS TV UP7750PTZ"},
	}
	got := choose(hits)
	if got.IP != "192.168.0.101" {
		t.Fatal(got.IP)
	}
}
