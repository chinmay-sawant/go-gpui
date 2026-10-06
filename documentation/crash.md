# Crash reports

`Run` stores every failure in a report file and returns an error naming that file: a panic goes through the recover, and any startup or window error goes through the same writer with the error as the reason. `BindMobile` recovers a panic on its own goroutine. `Serve` does not. The process does not panic again, and the file is not uploaded. In a browser build the recover still runs and returns an error, but no file is written, because Go's wasm file operations return an error there.

```go
ownframe.SetCrashDir(dir)
path, err := ownframe.Report("Hello", "disk full")
```

`Report` is the same writer, for a caller that already has a reason and does not want to panic. `SetCrashDir` points both `Report` and a recovered panic at `dir`. An empty dir restores the default.

The default folder is `ownframe/crashes` under `os.UserConfigDir`. If that directory cannot be resolved, the folder is `ownframe-crashes` under `os.TempDir`.

The file name starts with a UTC timestamp `20060102-150405`. A second write in that same second gets a numeric suffix. The body contains the title, the reason, `runtime.Version()`, and `debug.Stack()`.

A nil page passed to `Serve` or `BindMobile` still returns bare `ErrNilPage` from the existing prepare check. `Run` wraps it in the report error above, which keeps `ErrNilPage` for `errors.Is`. The recover does not call `Title` on a nil page.

A panic inside the Ebiten update loop is caught only when it unwinds through `Run` on that same goroutine.

The [crash example](../examples/crash) sets the folder to `crashes/`, writes a manual report from one button, and panics from another.
