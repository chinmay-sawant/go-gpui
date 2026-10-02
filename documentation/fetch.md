# Fetch and XHR

`Fetch` and `XHR` use `net/http` for one request. They are not a browser network stack. The root file `fetch.go` forwards them to `internal/fetch`.

```go
res, err := gpui.Fetch(ctx, "https://example.com/health")
if err != nil {
    return err
}
_, err = gpui.XHR(ctx, "POST", "https://example.com/note",
    map[string]string{"Content-Type": "text/plain"},
    []byte("hi"),
)
```

`Fetch` always uses GET and sends no body. An empty method on `XHR` also means GET.

`FetchResponse` has `Status`, `Header`, and `Body`. `Header` is a flat map. Several values for one key are joined with `", "`.

The client rejects an empty URL and any scheme other than `http` or `https`, including `file`, `data`, and `javascript`. That error is `ErrScheme`. A redirect whose next URL is not `http` or `https` is also `ErrScheme`. After ten redirects the client stops.

A nil context returns an error whose text contains `nil context`. Cancellation uses `http.NewRequestWithContext`.

The shared client has no cookie jar and no cache. It does not store a session.
