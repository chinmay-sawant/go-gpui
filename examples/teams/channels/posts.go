package channels

// postSets maps a channel id to the builder for its posts. Each call builds
// fresh posts, so every channel open starts from the sample state.
var postSets = map[string]func() []Post{
	"av-general":   avengersPosts,
	"av-missions":  avengersMissions,
	"av-suitlab":   avengersSuitLab,
	"wk-general":   wakandaPosts,
	"wk-vibranium": wakandaVibranium,
	"wk-outreach":  wakandaOutreach,
	"sh-general":   shieldPosts,
	"sh-briefings": shieldBriefings,
	"st-general":   starkPosts,
	"st-rd":        starkRD,
	"gd-general":   guardiansPosts,
	"gd-milano":    guardiansMilano,
}

// postsFor returns the sample posts for one channel, or nil when unknown.
func postsFor(id string) []Post {
	build := postSets[id]
	if build == nil {
		return nil
	}

	return build()
}
