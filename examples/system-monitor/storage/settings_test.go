package storage

import (
	"testing"
)

// TestSettingsRoundTrip checks write, replace, list, delete, and the missing
// and empty-key cases.
func TestSettingsRoundTrip(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()

	if _, ok, err := st.Setting(ctx, "theme"); err != nil || ok {
		t.Fatalf("missing setting ok=%v err=%v", ok, err)
	}
	if err := st.SetSetting(ctx, "theme", "dark"); err != nil {
		t.Fatal(err)
	}
	if err := st.SetSetting(ctx, "theme", "light"); err != nil {
		t.Fatal(err)
	}
	if err := st.SetSetting(ctx, "mode", "live"); err != nil {
		t.Fatal(err)
	}

	got, ok, err := st.Setting(ctx, "theme")
	if err != nil || !ok || got != "light" {
		t.Fatalf("theme = %q ok=%v err=%v", got, ok, err)
	}

	all, err := st.Settings(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 2 || all[0].Key != "mode" || all[0].Updated.IsZero() {
		t.Fatalf("settings = %+v", all)
	}

	if err := st.DeleteSetting(ctx, "mode"); err != nil {
		t.Fatal(err)
	}
	if _, ok, err := st.Setting(ctx, "mode"); err != nil || ok {
		t.Fatalf("deleted setting ok=%v err=%v", ok, err)
	}

	if err := st.SetSetting(ctx, "", "x"); err == nil {
		t.Fatal("empty key accepted")
	}
	if _, _, err := st.Setting(ctx, ""); err == nil {
		t.Fatal("empty key read accepted")
	}
}

// TestViewsRoundTrip is in views_test.go.
