package flappy

import "strconv"

// birdIDs are the six fills of the bird, in paint order: body, belly,
// wing, eye, pupil, beak.
var birdIDs = [birdCount]string{
	"b-body", "b-belly", "b-wing", "b-eye", "b-pupil", "b-beak",
}

// pipeParts are the four fills of one pipe slot: top body, top cap,
// bottom body, bottom cap.
var pipeParts = [pipePartCount]string{"t", "tc", "b", "bc"}

// cloudIDs are the drifting clouds.
var cloudIDs = [cloudMax]string{"c0", "c1", "c2"}

// stripeIDs are the ground marks that scroll with the pipes.
var stripeIDs = [stripeMax]string{"s0", "s1", "s2", "s3", "s4", "s5"}

// Counts the App and the paint step index by.
const (
	birdCount     = 6
	pipeSlots     = 3
	pipePartCount = 4
	cloudMax      = 3
	stripeMax     = 6
)

// pipeID names one pipe part, such as "p2tc".
func pipeID(slot int, part string) string {
	return "p" + strconv.Itoa(slot) + part
}
