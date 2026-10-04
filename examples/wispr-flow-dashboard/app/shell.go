package app

// NavItem is one sidebar link.
type NavItem struct {
	Label  string
	Icon   string
	Page   string
	Action string
}

// navItems is the sidebar's main link list.
func navItems() []NavItem {
	return []NavItem{
		{"Dictation", "icon-mic", "dictation", "nav-dictation"},
		{"Notetaker", "icon-note", "notetaker", "nav-notetaker"},
		{"Insights", "icon-chart", "insights", "nav-insights"},
		{"Dictionary", "icon-book", "dictionary", "nav-dictionary"},
		{"Snippets", "icon-scissors", "snippets", "nav-snippets"},
		{"Style", "icon-text", "style", "nav-style"},
		{"Transforms", "icon-wand", "transforms", "nav-transforms"},
		{"Scratchpad", "icon-scratch", "scratchpad", "nav-scratchpad"},
	}
}

// footNavItems is the sidebar's lower link list.
func footNavItems() []NavItem {
	return []NavItem{
		{"Invite your team", "icon-invite", "invite", "nav-invite"},
		{"Get a free month", "icon-gift", "free-month", "nav-free-month"},
		{"Settings", "icon-gear", "settings", "nav-settings"},
		{"Help", "icon-help", "help", "nav-help"},
	}
}
