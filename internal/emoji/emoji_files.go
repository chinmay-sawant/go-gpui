package emoji

import "embed"

//go:embed png/*.png
var pngs embed.FS

// PNG returns the bundled bytes for a file stem such as "1f602".
func PNG(name string) ([]byte, bool) {
	data, err := pngs.ReadFile("png/" + name + ".png")
	if err != nil {
		return nil, false
	}

	return data, true
}

// Files lists every bundled file stem.
func Files() []string {
	out := make([]string, 0, 32)
	seen := map[string]bool{}

	for _, e := range all() {
		if seen[e.file] {
			continue
		}

		seen[e.file] = true
		out = append(out, e.file)
	}

	return out
}
