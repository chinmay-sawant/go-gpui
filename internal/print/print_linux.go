//go:build linux && !android

package print

import "context"

// Print hands path to lp, or opens it with xdg-open when lp is missing.
func Print(ctx context.Context, path string) error {
	return run(ctx, path, linuxCommands)
}
