//go:build darwin && !ios

package filepick

import (
	"context"
	"os/exec"
	"strings"
)

const script = `on run argv
POSIX path of (choose file with prompt (item 1 of argv))
end run`

// Pick shows the macOS open dialog through osascript.
// ok is false when the person canceled.
func Pick(ctx context.Context, title string) (string, bool) {
	out, err := exec.CommandContext(ctx, "osascript", "-e", script, title).Output()
	if err != nil {
		return "", false
	}

	picked := strings.TrimRight(string(out), "\r\n")
	if picked == "" {
		return "", false
	}

	return picked, true
}
