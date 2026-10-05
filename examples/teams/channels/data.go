package channels

// descriptions maps a channel id to the line under its name.
var descriptions = map[string]string{
	"av-general":   "Mission wins, suit news, and shawarma plans.",
	"av-missions":  "Debriefs, targets, and extraction windows.",
	"av-suitlab":   "Suit tests, schematics, and armor upgrades.",
	"wk-general":   "Updates from the Wakanda R&D team.",
	"wk-vibranium": "Yields, testing, and secure transport.",
	"wk-outreach":  "Relief runs, school kits, and first aid.",
	"sh-general":   "Field ops updates and shift handoffs.",
	"sh-briefings": "Briefings, clearances, and field reports.",
	"st-general":   "Company updates, events, and badges.",
	"st-rd":        "Repulsors, reactors, and prototypes.",
	"gd-general":   "Crew updates and orbital logistics.",
	"gd-milano":    "Maintenance, jump points, and cargo.",
}

// describe returns the description for a channel id.
func describe(id string) string {
	if desc, ok := descriptions[id]; ok {
		return desc
	}

	return "Updates and conversation for this channel."
}
