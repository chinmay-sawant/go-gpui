package tv

// registerMsg is the prompt pairing handshake used by current webOS TVs.
// A saved key skips the on-screen prompt.
func registerMsg(key string) map[string]any {
	payload := map[string]any{
		"forcePairing": false,
		"pairingType":  "PROMPT",
		"manifest": map[string]any{
			"appVersion":      "1.1",
			"manifestVersion": 1,
			"permissions":     permissions(),
		},
	}
	if key != "" {
		payload["client-key"] = key
	}

	return map[string]any{
		"type":    "register",
		"id":      "register_0",
		"payload": payload,
	}
}
