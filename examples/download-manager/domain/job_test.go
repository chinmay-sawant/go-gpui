package domain

import (
	"errors"
	"testing"
)

// errorsAs is a tiny shim so state_test stays readable.
func errorsAs(err error, target **TransitionError) bool {
	return errors.As(err, target)
}

// TestRedactURL drops credentials from userinfo and query values.
func TestRedactURL(t *testing.T) {
	got := RedactURL("https://alice:secret@example.com/file.bin?token=abc123&name=x#frag")
	want := "https://%5Bredacted%5D@example.com/file.bin?name=%5Bredacted%5D&token=%5Bredacted%5D#%5Bredacted%5D"

	if got != want {
		t.Errorf("RedactURL = %q, want %q", got, want)
	}

	plain := RedactURL("https://example.com/a.bin")
	if plain != "https://example.com/a.bin" {
		t.Errorf("plain URL changed: %q", plain)
	}

	if RedactURL("://bad") != redacted {
		t.Errorf("bad URL not redacted")
	}
}

// TestRedactHeader blanks the value after the first colon.
func TestRedactHeader(t *testing.T) {
	if got := Redact("Authorization: Bearer abc"); got != "Authorization: [redacted]" {
		t.Errorf("Redact = %q", got)
	}

	if got := Redact("nocolon"); got != redacted {
		t.Errorf("Redact = %q", got)
	}
}

// TestFraction covers known, unknown, and over-full lengths.
func TestFraction(t *testing.T) {
	cases := []struct {
		job  Job
		want float64
	}{
		{Job{Done: 50, Expected: 100}, 0.5},
		{Job{Done: 150, Expected: 100}, 1},
		{Job{Done: 10, Expected: -1}, -1},
		{Job{Done: 10, Total: 20}, 0.5},
		{Job{Done: -1, Expected: 100}, -1},
	}

	for _, c := range cases {
		if got := c.job.Fraction(); got != c.want {
			t.Errorf("Fraction(%+v) = %v, want %v", c.job, got, c.want)
		}
	}

	if !(Job{Expected: -1, Total: -1}).UnknownLength() {
		t.Error("unknown length not detected")
	}

	if (Job{Expected: 5}).UnknownLength() {
		t.Error("known length reported unknown")
	}
}
