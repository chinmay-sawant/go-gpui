package scene

import (
	"errors"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// ErrShutdownTimeout reports a store worker that outlived ShutdownBudget.
var ErrShutdownTimeout = errors.New("scene: store worker did not stop")

// Store is the scene's asynchronous view of the storage worker. Every
// call returns at once; finished work arrives through Poll, and the scene
// drains a bounded number of results per frame. A zero request id means
// the request was refused and will be retried.
type Store interface {
	Settings() uint64
	SaveTheme(dark bool) uint64
	Scores(page int) uint64
	DemoScores() uint64
	SaveGame(game.Result, *game.Replay) uint64
	SaveSnapshot(game.Snapshot) uint64
	ClearSnapshot() uint64
	Poll() (Result, bool)
	Close() error
}
