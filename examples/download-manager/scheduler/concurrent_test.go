package scheduler

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestConcurrentAddsReserveDistinctDestinations gives two adds the same
// requested name and checks that reservation serializes them.
func TestConcurrentAddsReserveDistinctDestinations(t *testing.T) {
	eng, _ := newTestEngine(t, instantTransport{}, Options{Workers: 1})
	dir := t.TempDir()
	ctx := context.Background()

	results := make([]domain.Job, 2)

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			job, err := eng.Add(ctx, AddRequest{
				URL:  fmt.Sprintf("https://example.invalid/%d.bin", i),
				Dir:  dir,
				Name: "same.bin",
			})
			if err != nil {
				t.Errorf("add %d: %v", i, err)

				return
			}

			results[i] = job
		}(i)
	}

	wg.Wait()

	if results[0].Destination == results[1].Destination {
		t.Errorf("both adds reserved %s", results[0].Destination)
	}
}
