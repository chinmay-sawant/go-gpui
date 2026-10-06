package scene

import (
	"errors"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// errTest is a store failure the tests inject.
var errTest = errors.New("test failure")

// fakeStore records requests and answers them from a queue.
type fakeStore struct {
	queue []Result
	seq   uint64

	settingsID, themeID, scoresID, demoID, saveID, snapID, clearID uint64

	themeCalls  []bool
	scoresCalls []int
	saves       []game.Result
	replays     []*game.Replay
	snaps       []game.Snapshot
	clears      int
	closed      bool
	refuse      int
}

func (f *fakeStore) id() uint64 { f.seq++; return f.seq }

func (f *fakeStore) Settings() uint64 { f.settingsID = f.id(); return f.settingsID }

func (f *fakeStore) SaveTheme(dark bool) uint64 {
	f.themeCalls = append(f.themeCalls, dark)
	f.themeID = f.id()

	return f.themeID
}

func (f *fakeStore) Scores(page int) uint64 {
	f.scoresCalls = append(f.scoresCalls, page)
	f.scoresID = f.id()

	return f.scoresID
}

func (f *fakeStore) DemoScores() uint64 { f.demoID = f.id(); return f.demoID }

func (f *fakeStore) SaveGame(res game.Result, rep *game.Replay) uint64 {
	if f.refuse > 0 {
		f.refuse--

		return 0
	}

	f.saves = append(f.saves, res)
	f.replays = append(f.replays, rep)
	f.saveID = f.id()

	return f.saveID
}

func (f *fakeStore) SaveSnapshot(s game.Snapshot) uint64 {
	f.snaps = append(f.snaps, s)
	f.snapID = f.id()

	return f.snapID
}

func (f *fakeStore) ClearSnapshot() uint64 {
	f.clears++
	f.clearID = f.id()

	return f.clearID
}

func (f *fakeStore) Poll() (Result, bool) {
	if len(f.queue) == 0 {
		return Result{}, false
	}

	r := f.queue[0]
	f.queue = f.queue[1:]

	return r, true
}

func (f *fakeStore) Close() error {
	f.closed = true

	return nil
}

func (f *fakeStore) push(r Result) { f.queue = append(f.queue, r) }
