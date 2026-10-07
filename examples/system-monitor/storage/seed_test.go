package storage

import (
	"testing"
)

// fixtureSamples and fixture are in seed_fixture_test.go.

// TestSeedIdempotent checks that a second seed of the same version writes
// nothing and that a version bump replaces only the fixture session.
func TestSeedIdempotent(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()

	wrote, err := st.Seed(ctx, fixture(1))
	if err != nil || !wrote {
		t.Fatalf("seed wrote=%v err=%v", wrote, err)
	}

	sessions, err := st.Sessions(ctx, 10)
	if err != nil || len(sessions) != 1 {
		t.Fatalf("sessions = %+v err=%v", sessions, err)
	}
	if sessions[0].Source != "dummy" || sessions[0].State != StateDone || sessions[0].Rows == 0 {
		t.Fatalf("fixture session = %+v", sessions[0])
	}

	if err := st.SetSetting(ctx, "theme", "dark"); err != nil {
		t.Fatal(err)
	}
	userID := newTestSession(t, st)

	wrote, err = st.Seed(ctx, fixture(1))
	if err != nil || wrote {
		t.Fatalf("second seed wrote=%v err=%v", wrote, err)
	}

	wrote, err = st.Seed(ctx, fixture(2))
	if err != nil || !wrote {
		t.Fatalf("version bump wrote=%v err=%v", wrote, err)
	}

	sessions, err = st.Sessions(ctx, 10)
	if err != nil || len(sessions) != 2 {
		t.Fatalf("sessions = %+v err=%v", sessions, err)
	}
	if _, ok, err := st.Session(ctx, userID); err != nil || !ok {
		t.Fatalf("user session gone: ok=%v err=%v", ok, err)
	}
	if got, _, err := st.Setting(ctx, "theme"); err != nil || got != "dark" {
		t.Fatalf("setting = %q err=%v", got, err)
	}

	version, err := st.FixtureVersion(ctx)
	if err != nil || version != 2 {
		t.Fatalf("version = %d err=%v", version, err)
	}
}

// TestSeedRejectsBadVersion checks the version guard.
func TestSeedRejectsBadVersion(t *testing.T) {
	st := openTest(t)

	if _, err := st.Seed(t.Context(), fixture(0)); err == nil {
		t.Fatal("version zero accepted")
	}
}
