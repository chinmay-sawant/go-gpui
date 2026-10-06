// Package game is the pure Tetris model for the ownframe tetris example:
// board coordinates, piece rotations, collision, scoring, and game states.
// It imports no UI or storage, so tests and replays drive it directly.
package game

import "time"

// Ruleset and fixture versions identify scores, snapshots, and replays.
const (
	Ruleset        = "tetris-classic-1"
	RulesetVersion = 1
	FixtureVersion = 1
	DummySeed      = 0x7e7e15
)

// Board geometry, bounded values, and the fixed simulation step.
const (
	Cols        = 10
	Rows        = 20
	Preview     = 3
	MaxLevel    = 20
	MaxScore    = 999999999
	MaxLines    = 9999
	MaxPieces   = 100000
	MaxDuration = 24 * time.Hour
	FixedStep   = time.Second / 60
)

// Phase is the game state.
type Phase uint8

// The game states.
const (
	PhaseReady Phase = iota
	PhaseRunning
	PhasePaused
	PhaseOver
)

// Cell is one board square; zero is empty, any other value is a Piece.
type Cell uint8

// Board is the playfield, row 0 at the top.
type Board [Rows][Cols]Cell

// Piece is one of the seven tetrominoes.
type Piece uint8

// The piece kinds.
const (
	PieceI Piece = 1 + iota
	PieceO
	PieceT
	PieceS
	PieceZ
	PieceJ
	PieceL
)

// PieceCount is the number of tetromino kinds.
const PieceCount = 7

// Rotation is one of four piece orientations, clockwise from spawn.
type Rotation uint8

// The four orientations.
const (
	Rot0 Rotation = iota
	RotR
	Rot2
	RotL
)

// Point is a board cell or a box-relative cell.
type Point struct{ X, Y int }

// Action is one player command the input adapter emits.
type Action uint8

// The player actions.
const (
	ActionNone Action = iota
	ActionLeft
	ActionRight
	ActionRotateCW
	ActionRotateCCW
	ActionSoftDrop
	ActionHardDrop
	ActionPause
	ActionRestart
	ActionStart
)
