package telegram_test

import (
	"context"
	"strings"
	"testing"
)

var bg = context.Background()

func TestListScreensKeepTheirBars(t *testing.T) {
	app := newApp(t, bg)
	for _, tab := range []string{"tab-chats", "tab-contacts", "tab-settings", "tab-chats"} {
		click(t, bg, app, tab, "")
		top := boxByID(t, app, "tab-chats").Y
		for _, b := range app.Boxes() {
			if strings.HasPrefix(b.ID, "chat-") && b.Y >= top {
				t.Fatalf("%s %.1f >= %.1f", b.ID, b.Y, top)
			}
		}
	}
}

func TestDrawerCoveredByRegression(t *testing.T) {
	app := newApp(t, bg)
	click(t, bg, app, "", "open-anna")
	click(t, bg, app, "attach", "")
	for _, id := range []string{"attach-camera", "attach-gallery", "compose"} {
		boxByID(t, app, id)
	}
	r := boxByID(t, app, "attach-gallery")
	c := boxByID(t, app, "compose")
	if r.Y+r.H > c.Y+1 {
		t.Fatal("sheet not above composer")
	}
	click(t, bg, app, "attach", "")
	if hasBox(app, "attach-camera") {
		t.Fatal("sheet still drawn")
	}
}

func TestGiftOnlyInAnnasChatStill(t *testing.T) {
	app := newApp(t, bg)
	click(t, bg, app, "", "open-anna")
	click(t, bg, app, "gift", "")
	if v := app.View().Thread; !v[len(v)-1].Gift {
		t.Fatal("last is not a gift")
	}
	click(t, bg, app, "chat-back", "")
	click(t, bg, app, "", "open-max")
	if app.View().CanGift || hasBox(app, "gift") {
		t.Fatal("gift in Max's chat")
	}
}

func TestReactionRowStillWorks(t *testing.T) {
	app := newApp(t, bg)
	click(t, bg, app, "", "open-anna")
	id := ""
	for _, b := range app.Boxes() {
		if strings.HasPrefix(b.ID, "msg-") {
			id = b.ID
			break
		}
	}
	if id == "" {
		t.Fatalf("no msg:%s", dumpBoxes(app))
	}
	longPress(t, bg, app, id)
	react := ""
	for _, b := range app.Boxes() {
		if strings.HasPrefix(b.ID, "react-") {
			react = b.ID
			break
		}
	}
	if react == "" {
		t.Fatal("no reaction box")
	}
	click(t, bg, app, react, "")
	for _, m := range app.View().Thread {
		if "msg-"+m.ID == id && m.Reaction != "" {
			return
		}
	}
	t.Fatal("no reaction")
}
