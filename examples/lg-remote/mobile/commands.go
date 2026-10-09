package mobile

import "github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"

// CommandListener wakes the Android UI thread when Bluetooth work arrives.
type CommandListener interface{ Ready() }

// SetCommandListener installs Android's dispatcher without a polling delay.
func SetCommandListener(listener CommandListener) {
	if listener == nil {
		bridge.SetCommandNotifier(nil)
		return
	}
	bridge.SetCommandNotifier(listener.Ready)
}
