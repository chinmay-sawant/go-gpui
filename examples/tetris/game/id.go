package game

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync/atomic"
	"time"
)

// idSeq makes fallback IDs unique in this process.
var idSeq atomic.Uint64

// newID returns a fresh stable game ID: 32 hex characters when the system
// random source works, a time and counter form otherwise.
func newID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err == nil {
		return hex.EncodeToString(b[:])
	}

	return fmt.Sprintf("g%016x%08x", time.Now().UnixNano(), idSeq.Add(1))
}
