package scene

import (
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestPadScoreBounds(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "000000000"},
		{42, "000000042"},
		{-5, "000000000"},
		{game.MaxScore, "999999999"},
		{game.MaxScore + 1, "999999999"},
	}

	for _, c := range cases {
		if got := padScore(c.in); got != c.want {
			t.Fatalf("padScore(%d) = %q, want %q", c.in, got, c.want)
		}

		if len(padScore(c.in)) != 9 {
			t.Fatalf("padScore(%d) is not nine digits", c.in)
		}
	}
}

func TestPadLinesAndLevel(t *testing.T) {
	if got := padLines(7); got != "0007" {
		t.Fatalf("padLines(7) = %q", got)
	}

	if got := padLevel(0); got != "01" {
		t.Fatalf("padLevel(0) = %q, want the clamped minimum", got)
	}

	if got := padLevel(game.MaxLevel + 5); got != "20" {
		t.Fatalf("padLevel over max = %q", got)
	}
}

func TestPageLabel(t *testing.T) {
	if got := pageLabel(0, false); got != "PAGE 1" {
		t.Fatalf("pageLabel(0) = %q", got)
	}

	if got := pageLabel(2, true); got != "PAGE 3+" {
		t.Fatalf("pageLabel(2, more) = %q", got)
	}
}

func TestEntryLine(t *testing.T) {
	live := ScoreEntry{Rank: 1, ID: "abcdef012345", Score: 1200, Lines: 12, Level: 2}
	got := entryLine(live)

	if !strings.HasPrefix(got, "    1") {
		t.Fatalf("row %q is not rank-aligned", got)
	}

	if !strings.Contains(got, "abcdef01") {
		t.Fatalf("row %q lost the short id", got)
	}

	if strings.Contains(got, "DEMO") {
		t.Fatalf("live row %q carries a demo tag", got)
	}

	demo := entryLine(ScoreEntry{Rank: 20, Dummy: true})
	if !strings.HasSuffix(demo, "DEMO") {
		t.Fatalf("demo row %q is not tagged", demo)
	}
}

func TestThemeLabel(t *testing.T) {
	if got := themeLabel(false); got != "THEME DARK" {
		t.Fatalf("light label = %q", got)
	}

	if got := themeLabel(true); got != "THEME LIGHT" {
		t.Fatalf("dark label = %q", got)
	}
}
