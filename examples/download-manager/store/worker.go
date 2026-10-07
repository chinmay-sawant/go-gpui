package store

import (
	"context"
)

// Close stops the worker and closes the connection. It is safe twice.
func (s *Store) Close() error {
	s.once.Do(func() {
		close(s.quit)
		<-s.done
	})

	return s.db.Close()
}

// JournalMode reports the active journal mode, "wal", "delete", or
// "memory".
func (s *Store) JournalMode() string { return s.journal }

// Dir reports the data directory, empty for temporary mode.
func (s *Store) Dir() string { return s.dir }

// Memory reports whether the store lives only in RAM.
func (s *Store) Memory() bool { return s.memory }

// do runs fn on the single worker and waits for it. A caller that loses
// patience gets its context error; the worker still finishes the piece it
// started.
func (s *Store) do(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, queryTimeout)
		defer cancel()
	}

	reply := make(chan error, 1)
	op := func() { reply <- fn(ctx) }

	select {
	case s.ops <- op:
	case <-s.quit:
		return ErrClosed
	case <-ctx.Done():
		return ctx.Err()
	}

	select {
	case err := <-reply:
		return err
	case <-s.quit:
		return ErrClosed
	case <-ctx.Done():
		return ctx.Err()
	}
}

// loop runs queued operations until Close, then drains what is queued.
func (s *Store) loop() {
	defer close(s.done)

	for {
		select {
		case op := <-s.ops:
			op()
		case <-s.quit:
			for {
				select {
				case op := <-s.ops:
					op()
				default:
					return
				}
			}
		}
	}
}
