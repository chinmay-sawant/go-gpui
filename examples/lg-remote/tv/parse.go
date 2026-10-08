package tv

import (
	"encoding/json"
	"strings"
)

func mapOf(raw json.RawMessage) map[string]any {
	out := map[string]any{}
	if len(raw) == 0 {
		return out
	}

	_ = json.Unmarshal(raw, &out)

	return out
}

func modelOf(raw json.RawMessage) string {
	p := mapOf(raw)
	s, _ := p["modelName"].(string)

	return s
}

func clientKey(raw json.RawMessage) string {
	s, _ := mapOf(raw)["client-key"].(string)

	return s
}

func macsOf(p map[string]any) []string {
	out := []string{}

	for _, name := range []string{"wifiInfo", "wiredInfo"} {
		info, _ := p[name].(map[string]any)
		if info == nil {
			continue
		}

		mac, _ := info["macAddress"].(string)
		mac = strings.TrimSpace(mac)
		if mac == "" || mac == "00:00:00:00:00:00" {
			continue
		}

		out = append(out, mac)
	}

	return out
}

func volumeText(p map[string]any) string {
	if p == nil {
		return "Volume"
	}

	if v, ok := p["volume"]; ok {
		return "Volume " + stringify(v)
	}

	status, _ := p["volumeStatus"].(map[string]any)
	if status == nil {
		return "Volume"
	}

	if v, ok := status["volume"]; ok {
		return "Volume " + stringify(v)
	}

	return "Volume"
}

func mutedOf(p map[string]any) (bool, bool) {
	if v, ok := p["muted"].(bool); ok {
		return v, true
	}

	status, _ := p["volumeStatus"].(map[string]any)
	if status == nil {
		return false, false
	}

	v, ok := status["muteStatus"].(bool)

	return v, ok
}

func stringify(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return strings.Trim(string(b), `"`)
}
