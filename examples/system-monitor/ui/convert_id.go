package ui

import (
	"strconv"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// parseID turns "pid:start" back into an identity. A malformed or empty id
// clears the selection.
func parseID(id string) domain.ProcessIdentity {
	pid, start, ok := strings.Cut(id, ":")
	if !ok {
		return domain.ProcessIdentity{}
	}

	p, err1 := strconv.ParseInt(pid, 10, 32)
	s, err2 := strconv.ParseUint(start, 10, 64)

	if err1 != nil || err2 != nil {
		return domain.ProcessIdentity{}
	}

	return domain.ProcessIdentity{PID: int32(p), Start: s}
}
