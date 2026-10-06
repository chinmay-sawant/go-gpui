package channels

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/ownframe"
)

// Handle applies one click action; it reports whether the action was ours.
func Handle(_ context.Context, page *ownframe.Page, d *Data, action string) bool {
	switch {
	case strings.HasPrefix(action, "channels-toggle-"):
		toggleTeam(d, strings.TrimPrefix(action, "channels-toggle-"))
	case strings.HasPrefix(action, "channels-open-"):
		openChannel(d, strings.TrimPrefix(action, "channels-open-"))
	case strings.HasPrefix(action, "channels-tab-"):
		setTab(d, strings.TrimPrefix(action, "channels-tab-"))
	case strings.HasPrefix(action, "channels-like-"):
		toggleLike(d, strings.TrimPrefix(action, "channels-like-"))
	case strings.HasPrefix(action, "channels-pin-"):
		togglePin(d, strings.TrimPrefix(action, "channels-pin-"))
	case strings.HasPrefix(action, "channels-star-"):
		toggleStar(d, strings.TrimPrefix(action, "channels-star-"))
	case strings.HasPrefix(action, "channels-expand-"):
		toggleExpand(d, strings.TrimPrefix(action, "channels-expand-"))
	case strings.HasPrefix(action, "channels-reply-send-"):
		sendReply(page, d, strings.TrimPrefix(action, "channels-reply-send-"))
	case strings.HasPrefix(action, "channels-react-"):
		toggleReaction(d, strings.TrimPrefix(action, "channels-react-"))
	case strings.HasPrefix(action, "channels-add-"):
		addReaction(d, strings.TrimPrefix(action, "channels-add-"))
	case strings.HasPrefix(action, "channels-picker-"):
		togglePicker(d, strings.TrimPrefix(action, "channels-picker-"))
	default:
		return false
	}

	return true
}

// setTab switches the main pane tab.
func setTab(d *Data, tab string) {
	switch tab {
	case "posts", "files", "notes":
		d.Tab = tab
	}
}
