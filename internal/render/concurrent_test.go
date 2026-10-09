package render_test

import (
	"context"
	"sync"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/render"
)

func TestConcurrentIndependentPages(t *testing.T) {
	var wg sync.WaitGroup
	for range 4 {
		wg.Go(func() {
			for range 3 {
				cache, err := render.NewCache(context.Background(), `<p>Hello</p>`, 200, 100, render.State{})
				if err != nil {
					t.Error(err)
					return
				}
				if _, err := render.DisplayListDocument(context.Background(), cache.Styled(), nil); err != nil {
					t.Error(err)
					return
				}
			}
		})
	}
	wg.Wait()
}
