package host

import "time"

// RepeatRate selects an interval after a held duration.
type RepeatRate struct {
	After    time.Duration
	Interval time.Duration
}

// RepeatPolicy schedules repeats after the caller's initial action.
type RepeatPolicy struct {
	Delay time.Duration
	Rates []RepeatRate
}

// RepeatScheduler advances on the host's frame loop; it starts no goroutines.
type RepeatScheduler struct {
	policy  RepeatPolicy
	started time.Time
	next    time.Time
	active  bool
}

// Start arms a schedule. Rates must have ascending After values and positive
// intervals. Invalid policies leave the scheduler stopped.
func (s *RepeatScheduler) Start(p RepeatPolicy, now time.Time) bool {
	s.Stop()
	if p.Delay < 0 || len(p.Rates) == 0 {
		return false
	}
	for i, rate := range p.Rates {
		if rate.After < 0 || rate.Interval <= 0 || i > 0 && rate.After <= p.Rates[i-1].After {
			return false
		}
	}
	p.Rates = append([]RepeatRate(nil), p.Rates...)
	s.policy, s.started, s.next, s.active = p, now, now.Add(p.Delay), true
	return true
}

// Stop cancels repeats and clears the active policy.
func (s *RepeatScheduler) Stop() {
	s.policy, s.started, s.next, s.active = RepeatPolicy{}, time.Time{}, time.Time{}, false
}

// Tick reports whether one repeat is due. Missed intervals do not burst; the
// next repeat is scheduled from now using the rate for the current hold time.
func (s *RepeatScheduler) Tick(now time.Time) bool {
	if !s.active || now.Before(s.next) {
		return false
	}
	interval := s.policy.Rates[0].Interval
	held := now.Sub(s.started)
	for _, rate := range s.policy.Rates {
		if held < rate.After {
			break
		}
		interval = rate.Interval
	}
	s.next = now.Add(interval)
	return true
}
