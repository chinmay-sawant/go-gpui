package page

import (
	"reflect"
	"testing"
)

func TestBindInert(t *testing.T) {
	t.Parallel()

	c := Control{Type: "text", Bind: "Text", Value: "x"}
	d := &bindData{}
	for _, data := range []any{nil, bindData{}, (*bindData)(nil), new(int)} {
		if _, ok := bindRead(&Page{data: data}, c); ok {
			t.Fatalf("read %T", data)
		}
		if bindWrite(&Page{data: data}, c) || d.Text != "" {
			t.Fatalf("write %T", data)
		}
	}
	if _, ok := bindRead(nil, c); ok {
		t.Fatal("nil page")
	}
	if bindWrite(nil, c) {
		t.Fatal("nil page write")
	}
}

func TestBindIgnored(t *testing.T) {
	t.Parallel()

	d := &bindData{Text: "keep", Count: 3, inner: "hidden"}
	for _, c := range []Control{
		{Type: "text", Bind: "Nope"},
		{Type: "text", Bind: "inner"},
		{Type: "text", Bind: "Count"},
		{Type: "text"},
		{Type: "checkbox", Bind: "Text"},
		{Type: "checkbox", Bind: "Count"},
		{Tag: "textarea", Bind: "Count"},
		{Type: "radio", Bind: "Count"},
		{Type: "submit", Bind: "Text"},
	} {
		if _, ok := bindRead(&Page{data: d}, c); ok {
			t.Fatalf("read %+v", c)
		}
		if bindWrite(&Page{data: d}, c) {
			t.Fatalf("write %+v", c)
		}
	}
	if d.Text != "keep" || d.Count != 3 || d.inner != "hidden" {
		t.Fatalf("mutated %+v", d)
	}
}

func TestBindTarget(t *testing.T) {
	t.Parallel()

	f, ok := bindTarget(&Page{data: &bindData{Text: "hi"}}, "Text")
	if !ok || f.Kind() != reflect.String || f.String() != "hi" {
		t.Fatalf("target %v %v", f, ok)
	}
	for _, name := range []string{"", "Nope", "inner"} {
		if _, ok := bindTarget(&Page{data: &bindData{}}, name); ok {
			t.Fatalf("target %q", name)
		}
	}
}
