package store

// dummyScores are the seeded demo entries. Their IDs are stable, so
// reseeding never duplicates and never overwrites a real score. Each
// record pins the ruleset and fixture versions that produced it.
var dummyScores = []struct {
	score, lines, level, pieces int
	seed                        uint64
	day                         int
}{
	{240000, 120, 12, 410, 0x7e7e15, 1},
	{151500, 96, 9, 350, 0x51a2b3, 2},
	{98700, 70, 8, 280, 0x9c4d5e, 3},
	{84100, 62, 7, 260, 0x3f6a7b, 4},
	{71250, 54, 6, 240, 0x2b8c9d, 5},
	{62500, 48, 5, 220, 0xd4e5f6, 6},
	{54000, 42, 5, 200, 0xa1b2c3, 7},
	{47100, 37, 4, 185, 0x8e9fa0, 8},
	{39800, 31, 4, 170, 0x6d7e8f, 9},
	{31200, 25, 3, 150, 0x5c6d7e, 10},
	{22600, 19, 2, 130, 0x4b5c6d, 11},
	{15400, 13, 2, 110, 0x3a4b5c, 12},
}
