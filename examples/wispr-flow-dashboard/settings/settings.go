// Package settings is the settings page: the settings column and the General panel.
package settings

// Data is the data the settings page prints.
type Data struct {
	Tabs      []SettingTab
	Languages []SettingOption
	Toggles   []SettingToggle
	Note      string
}

// SettingTab is one item in the settings side navigation.
type SettingTab struct {
	ID     string
	Label  string
	Active bool
}

// SettingOption is one language choice.
type SettingOption struct {
	Value    string
	Label    string
	Selected bool
}

// SettingToggle is one checkbox row in the General panel.
type SettingToggle struct {
	ID      string
	Label   string
	Checked bool
}

// Default returns the settings page data.
func Default() Data {
	return Data{
		Tabs: []SettingTab{
			{"set-tab-general", "General", true},
			{"set-tab-dictation", "Dictation", false},
			{"set-tab-microphone", "Microphone", false},
			{"set-tab-appearance", "Appearance", false},
			{"set-tab-shortcuts", "Shortcuts", false},
			{"set-tab-account", "Account", false},
			{"set-tab-privacy", "Privacy", false},
			{"set-tab-about", "About", false},
		},
		Languages: []SettingOption{
			{"en-in", "English (India)", true},
			{"en-us", "English (US)", false},
			{"de", "Deutsch", false},
			{"ja", "日本語", false},
		},
		Toggles: []SettingToggle{
			{"set-login", "Launch Flow at login", true},
			{"set-sound", "Play a sound when dictation starts", true},
			{"set-format", "Auto-format punctuation", true},
			{"set-space", "Insert trailing space after dictation", false},
		},
		Note: "Changes apply to this device only.",
	}
}
