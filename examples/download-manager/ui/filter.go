package ui

// Filter selects which durable history rows a page shows.
type Filter uint8

// The four history filters. The active queue ignores them.
const (
	FilterAll Filter = iota
	FilterCompleted
	FilterFailed
	FilterCancelled
)

// Label returns the button text.
func (f Filter) Label() string {
	switch f {
	case FilterCompleted:
		return "Completed"
	case FilterFailed:
		return "Failed"
	case FilterCancelled:
		return "Cancelled"
	default:
		return "All"
	}
}

// Key returns the stable id used in data-action names.
func (f Filter) Key() string {
	switch f {
	case FilterCompleted:
		return "completed"
	case FilterFailed:
		return "failed"
	case FilterCancelled:
		return "cancelled"
	default:
		return "all"
	}
}

// ParseFilter turns a data-action suffix back into a filter; an unknown key
// maps to FilterAll.
func ParseFilter(key string) Filter {
	switch key {
	case "completed":
		return FilterCompleted
	case "failed":
		return FilterFailed
	case "cancelled":
		return FilterCancelled
	default:
		return FilterAll
	}
}
