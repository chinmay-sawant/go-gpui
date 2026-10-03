// Package filepick opens the dialog that picks one file.
// The compiler selects the implementation for the current system:
// comdlg32 on Windows, zenity or kdialog on Linux, and osascript on macOS.
// A system without a dialog returns ok false.
package filepick
