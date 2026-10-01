package login

import (
	"unicode"
	"unicode/utf8"
)

func (a *App) signIn() {
	status := ""
	errMsg := "Unknown email or password."

	if a.view.Email == "secret" && a.view.Password == "secret" {
		status = "Signed in."
		errMsg = ""
	}

	if a.view.Status == status && a.view.Error == errMsg {
		return
	}

	a.push()
	a.view.Status = status
	a.view.Error = errMsg
}

func (a *App) setField(field *string, next string) {
	if *field == next && !a.view.Selected && a.view.Status == "" && a.view.Error == "" {
		return
	}

	a.push()
	*field = next
	a.view.Selected = false
	a.view.Status = ""
	a.view.Error = ""
	a.page.SetData(a.view)
}

func (a *App) push() {
	a.undo = append(a.undo, a.view)
	if len(a.undo) > undoLimit {
		a.undo = a.undo[len(a.undo)-undoLimit:]
	}

	a.redo = nil
}

func (a *App) focused() *string {
	switch a.view.Focus {
	case emailField:
		return &a.view.Email
	case passwordField:
		return &a.view.Password
	default:
		return nil
	}
}

func dropLastRune(s string) string {
	if s == "" {
		return ""
	}

	_, size := utf8.DecodeLastRuneInString(s)

	return s[:len(s)-size]
}

func dropLastWord(s string) string {
	runes := []rune(s)
	i := len(runes)

	for i > 0 && unicode.IsSpace(runes[i-1]) {
		i--
	}

	for i > 0 && !unicode.IsSpace(runes[i-1]) {
		i--
	}

	return string(runes[:i])
}
