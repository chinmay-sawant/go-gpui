//go:build (!linux && !windows && !darwin) || android || ios

package print

import "context"

// Print returns ErrNoPrinter where the system has no print helper.
func Print(_ context.Context, _ string) error {
	return ErrNoPrinter
}
