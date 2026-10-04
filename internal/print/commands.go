package print

import "strings"

// linuxCommands print directly, then open the file in a viewer.
var linuxCommands = []command{
	{name: "lp", args: fileArg},
	{name: "xdg-open", args: fileArg},
}

// darwinCommands open the print dialog, then fall back to the viewer.
var darwinCommands = []command{
	{name: "osascript", args: appleScriptPrint},
	{name: "open", args: fileArg},
}

// windowsCommands hand the file to the shell print verb.
var windowsCommands = []command{
	{name: "powershell.exe", args: powershellPrint},
	{name: "pwsh.exe", args: powershellPrint},
}

func fileArg(path string) []string {
	return []string{path}
}

// appleScriptPrint asks Preview to print the file.
func appleScriptPrint(path string) []string {
	escaped := strings.ReplaceAll(path, `\`, `\\`)
	escaped = strings.ReplaceAll(escaped, `"`, `\"`)

	return []string{"-e", `tell application "Preview" to print (POSIX file "` + escaped + `")`}
}

// powershellPrint asks the shell to run the file's print verb.
func powershellPrint(path string) []string {
	quoted := strings.ReplaceAll(path, "'", "''")

	return []string{
		"-NoProfile",
		"-Command",
		"Start-Process -FilePath '" + quoted + "' -Verb Print",
	}
}
