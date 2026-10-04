package player

import (
	"context"
	"strings"
)

// QueueData is the data behind the Queue screen: the songs waiting next and
// how many songs the queue holds in total.
type QueueData struct {
	Count int
	Next  []Track
}

// defaultQueue returns the offline queue screen.
func defaultQueue() QueueData {
	return QueueData{
		Count: 6,
		Next: []Track{
			{Index: 0, Title: "Faded Signals", Artist: "Echo Valley", Length: "3:18", Cover: "track-0"},
			{Index: 1, Title: "Glass Highway", Artist: "Nova Waves", Length: "4:02", Cover: "track-1"},
			{Index: 2, Title: "Silver Rainfall", Artist: "Aster Field", Length: "2:47", Cover: "track-2"},
			{Index: 3, Title: "Copper Skyline", Artist: "Mono Arcade", Length: "3:36", Cover: "track-3"},
			{Index: 4, Title: "Quiet Machines", Artist: "Night Cartography", Length: "4:14", Cover: "track-4"},
			{Index: 5, Title: "Last Transmission", Artist: "The Far Coast", Length: "3:05", Cover: "track-5"},
		},
	}
}

// clickQueue applies one action on the queue screen.
func (a *App) clickQueue(ctx context.Context, action string) error {
	if strings.HasPrefix(action, "queue-row-") {
		i := slot(action, "queue-row-")
		if i < 0 || i >= len(a.view.Queue.Next) {
			return nil
		}

		for j := range a.view.Queue.Next {
			a.view.Queue.Next[j].Active = j == i
		}

		a.view.Now = a.view.Queue.Next[i]
		a.resetProgress()
		a.playNow(ctx)

		return nil
	}

	if strings.HasPrefix(action, "queue-from-") {
		i := slot(action, "queue-from-")
		if i < 0 || i >= len(a.view.Shelf) {
			return nil
		}

		for j := range a.view.Shelf {
			a.view.Shelf[j].Active = j == i
		}

		a.setNow(a.view.Shelf[i])
		a.playNow(ctx)
	}

	return nil
}
