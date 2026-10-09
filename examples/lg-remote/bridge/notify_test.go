package bridge

import "testing"

func TestCommandNotificationCanDrainImmediately(t *testing.T) {
	defer SetCommandNotifier(nil)
	for Take() != "" {
	}
	var got string
	SetCommandNotifier(func() { got = Take() })
	if !Enqueue("key:80") || got != "key:80" {
		t.Fatal("dispatcher did not receive accepted command")
	}
	SetCommandNotifier(nil)
	Enqueue("key:79")
	SetCommandNotifier(func() { got = Take() })
	if got != "key:79" {
		t.Fatal("resume did not drain pending command")
	}
}
