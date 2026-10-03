package music

import "strings"

// licenseName formats an Openverse license for the credit line.
func licenseName(license, version string) string {
	name := strings.ToLower(strings.TrimSpace(license))
	if name == "" {
		return ""
	}

	if name == "cc0" {
		return "CC0 1.0"
	}

	out := "CC " + strings.ToUpper(name)
	if version != "" {
		out += " " + version
	}

	return out
}
