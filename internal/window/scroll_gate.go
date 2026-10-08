package window

import "github.com/chinmay-sawant/ownframe/internal/host"

func (s *shell) allowPageScroll() bool {
	gate, ok := s.app.(host.ScrollGate)

	if !ok {
		return true
	}

	return gate.AllowScroll()
}
