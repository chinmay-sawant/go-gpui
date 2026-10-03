package render_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

const imageSource = `<body style="margin:0">` +
	`<div style="width:40px;height:40px;background-image:url('bg');` +
	`background-size:100% 100%"></div></body>`

const redPNG64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJ" +
	"AAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="

// imageResolver answers "bg" with the PNG and rejects any other src.
func imageResolver(t *testing.T) func(string) ([]byte, error) {
	t.Helper()

	data, err := base64.StdEncoding.DecodeString(redPNG64)
	if err != nil {
		t.Fatal(err)
	}

	return func(src string) ([]byte, error) {
		if src != "bg" {
			return nil, errors.New("unexpected src " + src)
		}

		return data, nil
	}
}

func TestDisplayListStateResolvesImage(t *testing.T) {
	t.Parallel()

	display, err := render.DisplayListState(
		context.Background(), imageSource, 120, 80, render.State{Images: imageResolver(t)},
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, paintOp := range display.Ops {
		if paintOp.Kind != render.OpImage {
			continue
		}

		data, width, height := paintOp.ImageBytes()
		if len(data) == 0 || width <= 0 || height <= 0 {
			t.Fatalf("image op %d bytes %dx%d", len(data), width, height)
		}

		if !bytes.HasPrefix(data, []byte("\x89PNG")) {
			t.Fatal("payload is not a PNG")
		}

		return
	}

	t.Fatal("no image op for a resolved background")
}

func TestDisplayListStateKeepsImagesOffWithoutResolver(t *testing.T) {
	t.Parallel()

	display, err := render.DisplayListState(
		context.Background(), imageSource, 120, 80, render.State{},
	)
	if err != nil {
		t.Fatal(err)
	}

	for _, paintOp := range display.Ops {
		if paintOp.Kind == render.OpImage {
			t.Fatal("image resolved without a resolver")
		}
	}
}
