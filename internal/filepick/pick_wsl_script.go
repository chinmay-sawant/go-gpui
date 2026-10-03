//go:build linux && !android

package filepick

import "strings"

func windowScript(title string) string {
	quoted := strings.ReplaceAll(title, "'", "''")

	return `[Console]::OutputEncoding = [System.Text.Encoding]::UTF8;` +
		`Add-Type -AssemblyName System.Windows.Forms;` +
		`$f = New-Object System.Windows.Forms.OpenFileDialog;` +
		`$f.Title = '` + quoted + `';` +
		`if ($f.ShowDialog() -eq [System.Windows.Forms.DialogResult]::OK) ` +
		`{ [Console]::Out.Write($f.FileName) }`
}
