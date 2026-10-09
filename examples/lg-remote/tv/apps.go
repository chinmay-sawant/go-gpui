package tv

func (c *Client) hub() (string, error) {
	_, errPick := c.Request("com.webos.surfacemanager/showInputPicker", nil)
	errKey := c.Button("INPUT_HUB")
	if errPick != nil && errKey != nil {
		return "", errKey
	}

	return "Input", nil
}

func (c *Client) hdmi(n string) (string, error) {
	_, err := c.Request("tv/switchInput", map[string]any{"inputId": "HDMI_" + n})
	if err == nil {
		return "HDMI " + n, nil
	}

	id := "com.webos.app.hdmi" + n
	_, err2 := c.Request("com.webos.applicationManager/launch", map[string]any{"id": id})
	if err2 == nil {
		return "HDMI " + n, nil
	}

	return "", err
}

func (c *Client) appButton(spec string) (string, error) {
	button, id, _ := cutApp(spec)
	if c.Button(button) == nil {
		return button, nil
	}

	return c.launch(id)
}

func cutApp(spec string) (string, string, bool) {
	for i := 0; i < len(spec); i++ {
		if spec[i] == '|' {
			return spec[:i], spec[i+1:], true
		}
	}

	return spec, spec, false
}

func (c *Client) launch(id string) (string, error) {
	ids := []string{id}
	if id == "youtube.leanback.v4" {
		ids = []string{"youtube.leanback.v4", "youtube.leanback"}
	}

	var last error
	for _, one := range ids {
		_, err := c.Request("com.webos.applicationManager/launch", map[string]any{"id": one})
		if err == nil {
			return "Launched", nil
		}

		last = err
	}

	return "", last
}
