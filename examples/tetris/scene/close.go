package scene

import "time"

// ShutdownBudget bounds how long Close waits for the store worker.
const ShutdownBudget = 2 * time.Second

// Close stops the frame callback and joins the store worker within
// ShutdownBudget. A worker that outlives the budget is reported and left
// to the process exit.
func (s *Scene) Close() error {
	s.page.SetTick(nil)
	s.savePoint()

	if s.store == nil {
		return nil
	}

	done := make(chan error, 1)

	go func() { done <- s.store.Close() }()

	select {
	case err := <-done:
		return err
	case <-time.After(ShutdownBudget):
		return ErrShutdownTimeout
	}
}
