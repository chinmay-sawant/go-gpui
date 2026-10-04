package window

import (
	"fmt"
	"os"
	"time"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// reloadEvery is how often a watched page is asked for file changes.
const reloadEvery = 250 * time.Millisecond

// pollReload asks the screen to reread its files when the interval has
// passed. A change is left to syncImage. An error is printed once and the
// last good frame keeps drawing.
func (s *shell) pollReload() {
	reloader, ok := s.app.(host.Reloader)
	if !ok {
		return
	}

	if !s.lastPoll.IsZero() && time.Since(s.lastPoll) < reloadEvery {
		return
	}

	s.lastPoll = time.Now()

	changed, err := reloader.PollReload(s.ctx)
	if err != nil {
		if line, show := s.reloadNote(err); show {
			fmt.Fprintln(os.Stderr, line)
		}

		return
	}

	s.lastNote = ""
	if changed {
		s.clampView()
	}
}

// reloadNote returns the stderr line for err, once per distinct message.
func (s *shell) reloadNote(err error) (string, bool) {
	line := "hot reload: " + err.Error()
	if line == s.lastNote {
		return "", false
	}

	s.lastNote = line

	return line, true
}

// clampView pulls the scroll offset back inside a page that got shorter.
func (s *shell) clampView() {
	contentW, contentH := s.contentSize()
	s.scrollX, s.scrollY = clampScroll(s.scrollX, s.scrollY, contentW, contentH, s.screenW, s.screenH)
}
