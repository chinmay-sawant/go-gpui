package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// menuState is the shell context menu.
type menuState struct {
	open  bool
	x, y  int
	items []host.MenuItem
}

// openMenu asks the page for its rows and shows the menu at a point. A page
// without a context menu keeps the right click.
func (s *shell) openMenu(x, y int) {
	menu, ok := s.app.(host.ContextMenu)
	if !ok {
		return
	}

	items := menu.ContextMenu()
	if len(items) == 0 {
		return
	}

	s.menu = menuState{open: true, x: x, y: y, items: items}
}

// closeMenu hides the menu.
func (s *shell) closeMenu() {
	s.menu.open = false
	s.menu.items = nil
}

// menuPointer handles right and left presses over the menu. It reports
// whether the event was consumed, so the page does not also see the click.
func (s *shell) menuPointer(x, y int) (bool, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.openMenu(x, y)

		return true, nil
	}

	if !s.menu.open || !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false, nil
	}

	return s.menuClick(x, y)
}

// menuClick acts on a left click at a point and closes the menu. It reports
// whether the menu consumed the click.
func (s *shell) menuClick(x, y int) (bool, error) {
	if !s.menu.open {
		return false, nil
	}

	row := s.menuIndex(x, y)
	if row >= 0 && s.menu.items[row].Enabled {
		err := s.menuAction(s.menu.items[row].ID)
		s.closeMenu()

		return true, err
	}

	s.closeMenu()

	return true, nil
}
