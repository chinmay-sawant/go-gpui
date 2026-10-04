//go:build windows

package print

import "context"

// Print asks PowerShell to run the shell print verb on path.
func Print(ctx context.Context, path string) error {
	return run(ctx, path, windowsCommands)
}
