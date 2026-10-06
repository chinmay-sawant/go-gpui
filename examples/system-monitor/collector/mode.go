package collector

// Mode selects where samples come from.
type Mode string

const (
	// ModeDummy reads the labeled fixture, so the example is useful without
	// permissions or platform support.
	ModeDummy Mode = "dummy"
	// ModeLive reads the host.
	ModeLive Mode = "live"
)
