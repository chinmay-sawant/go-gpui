package scene

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/store"
)

// exec runs one request against the store with a bounded deadline.
func (w *worker) exec(req request) {
	ctx, cancel := context.WithTimeout(w.base, store.QueryTimeout+time.Second)
	defer cancel()

	switch req.kind {
	case ResultSettings:
		w.execSettings(ctx, req)
	case ResultTheme:
		w.set.Theme = themeName(req.dark)
		w.push(Result{ID: req.id, Kind: req.kind, Err: w.st.SaveSettings(ctx, w.set)})
	case ResultScores:
		w.execScores(ctx, req)
	case ResultDummy:
		all, err := w.st.DummyScores(ctx, PageSize)
		w.push(Result{ID: req.id, Kind: req.kind, Err: err, Scores: pageOf(all, 0, PageSize)})
	case ResultSaved:
		w.push(Result{ID: req.id, Kind: req.kind, Err: w.st.SaveGame(ctx, req.res, req.rep)})
	case ResultSnapshot:
		w.execSnapshot(ctx, req)
	}
}

// execSettings loads the configuration and the resume snapshot.
func (w *worker) execSettings(ctx context.Context, req request) {
	set, err := w.st.Settings(ctx)
	if err == nil {
		w.set = set
	}

	snap, has, serr := w.st.LoadSnapshot(ctx)
	if err == nil && serr != nil {
		err = serr
	}

	w.push(Result{ID: req.id, Kind: req.kind, Err: err, Set: SettingsResult{
		Settings: sceneSettings(set),
		Snapshot: snap,
		Has:      has,
	}})
}

// execScores pages the live ranking by asking for one page more than the
// previous request, because the store lists a limit and not an offset.
func (w *worker) execScores(ctx context.Context, req request) {
	limit := PageSize * (req.page + 1)

	all, err := w.st.TopScores(ctx, limit)
	w.push(Result{ID: req.id, Kind: req.kind, Err: err, Scores: pageOf(all, req.page, limit)})
}

// execSnapshot saves or clears the single resume slot.
func (w *worker) execSnapshot(ctx context.Context, req request) {
	var err error

	if req.clear {
		err = w.st.ClearSnapshot(ctx)
	} else {
		err = w.st.SaveSnapshot(ctx, req.snap)
	}

	w.push(Result{ID: req.id, Kind: req.kind, Err: err})
}
