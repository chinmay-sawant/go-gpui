package page

import "testing"

type bindData struct {
	Text  string
	Flag  bool
	Pick  string
	Count int
	inner string
}

func bindCtl(tag, kind string) Control {
	return Control{ID: "x", Tag: tag, Type: kind}
}

func TestBindText(t *testing.T) {
	t.Parallel()

	for _, c := range []Control{
		bindCtl("input", "text"),
		bindCtl("textarea", ""),
		bindCtl("select", ""),
	} {
		c.Bind, c.Value = "Text", "old"
		d := &bindData{Text: "Ann"}
		got, ok := bindRead(&Page{data: d}, c)
		if !ok || got.Value != "Ann" {
			t.Fatalf("%s read %q %v", c.Tag, got.Value, ok)
		}
		c.Value = "Bob"
		if !bindWrite(&Page{data: d}, c) || d.Text != "Bob" {
			t.Fatalf("%s write %q", c.Tag, d.Text)
		}
	}
}

func TestBindCheckbox(t *testing.T) {
	t.Parallel()

	c := bindCtl("input", "checkbox")
	c.Bind = "Flag"
	d := &bindData{Flag: true}
	got, ok := bindRead(&Page{data: d}, c)
	if !ok || !got.Checked {
		t.Fatalf("read %v %v", got.Checked, ok)
	}
	if !bindWrite(&Page{data: d}, got) || !d.Flag {
		t.Fatalf("write %v", d.Flag)
	}
	c.Checked = false
	if !bindWrite(&Page{data: d}, c) || d.Flag {
		t.Fatalf("write %v", d.Flag)
	}
}

func TestBindRadio(t *testing.T) {
	t.Parallel()

	yes := bindCtl("input", "radio")
	yes.Bind, yes.Value = "Pick", "b"
	no := yes
	no.Value = "a"

	d := &bindData{Pick: "b"}
	got, ok := bindRead(&Page{data: d}, yes)
	if !ok || !got.Checked {
		t.Fatalf("b read %v %v", got.Checked, ok)
	}
	got, ok = bindRead(&Page{data: d}, no)
	if !ok || got.Checked {
		t.Fatalf("a read %v %v", got.Checked, ok)
	}

	no.Checked = true
	if !bindWrite(&Page{data: d}, no) || d.Pick != "a" {
		t.Fatalf("write %q", d.Pick)
	}
	if !bindWrite(&Page{data: d}, yes) || d.Pick != "a" {
		t.Fatalf("unchecked %q", d.Pick)
	}
}
