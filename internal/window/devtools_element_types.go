package window

// devElementJSON is the picked element on the Elements tab.
type devElementJSON struct {
	Tag    string         `json:"tag"`
	ID     string         `json:"id"`
	Action string         `json:"action"`
	Text   string         `json:"text"`
	Rect   devElementRect `json:"rect"`
	Ops    int            `json:"ops"`
}

// devElementRect is an element border box in CSS pixels.
type devElementRect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
