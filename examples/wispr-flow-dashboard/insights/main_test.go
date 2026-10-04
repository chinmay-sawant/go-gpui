package insights

import (
	"os"
	"testing"

	osclip "github.com/chinmay-sawant/go-gpui/internal/clipboard"
)

// TestMain keeps every test off the desktop clipboard.
func TestMain(m *testing.M) {
	osclip.UseMemory(true)
	os.Exit(m.Run())
}
