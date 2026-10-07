package ui

import "testing"

func TestFormulaBarStartsEdit(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	var box Box
	for _, b := range app.View().Toolbar {
		if string(b.Action) == "formula" && b.Class == "ftext" {
			box = b
		}
	}

	if box.W == 0 {
		t.Fatal("formula bar missing")
	}

	clickBox(t, app, box)
	if app.edit == nil || app.edit.Text != "name" {
		t.Fatalf("formula bar edit = %+v", app.edit)
	}
}

func TestSheetTabSwitches(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	for _, b := range app.View().Tabs {
		if string(b.Action) == "sheet:s2" {
			clickBox(t, app, b)
		}
	}

	settle(t, app)
	if app.active != "s2" {
		t.Fatalf("active = %q", app.active)
	}

	if app.View().Ref != "A1" {
		t.Fatalf("ref = %q", app.View().Ref)
	}
}

func TestSelectionSurvivesThemeToggle(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	app.selectRect(1, 1, 2, 2)
	settle(t, app)

	clickAction(t, app, "theme")
	flush(t, app)

	if s := app.selection(); s != (Selection{1, 1, 2, 2}) {
		t.Fatalf("selection after theme = %+v", s)
	}

	if !app.dark {
		t.Fatal("theme did not toggle")
	}
}
