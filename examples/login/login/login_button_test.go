package login_test

import (
	"context"
	"testing"
)

func TestLoginButtonIsAButton(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	login := boxByID(t, app, "login")
	if login.Tag != "button" || login.W <= 0 || login.H <= 0 {
		t.Fatalf("login control = %+v", login)
	}
}
