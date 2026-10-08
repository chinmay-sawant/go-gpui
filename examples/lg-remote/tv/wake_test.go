package tv

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMagicPacket(t *testing.T) {
	pkt, err := Magic("aa:bb:cc:dd:ee:ff")
	if err != nil {
		t.Fatal(err)
	}

	if len(pkt) != 102 || pkt[0] != 0xff || pkt[6] != 0xaa || pkt[11] != 0xff {
		t.Fatalf("packet %d %x", len(pkt), pkt[:12])
	}
}

func TestRegisterCarriesKey(t *testing.T) {
	b, err := json.Marshal(registerMsg("abc"))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(b), `"client-key":"abc"`) {
		t.Fatal(string(b))
	}
}

func TestMacsAndVolume(t *testing.T) {
	p := map[string]any{
		"wifiInfo": map[string]any{"macAddress": "aa:bb:cc:dd:ee:ff"},
		"volume":   12.0,
	}
	macs := macsOf(p)
	if len(macs) != 1 || macs[0] != "aa:bb:cc:dd:ee:ff" {
		t.Fatal(macs)
	}

	if volumeText(p) != "Volume 12" {
		t.Fatal(volumeText(p))
	}
}

func TestStoreRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ownframe", "lg-remote.json")
	want := Store{Host: "10.0.0.4", Key: "k", Model: "UP7750PTZ", MACs: []string{"aa:bb:cc:dd:ee:ff"}}
	if err := Save(path, want); err != nil {
		t.Fatal(err)
	}

	got := Load(path)
	if got.Host != want.Host || got.Key != want.Key || got.Model != want.Model || len(got.MACs) != 1 {
		t.Fatalf("%+v", got)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}

func TestWakeTargets(t *testing.T) {
	got := wakeTargets("192.168.0.101")
	found := false
	for _, item := range got {
		if item == "192.168.0.255:9" {
			found = true
		}
	}
	if !found {
		t.Fatal(got)
	}
}

func TestCleanHost(t *testing.T) {
	if got := cleanHost("ws://10.0.0.5:3000/"); got != "10.0.0.5" {
		t.Fatal(got)
	}
}
