package dino

import "strconv"

// dinoParts are the fill elements of the dinosaur, in paint order: tail,
// body, neck, head, eye, arm, and the two legs.
var dinoParts = [8]string{
	"d-tail", "d-body", "d-neck", "d-head",
	"d-eye", "d-arm", "d-leg1", "d-leg2",
}

// obstacleParts are the fill elements of one obstacle slot: the bird body,
// its two wings and beak, then the small cactus and the big cactus.
var obstacleParts = [10]string{
	"w1", "w2", "w3", "w4", "s1", "s2", "s3", "b1", "b2", "b3",
}

// cloudParts are the six puffs of the two clouds.
var cloudParts = [6]string{"c0a", "c0b", "c0c", "c1a", "c1b", "c1c"}

// pebbleParts are the six ground marks.
var pebbleParts = [6]string{"p0", "p1", "p2", "p3", "p4", "p5"}

// slotID names one obstacle part, such as "o2-s1".
func slotID(slot int, part string) string {
	return "o" + strconv.Itoa(slot) + "-" + part
}
