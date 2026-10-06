// Package mobile is the Android and iOS entry for the Telegram demo.
//
// Desktop and browser builds use examples/telegram. A phone build binds this
// package with ebitenmobile. The generated view calls the game registered
// here. Do not call ownframe.Run from this package.
package mobile

// Dummy exists so ebitenmobile bind will compile this package.
// The bind tool skips a package that exports nothing.
func Dummy() {}

func init() {
	if err := start(); err != nil {
		panic(err)
	}
}
