package formula

import "testing"

func TestRanges(t *testing.T) {
	env := testEnv{
		{0, 0}: Number(1),
		{1, 0}: Number(2),
		{2, 0}: Text("skip"),
		{0, 1}: Number(10),
	}

	cases := []struct {
		src  string
		want string
	}{
		{"SUM(A1:A4)", "3"},
		{"SUM(A1:B1)", "11"},
		{"SUM(A1:A4,10)", "13"},
		{"SUM(A1,A2)", "3"},
		{"AVERAGE(A1:A4)", "1.5"},
		{"SUM(B1:B2)", "10"},
		{"AVERAGE(B1:B2)", "10"},
	}

	for _, c := range cases {
		if got := evalStr(t, c.src, env).Display(); got != c.want {
			t.Errorf("%s = %q want %q", c.src, got, c.want)
		}
	}

	if v := evalStr(t, "AVERAGE(A3)", env); v.Code != ErrDiv {
		t.Errorf("AVERAGE(A3) = %v want %s", v, ErrDiv)
	}
}

func TestRefs(t *testing.T) {
	e, err := Parse("SUM(A1:B2)+C3", DefaultLimits())
	if err != nil {
		t.Fatal(err)
	}

	want := []Rect{
		{MinRow: 0, MinCol: 0, MaxRow: 1, MaxCol: 1},
		{MinRow: 2, MinCol: 2, MaxRow: 2, MaxCol: 2},
	}

	got := Refs(e)
	if len(got) != len(want) {
		t.Fatalf("Refs = %v want %v", got, want)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Refs[%d] = %v want %v", i, got[i], want[i])
		}
	}
}
