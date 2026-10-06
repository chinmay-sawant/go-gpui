package scene

import (
	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// Settings asks for the saved configuration and resume snapshot.
func (w *worker) Settings() uint64 {
	return w.send(request{kind: ResultSettings, id: w.next()})
}

// SaveTheme stores the theme choice.
func (w *worker) SaveTheme(dark bool) uint64 {
	return w.send(request{kind: ResultTheme, id: w.next(), dark: dark})
}

// Scores asks for one live score page.
func (w *worker) Scores(page int) uint64 {
	return w.send(request{kind: ResultScores, id: w.next(), page: page})
}

// DemoScores asks for the seeded demo entries.
func (w *worker) DemoScores() uint64 {
	return w.send(request{kind: ResultDummy, id: w.next()})
}

// SaveGame queues a completed run. A full queue returns 0 so the scene
// retries the important command later.
func (w *worker) SaveGame(res game.Result, rep *game.Replay) uint64 {
	req := request{kind: ResultSaved, id: w.next(), res: res, rep: rep}

	select {
	case w.reqs <- req:
		return req.id
	case <-w.quit:
		return 0
	default:
		return 0
	}
}

// SaveSnapshot stores the resume slot.
func (w *worker) SaveSnapshot(snap game.Snapshot) uint64 {
	return w.send(request{kind: ResultSnapshot, id: w.next(), snap: snap})
}

// ClearSnapshot drops the resume slot.
func (w *worker) ClearSnapshot() uint64 {
	return w.send(request{kind: ResultSnapshot, id: w.next(), clear: true})
}
