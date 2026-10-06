//go:build linux

package collector

import "testing"

// statFixture has spaces and parentheses in the process name.
const statFixture = `42 (weird (name) here) S 1 42 42 0 -1 4194560 100 0 0 0 10 5 0 0 20 0 8 0 12345 1000000 500 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
`

// TestParseProcStat checks the fields after the last parenthesis and the comm
// that holds spaces.
func TestParseProcStat(t *testing.T) {
	st, err := parseProcStat([]byte(statFixture))
	if err != nil {
		t.Fatal(err)
	}

	if st.comm != "weird (name) here" {
		t.Fatalf("comm = %q", st.comm)
	}
	if st.state != "S" || st.ppid != 1 {
		t.Fatalf("state=%q ppid=%d", st.state, st.ppid)
	}
	if st.utime != 10 || st.stime != 5 {
		t.Fatalf("utime=%d stime=%d", st.utime, st.stime)
	}
	if st.nice != 0 || st.threads != 8 {
		t.Fatalf("nice=%d threads=%d", st.nice, st.threads)
	}
	if st.startTicks != 12345 || st.vsize != 1000000 || st.rss != 500 {
		t.Fatalf("start=%d vsize=%d rss=%d", st.startTicks, st.vsize, st.rss)
	}
}

// TestParseProcStatBad checks truncated and malformed lines.
func TestParseProcStatBad(t *testing.T) {
	for _, bad := range []string{"", "42 no parens", "42 (short) S 1"} {
		if _, err := parseProcStat([]byte(bad)); err == nil {
			t.Fatalf("accepted %q", bad)
		}
	}
}

// TestStateName checks the state letters this example maps.
func TestStateName(t *testing.T) {
	cases := map[string]string{
		"R": "running",
		"S": "sleeping",
		"D": "disk sleep",
		"Z": "zombie",
		"?": "unknown",
	}
	for in, want := range cases {
		if got := stateName(in); got != want {
			t.Fatalf("stateName(%q) = %q, want %q", in, got, want)
		}
	}
}
