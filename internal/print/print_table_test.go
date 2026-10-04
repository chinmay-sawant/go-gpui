package print

import (
	"context"
	"strings"
	"testing"
)

// TestCommandsPerGOOS pins the program and argument shape of every system
// before any fake runner is installed.
func TestCommandsPerGOOS(t *testing.T) {
	cases := []struct {
		name     string
		commands []command
		present  string
		want     string
		contains string
	}{
		{"linux lp", linuxCommands, "lp", "lp", "/tmp/r.pdf"},
		{"linux xdg-open", linuxCommands, "xdg-open", "xdg-open", "/tmp/r.pdf"},
		{"darwin osascript", darwinCommands, "osascript", "osascript", "POSIX file"},
		{"darwin open", darwinCommands, "open", "open", "/tmp/r.pdf"},
		{"windows powershell", windowsCommands, "powershell.exe", "powershell.exe", "-Verb Print"},
		{"windows pwsh", windowsCommands, "pwsh.exe", "pwsh.exe", "-Verb Print"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := fakeRunner(t, map[string]bool{tc.present: true})

			if err := run(context.Background(), "/tmp/r.pdf", tc.commands); err != nil {
				t.Fatal(err)
			}

			got := *calls
			if len(got) != 1 || got[0].name != "/bin/"+tc.want {
				t.Fatalf("calls = %+v", got)
			}

			joined := strings.Join(got[0].args, " ")
			if !strings.Contains(joined, tc.contains) {
				t.Fatalf("args = %q, want %q", joined, tc.contains)
			}
		})
	}
}
