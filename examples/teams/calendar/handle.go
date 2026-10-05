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

// goWeek rebuilds the state for another sample week and merges in the
// meetings the user created for that week.
func (d *Data) goWeek(i int) {
	extra := d.extra
	*d = weekData(i)
	d.extra = extra

	if d.extra != nil && d.Week < len(d.extra) {
		d.Events = append(d.Events, d.extra[d.Week]...)
	}

	d.sortEvents()
}

// clamp keeps a week index inside the sample range.
func clamp(i int) int {
	if i < 0 {
		return 0
	}

	if i >= len(weeks) {
		return len(weeks) - 1
	}

	return i
}
