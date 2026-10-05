package calendar

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// Handle applies one click action; it reports whether the action was ours.
func Handle(_ context.Context, page *gpui.Page, d *Data, action string) bool {
	switch {
	case action == "calendar-prev":
		d.goWeek(clamp(d.Week - 1))
	case action == "calendar-next":
		d.goWeek(clamp(d.Week + 1))
	case action == "calendar-today":
		d.goWeek(defaultWeek)
	case action == "calendar-new":
		d.Compose = true
	case action == "calendar-cancel":
		d.Compose = false
	case action == "calendar-create":
		d.create(page)
	case strings.HasPrefix(action, "calendar-date-"):
		d.SelDate = pickIndex(action, "calendar-date-", len(d.Days))
	case strings.HasPrefix(action, "calendar-start-"):
		d.SelStart = pickIndex(action, "calendar-start-", len(d.StartTimes))
	case strings.HasPrefix(action, "calendar-dur-"):
		d.SelDur = pickIndex(action, "calendar-dur-", len(d.Durs))
	case strings.HasPrefix(action, "calendar-open-"):
		d.selectEvent(strings.TrimPrefix(action, "calendar-open-"))
	default:
		return false
	}

	return true
}

// goWeek shows another week from the stored event list.
func (d *Data) goWeek(i int) {
	i = clamp(i)
	days, label, today := weekMeta(i)

	d.Week = i
	d.WeekLabel = label
	d.Today = today
	d.Days = days
	d.Events = d.Events[:0]
	d.Compose = false
	d.Selected = ""
	d.SelDate = 0
	d.SelStart = 0
	d.SelDur = 1

	for _, e := range d.allEvents {
		if e.Week == i {
			d.Events = append(d.Events, e)
		}
	}

	d.sortEvents()
}

// clamp keeps a week index inside the week range.
func clamp(i int) int {
	if i < 0 {
		return 0
	}

	if i >= weekCount {
		return weekCount - 1
	}

	return i
}
