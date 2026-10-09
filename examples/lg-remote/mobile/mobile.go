// Package mobile is the Android entry for the LG remote.
//
// Desktop builds use examples/lg-remote. A phone build binds this package
// with ebitenmobile. The activity polls TakeCommand for Bluetooth HID
// reports. Do not call ownframe.Run from this package.
package mobile

import "github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"

// Dummy exists so ebitenmobile bind will compile this package.
func Dummy() {}

func init() {
	if err := start(); err != nil {
		panic(err)
	}
}

// SetStoreDir selects where the paired TV key is saved.
func SetStoreDir(path string) { bridge.SetDir(path) }

// TakeCommand returns the next Bluetooth command, or "".
func TakeCommand() string { return bridge.Take() }

// SetBluetooth records a radio status line for the page.
func SetBluetooth(state string) { bridge.SetBluetooth(state) }

// SetFontSize reports the system-scaled base text size in CSS pixels.
func SetFontSize(size int) { bridge.SetFontSize(size) }
