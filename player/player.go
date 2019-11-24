// Package player contains both the interface description for a player, and
// different player implementations.
package player

import "4iar/board"

// Player describes a player that is able to play a move on the given board.
type Player interface {
	// Play returns a move with the given field to be applied to the board.
	Play(b *board.Board) *board.Move

	// Field returns the field assigned to the player.
	Field() board.Field
}
