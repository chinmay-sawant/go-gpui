package telegram

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"image/png"
	"strconv"

	"golang.org/x/image/draw"
)

// photoMax is the widest a photo bubble shows, in CSS pixels.
const photoMax = 240

// addPhoto registers data as an image and appends a photo message to the
// chat that asked for it.
func (a *App) addPhoto(data []byte) {
	src, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return
	}

	src = fitPhoto(src)
	src = roundPhoto(src)

	var buf bytes.Buffer
	if err := png.Encode(&buf, src); err != nil {
		return
	}

	id := "photo-" + strconv.Itoa(a.photoN)
	a.photoN++
	a.page.SetImage(id, buf.Bytes())

	a.mu.Lock()
	chat := a.attachID
	a.mu.Unlock()

	if chat == "" || a.chatIndex(chat) < 0 {
		return
	}

	b := src.Bounds()
	a.threads[chat] = append(a.threads[chat], Message{
		ID: "shot-" + strconv.Itoa(a.photoN), Time: "now",
		Own: true, Read: true,
		Photo: id, PhotoW: b.Dx(), PhotoH: b.Dy(),
	})

	if i := a.chatIndex(chat); i >= 0 {
		a.chats[i].Preview = "You: a photo"
		a.chats[i].Time = "now"
	}

	a.page.ScrollTo(0, 1<<20)
	a.mark()
}

// fitPhoto scales src down to photoMax wide, keeping the aspect ratio.
func fitPhoto(src image.Image) image.Image {
	b := src.Bounds()
	w, h := b.Dx(), b.Dy()
	if w <= 0 || h <= 0 || w <= photoMax {
		return src
	}

	h = h * photoMax / w
	w = photoMax
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, b, draw.Over, nil)

	return dst
}
