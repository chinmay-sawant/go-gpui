package tv

import (
	"errors"
	"strings"
)

// Exec runs one spec such as "button:HOME" or "ssap:audio/volumeUp".
func (c *Client) Exec(spec string) (string, error) {
	kind, name, _ := strings.Cut(spec, ":")

	switch kind {
	case "button":
		return name, c.Button(name)
	case "ssap":
		_, err := c.Request(name, nil)
		return name, err
	case "launch":
		_, err := c.Request(
			"com.webos.applicationManager/launch",
			map[string]any{"id": name},
		)
		return "Launched", err
	case "input":
		_, err := c.Request("tv/switchInput", map[string]any{"inputId": name})
		return name, err
	case "power":
		_, err := c.Request("system/turnOff", nil)
		return "Turning the TV off", err
	case "pair":
		if c.Model != "" {
			return c.Model, nil
		}
		return "Paired", nil
	case "mute":
		return c.toggleMute()
	case "vol":
		return c.volume(name)
	case "ch":
		return c.channel(name)
	case "media":
		return c.media(name)
	case "screen":
		return c.screen(name)
	default:
		return "", errors.New("unknown control")
	}
}
