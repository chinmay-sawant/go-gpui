package render

import "sync"

// The pinned engine replaces shared document readers during Parse and Apply.
// Guard registration and reads together until the engine publishes them once.
var engineMu sync.Mutex
