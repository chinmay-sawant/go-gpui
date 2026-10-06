package ui

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/chinmay-sawant/ownframe"
)

// App is the live log viewer screen.
type App struct {
	page      *ownframe.Page
	feed      Feed
	dark      bool
	perf      bool
	exportDir string
	pollEvery time.Duration

	view    View
	pager   Pager
	filters filterState
	follow  followState

	sources      []SourceInfo
	activeSource string
	minSev       string

	selected   int64
	selAnchor  int64
	detailOpen bool
	detail     Detail
	detailID   int64

	winStart, winEnd int
	winEmpty         bool
	viewH            int
	lastViewH        int
	width            int
	pendingOffset    int
	havePending      bool
	intent           pageIntent
	settingsLoaded   bool
	exporting        bool

	pulseGen uint64
	pulseDot *ownframe.DisplayOp

	mu           sync.Mutex
	pageCancel   context.CancelFunc
	detailCancel context.CancelFunc

	pageGen   atomic.Uint64
	srcGen    atomic.Uint64
	detailGen atomic.Uint64
	tailAfter atomic.Int64
	pollOn    atomic.Bool
	dropped   atomic.Uint64

	ctx    context.Context
	cancel context.CancelFunc
	reqs   chan req
	outs   chan out
	wg     sync.WaitGroup

	pendingMu  sync.Mutex
	pendingReq *req

	note       string
	noteUntil  time.Time
	sourcesDue time.Time
	tailDue    time.Time
	lastDraw   time.Duration
	maxDraw    time.Duration
	lastTick   time.Duration
	started    time.Time
}
