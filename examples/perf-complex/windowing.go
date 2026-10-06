package main

import "github.com/chinmay-sawant/ownframe/internal/page"

func installWindow(p *page.Page, v view) {
	all := v.Rows
	previousStart, previousEnd := -1, -1
	p.SetWindowing(true)
	p.SetScrollWindow(func(offsetY, viewH int) bool {
		gridY := 0
		for _, b := range p.Boxes() {
			if b.ID == "grid" {
				gridY = int(b.Y)
				break
			}
		}
		first := max(0, (offsetY-gridY)/76)
		last := min(len(all), max(0, (offsetY+viewH-gridY+75)/76))
		topSafe := previousStart == 0 || first >= previousStart+5
		bottomSafe := previousEnd == len(all) || last <= previousEnd-5
		if previousStart >= 0 && topSafe && bottomSafe {
			return false
		}
		start := (offsetY-gridY)/76 - 10
		if start < 0 {
			start = 0
		}
		count := viewH/76 + 21
		if count < 48 {
			count = 48
		}
		if start > len(all)-count {
			start = len(all) - count
		}
		if start < 0 {
			start = 0
		}
		// Keep nth-child stripes aligned with the full document.
		start -= start % 2
		end := min(start+count, len(all))
		if start == previousStart && end == previousEnd {
			return false
		}
		previousStart, previousEnd = start, end
		v.Rows = all[start:end]
		v.Top = start * 76
		v.Bottom = (len(all) - end) * 76
		p.SetData(v)
		return true
	})
}
