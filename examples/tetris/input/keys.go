package input

import "github.com/chinmay-sawant/ownframe/examples/tetris/game"

// keyGroup maps a set of key names to one action.
type keyGroup struct {
	keys []string
	a    game.Action
}

// groups lists the key groups in a fixed order.
func (km Keymap) groups() []keyGroup {
	return []keyGroup{
		{km.Left, game.ActionLeft},
		{km.Right, game.ActionRight},
		{km.SoftDrop, game.ActionSoftDrop},
		{km.HardDrop, game.ActionHardDrop},
		{km.RotateCW, game.ActionRotateCW},
		{km.RotateCCW, game.ActionRotateCCW},
		{km.Pause, game.ActionPause},
		{km.Restart, game.ActionRestart},
		{km.Start, game.ActionStart},
	}
}

// actionFor returns the action a key triggers, or ActionNone.
func (km Keymap) actionFor(key string) game.Action {
	for _, g := range km.groups() {
		for _, k := range g.keys {
			if k == key {
				return g.a
			}
		}
	}

	return game.ActionNone
}

// FirstAction returns the first key mapped to an action, for hints. An
// unmapped action returns "".
func (km Keymap) FirstAction(a game.Action) string {
	for _, g := range km.groups() {
		if g.a == a && len(g.keys) > 0 {
			return g.keys[0]
		}
	}

	return ""
}
