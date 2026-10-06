package ui

// detailValue formats one detail row.
func detailValue(key string, sel selection) string {
	d := sel.trk.Detail

	switch key {
	case "status":
		return detailStatus(sel)
	case "state":
		return nonEmpty(d.State)
	case "cpu":
		return cpuText(d.CPU, d.CPUKnown)
	case "mem":
		return formatBytes(float64(d.Mem), d.MemKnown)
	case "threads":
		return countOrNA(d.Threads, d.ThreadsKnown)
	case "handles":
		return countU64(d.Handles, d.HandlesOK)
	case "priority":
		if !d.PriorityOK {
			return unavailable
		}

		return formatCount(int(d.Priority))
	case "started":
		if d.Start.IsZero() {
			return unavailable
		}

		return d.Start.Format("2006-01-02 15:04")
	case "virtual":
		return formatBytes(float64(d.Virtual), d.VirtualOK)
	case "io":
		return formatRate(d.ReadRate, d.ReadRateOK) + " read \u00b7 " +
			formatRate(d.WriteRate, d.WriteRateOK) + " write"
	case "parent":
		return nonEmpty(d.Parent)
	case "exe":
		return truncate(nonEmpty(d.Exe), 96)
	case "cwd":
		return truncate(nonEmpty(d.Cwd), 96)
	case "command":
		return truncate(nonEmpty(d.Command), 160)
	}

	return unavailable
}
