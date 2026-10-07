package formula

import "testing"

type testEnv map[[2]int]Value

func (e testEnv) Cell(row, col int) Value { return e[[2]int{row, col}] }

func evalStr(t *testing.T, src string, env Env) Value {
	t.Helper()

	e, err := Parse(src, DefaultLimits())
	if err != nil {
		t.Fatalf("Parse(%q): %v", src, err)
	}

	return Eval(e, env, DefaultLimits(), nil)
}

func TestArithmetic(t *testing.T) {
	env := testEnv{}

	cases := []struct {
		src  string
		want string
	}{
		{"1+2*3", "7"},
		{"(1+2)*3", "9"},
		{"2^3^2", "512"},
		{"2^-1", "0.5"},
		{"-3+1", "-2"},
		{"-2*3", "-6"},
		{"10/4", "2.5"},
		{"1-2-3", "-4"},
		{`"a"`, "a"},
		{`"say ""hi"""`, `say "hi"`},
		{"1e3", "1000"},
	}

	for _, c := range cases {
		if got := evalStr(t, c.src, env).Display(); got != c.want {
			t.Errorf("%s = %q want %q", c.src, got, c.want)
		}
	}
}

func TestErrors(t *testing.T) {
	env := testEnv{
		{0, 0}: Text("hi"),
		{0, 2}: Err(ErrDiv),
	}

	cases := []struct {
		src  string
		code string
	}{
		{"1/0", ErrDiv},
		{"FOO(1)", ErrName},
		{"A1+1", ErrValue},
		{"C1+0", ErrDiv},
		{"A1:A3", ErrValue},
	}

	for _, c := range cases {
		v := evalStr(t, c.src, env)
		if !v.IsError() || v.Code != c.code {
			t.Errorf("%s = %v want error %s", c.src, v, c.code)
		}
	}

	if got := evalStr(t, "A1", env).Display(); got != "hi" {
		t.Errorf("A1 = %q want hi", got)
	}
}
