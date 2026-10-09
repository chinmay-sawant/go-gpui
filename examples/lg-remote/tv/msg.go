package tv

import "encoding/json"

// Message is one JSON text frame from the TV.
type Message struct {
	ID      flexID          `json:"id"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
	Error   string          `json:"error"`
}

type flexID string

func (f *flexID) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		*f = ""
		return nil
	}

	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}

		*f = flexID(s)

		return nil
	}

	*f = flexID(b)

	return nil
}

func helloMsg() map[string]any {
	return map[string]any{
		"id":      "hello",
		"type":    "hello",
		"payload": map[string]any{},
	}
}

func sysMsg() map[string]any {
	return map[string]any{
		"id":      "sys",
		"type":    "request",
		"uri":     "ssap://system/getSystemInfo",
		"payload": map[string]any{},
	}
}
