# Keys

`Handlers.KeyDown` and `Handlers.KeyUp` run when a key goes down and up in
the window. `key` is a lowercase name: `"space"`, `"arrowup"`,
`"arrowdown"`, `"a"`, or `"1"`. Digits lose the `Digit` prefix, so
`Digit1` arrives as `"1"`.

```go
page.Handle(gpui.Handlers{
    KeyDown: func(ctx context.Context, key string) error {
        if key == "space" {
            jump()
        }

        return nil
    },
    KeyUp: func(ctx context.Context, key string) error {
        if key == "space" {
            stopHolding()
        }

        return nil
    },
})
```

Neither handler draws. A game changes its state and paints from
`Page.SetTick`; a page that must change its HTML calls `Page.Redraw` from
the handler. The window sends one press and one release per real key
event: a platform that reports auto-repeat as a release and a press does
not turn a held key into a stream of pairs.

The window still handles `Ctrl+C`, `Ctrl+X`, `Ctrl+V`, `Ctrl+A`, `Ctrl+Z`,
`Ctrl+Y`, `Ctrl+Backspace`, and Enter on its own. A key handler sees those
keys too, and typing in a focused form control still edits the control.

`Serve` has no window, so its page receives no key events. The desktop
window, the phone build, and the WebAssembly canvas all send them.
