package store

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"
	"time"
)

func TestSaveGameRetriesALockedDatabaseThenSucceeds(t *testing.T) {
	dir := t.TempDir()

	s, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	raw, err := sql.Open("sqlite",
		"file:"+slashAbs(filepath.Join(dir, DBName))+"?_pragma=busy_timeout(100)")
	if err != nil {
		t.Fatal(err)
	}
	defer raw.Close()

	raw.SetMaxOpenConns(1)

	tx, err := raw.Begin()
	if err != nil {
		t.Fatal(err)
	}

	if _, err := tx.Exec(`UPDATE scores SET score = score WHERE id = 'dummy-0001'`); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	done := make(chan error, 1)

	go func() { done <- s.SaveGame(ctx, testResult("locked", 1), nil) }()

	time.Sleep(250 * time.Millisecond)

	if err := tx.Rollback(); err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("save after unlock: %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("save never finished after the lock cleared")
	}
}

func TestSaveGameGivesUpWhenTheLockHolds(t *testing.T) {
	dir := t.TempDir()

	s, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	raw, err := sql.Open("sqlite",
		"file:"+slashAbs(filepath.Join(dir, DBName))+"?_pragma=busy_timeout(100)")
	if err != nil {
		t.Fatal(err)
	}
	defer raw.Close()

	raw.SetMaxOpenConns(1)

	tx, err := raw.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`UPDATE scores SET score = score WHERE id = 'dummy-0001'`); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	start := time.Now()

	if err := s.SaveGame(ctx, testResult("never", 1), nil); err == nil {
		t.Fatal("a save under a held lock succeeded")
	}

	if time.Since(start) > 8*time.Second {
		t.Fatal("the save did not respect its deadline")
	}
}
