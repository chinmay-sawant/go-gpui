//go:build darwin && !ios

package print

import "context"

// Print asks Preview to print path, or opens it when osascript is missing.
func Print(ctx context.Context, path string) error {
	return run(ctx, path, darwinCommands)
}
