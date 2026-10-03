//go:build (!linux && !windows && !darwin) || android || ios

package filepick

import "context"

// Pick returns ok false where the system has no file dialog.
func Pick(_ context.Context, _ string) (string, bool) {
	return "", false
}
