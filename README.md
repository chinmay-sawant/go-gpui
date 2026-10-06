# go-gpui

<img src="assets/gopher.png" alt="go-gpui Gopher mascot" width="140" align="left" />

Write screens in HTML and CSS, and application logic in Go. go-gpui renders with [gowkhtmltopdf](https://github.com/chinmay-sawant/gowkhtmltopdf) and opens a window with Ebiten. It ships no Chromium, WebKit, or JavaScript runtime.

The same page can run on desktop, in a browser through WebAssembly, or on a phone.

## Demo

go-gpui preview. HTML and CSS layout with a layout engine written in Go, no Chromium or WebKit.

[![go-gpui preview](assets/preview.webp)](https://x.com/chinmay_sawant_/status/2106788230154871126)

[Watch the full preview on X](https://x.com/chinmay_sawant_/status/2106788230154871126)

Desktop cat overlay. Transparent, click-through cat that reports what is happening inside opencode.

[![Desktop cat overlay](assets/desktop-cat.webp)](https://x.com/chinmay_sawant_/status/2106829998409789929)

## Get started

```sh
go get github.com/chinmay-sawant/go-gpui
```

```go
package main

import (
    "context"
    "log"

    gpui "github.com/chinmay-sawant/go-gpui"
)

func main() {
    page, err := gpui.New(gpui.Config{
        Title:  "Hello",
        HTML:   `<h1>{{.Title}}</h1>`,
        Width:  480,
        Height: 640,
    })
    if err != nil {
        log.Fatal(err)
    }
    page.SetData(struct{ Title string }{"Hello"})
    if err := gpui.Run(context.Background(), page); err != nil {
        log.Fatal(err)
    }
}
```

## Run the demos

From this checkout:

```sh
go run ./examples/login
go run ./examples/login -web
sh scripts/browser.sh
go run ./examples/desktop-cat
sh scripts/android.sh install
```

The login demo accepts `secret` for both fields. [All examples](examples/readme.md), [desktop cat setup](examples/desktop-cat/readme.md), and [Android setup](examples/telegram/android/README.md) have the details. See [showcase.md](showcase.md) for Telegram screenshots.

## Documentation

Start with the [documentation index](documentation/README.md) or [feature list and limits](documentation/features.md).

- [Windows and sizing](documentation/window.md)
- [Desktop, WebAssembly, and mobile](documentation/platforms.md)
- [Templates and rendering](documentation/screen.md)
- [Forms and data binding](documentation/binding.md)
- [Theming](documentation/theming.md)
- [Packaging](documentation/packaging.md)
- [Comparisons with other UI libraries](documentation/compare-syntax.md)

The screen is a laid-out page, with no live DOM, JavaScript, video, canvas, or WebGL. There is one window and no accessibility tree. The [feature list](documentation/features.md) records the supported behavior and remaining limits.

## Development

`make build` compiles both Go modules without linking executables. `make test` runs their tests one package at a time.

The React documentation and demo website lives in [frontend/](frontend/README.md). Run `npm ci` and `npm run dev` there to preview it locally.
