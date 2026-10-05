package chat

// sampleThreads is the read-only message history per chat id.
var sampleThreads = map[string][]Message{
	"pepper":    pepperThread,
	"peter":     peterThread,
	"avengers":  avengersThread,
	"wakanda":   wakandaThread,
	"guardians": guardiansThread,
	"rhodey":    rhodeyThread,
	"strange":   strangeThread,
	"wanda":     wandaThread,
	"fury":      furyThread,
}

// cloneThread copies one sample thread so two Defaults never share state.
func cloneThread(id string) []Message {
	src := sampleThreads[id]
	if len(src) == 0 {
		return nil
	}

	out := make([]Message, len(src))
	copy(out, src)

	return out
}
