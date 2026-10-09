package remote

type accessibleNode struct {
	ID, Name, Value string
	X, Y, W, H      float64
	Enabled         bool
	Editable        bool
	Selected        bool
	Button          bool
	Slider          bool
	Progress        int
}

type accessState struct {
	gen  uint64
	x, y int
}

func (a *App) controlNames() map[string]Key {
	keys := map[string]Key{
		"theme":       {Name: "Switch to " + a.view.ThemeLabel + " theme"},
		"connect":     {Name: "Connect to TV"},
		"scan":        {Name: "Scan Wi-Fi for TVs"},
		"mode":        {Name: "Connection mode: " + a.view.Mode},
		"tab-remote":  {Name: "Remote"},
		"tab-pad":     {Name: "Pad"},
		"tab-nums":    {Name: "Numbers and more controls"},
		"pad":         {Name: "TV pointer pad. Tap to click. Drag to move.", Disabled: a.view.Mode != "Wi-Fi"},
		"host":        {Name: "TV IP address"},
		"sensitivity": {Name: "Trackpad sensitivity"},
	}
	for _, row := range append(a.view.RemoteRows, a.view.NumberRows...) {
		for _, key := range row.Keys {
			keys[key.ID] = key
		}
	}
	for _, key := range a.view.PadKeys {
		keys[key.ID] = key
	}
	return keys
}
