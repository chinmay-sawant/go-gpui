package telegram_test

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"testing"
)

func pngBytes(t *testing.T, w, h int) []byte {
	t.Helper()

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	var buf bytes.Buffer

	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}

	return buf.Bytes()
}

func TestAttachRequestAndPhotoBubble(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")
	click(t, ctx, app, "attach", "")
	click(t, ctx, app, "attach-camera", "")

	if app.TakeAttach() != "camera" {
		t.Fatal("no camera request")
	}

	if app.TakeAttach() != "" {
		t.Fatal("attach request repeated")
	}

	app.QueuePhoto(pngBytes(t, 400, 300))
	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	thread := app.View().Thread
	last := thread[len(thread)-1]

	if last.Photo == "" || !last.Own {
		t.Fatalf("last = %+v", last)
	}

	if last.PhotoW != 240 || last.PhotoH != 180 {
		t.Fatalf("photo = %dx%d", last.PhotoW, last.PhotoH)
	}
}

func TestPhotoFollowsTheAskingChat(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-max")
	click(t, ctx, app, "attach", "")
	click(t, ctx, app, "attach-gallery", "")
	click(t, ctx, app, "chat-back", "")

	app.QueuePhoto(pngBytes(t, 20, 10))
	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	for _, c := range app.View().Chats {
		if c.ID == "max" && c.Preview != "You: a photo" {
			t.Fatalf("max preview = %q", c.Preview)
		}
	}

	click(t, ctx, app, "", "open-max")

	thread := app.View().Thread
	if last := thread[len(thread)-1]; last.Photo == "" {
		t.Fatalf("last = %+v", last)
	}
}
