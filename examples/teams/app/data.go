package app

// RailItem is one app rail link.
type RailItem struct {
	Label   string
	Icon    string
	Section string
	Action  string
	Badge   string
}

// AppTile is one dummy app in the More apps flyout.
type AppTile struct {
	ID     string
	Name   string
	Letter string
	Color  string
}

// railSections maps a rail action to the menu it opens. The Teams menu is
// section "channels".
var railSections = map[string]string{
	"rail-activity": "activity",
	"rail-chat":     "chat",
	"rail-channels": "channels",
	"rail-calendar": "calendar",
	"rail-calls":    "calls",
	"rail-files":    "files",
}

// railItems is the app rail in order.
func railItems() []RailItem {
	return []RailItem{
		{"Activity", "rail-activity", "activity", "rail-activity", "3"},
		{"Chat", "rail-chat", "chat", "rail-chat", "2"},
		{"Teams", "rail-teams", "channels", "rail-channels", ""},
		{"Calendar", "rail-calendar", "calendar", "rail-calendar", ""},
		{"Calls", "rail-calls", "calls", "rail-calls", ""},
		{"Files", "rail-files", "files", "rail-files", ""},
	}
}

// appTiles are the dummy apps the More flyout shows.
func appTiles() []AppTile {
	return []AppTile{
		{"word", "Word", "W", "#2b579a"},
		{"excel", "Excel", "X", "#217346"},
		{"powerpoint", "PowerPoint", "P", "#d24726"},
		{"onenote", "OneNote", "N", "#7719aa"},
		{"outlook", "Outlook", "O", "#0f6cbd"},
		{"planner", "Planner", "P", "#31752f"},
		{"lists", "Lists", "L", "#0078d4"},
		{"shifts", "Shifts", "S", "#0b6a0b"},
		{"loop", "Loop", "L", "#7a1fa2"},
		{"viva", "Viva Engage", "V", "#a4262c"},
		{"forms", "Forms", "F", "#0f7b6c"},
		{"bookings", "Bookings", "B", "#8764b8"},
	}
}
