package domain

// deriveMemory fills used and percentage fields that a source left out. A
// source may know used or available, and either one fills the other.
func deriveMemory(m Memory) Memory {
	if total, ok := m.Total.Get(); ok && total > 0 {
		if used, ok := m.Used.Get(); ok {
			m.UsedPercent = Percent(used / total * 100)
		} else if free, ok := m.Available.Get(); ok {
			m.Used = Bytes(total - free)
			m.UsedPercent = Percent((total - free) / total * 100)
		} else if free, ok := m.Free.Get(); ok {
			m.Used = Bytes(total - free)
			m.UsedPercent = Percent((total - free) / total * 100)
		}
	}

	if total, ok := m.SwapTotal.Get(); ok && total > 0 {
		if used, ok := m.SwapUsed.Get(); ok {
			m.SwapUsedPercent = Percent(used / total * 100)
		}
	}

	return m
}
