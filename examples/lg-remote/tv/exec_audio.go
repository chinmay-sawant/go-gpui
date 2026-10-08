package tv

import "strings"

func (c *Client) channel(name string) (string, error) {
	_, err := c.Request("tv/channel"+name, nil)
	if err == nil {
		return "Channel", nil
	}

	key := "CHANNELUP"
	if name == "Down" {
		key = "CHANNELDOWN"
	}

	if err2 := c.Button(key); err2 != nil {
		return "", err
	}

	return "Channel", nil
}

func (c *Client) volume(name string) (string, error) {
	if _, err := c.Request("audio/volume"+name, nil); err != nil {
		key := "VOLUMEUP"
		if name == "Down" {
			key = "VOLUMEDOWN"
		}
		if err2 := c.Button(key); err2 != nil {
			return "", err
		}
	}

	p, err := c.Request("audio/getVolume", nil)
	if err != nil {
		return "Volume", nil
	}

	return volumeText(p), nil
}

func (c *Client) toggleMute() (string, error) {
	p, err := c.Request("audio/getVolume", nil)
	muted := false
	if err == nil {
		if v, ok := mutedOf(p); ok {
			muted = v
		}
	}

	_, err = c.Request("audio/setMute", map[string]any{"mute": !muted})
	if err != nil && c.Button("MUTE") != nil {
		return "", err
	}

	if muted {
		return "Sound on", nil
	}

	return "Muted", nil
}

func (c *Client) media(name string) (string, error) {
	uri := map[string]string{
		"PLAY": "play", "PAUSE": "pause", "STOP": "stop",
		"REWIND": "rewind", "FASTFORWARD": "fastForward",
	}[name]
	if uri == "" {
		uri = strings.ToLower(name)
	}

	_, errSSAP := c.Request("media.controls/"+uri, nil)
	errKey := c.Button(name)
	if errSSAP != nil && errKey != nil {
		return "", errKey
	}

	return name, nil
}

func (c *Client) screen(name string) (string, error) {
	uri := "com.webos.service.tvpower/power/turnOffScreen"
	label := "Screen off"
	if name == "on" {
		uri = "com.webos.service.tvpower/power/turnOnScreen"
		label = "Screen on"
	}

	_, err := c.Request(uri, nil)

	return label, err
}
