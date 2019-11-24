package player

import (
	"4iar/board"
	"math/rand"
	"time"
)

// RandomPlayer is a player that plays random moves.
type RandomPlayer struct {
	PlayerField board.Field
}

// NewRandomPlayer creates a new random player.
func NewRandomPlayer(field board.Field) *Player {
	randomPlayer := RandomPlayer{field}
	p := Player(&randomPlayer)
	rand.Seed(time.Now().Unix())
	return &p
}

// Play picks a random move.
func (p *RandomPlayer) Play(b *board.Board) *board.Move {
	candidates := b.ValidMoves()
	if len(candidates) == 0 {
		return nil
	}
	if len(candidates) == 1 {
		return &candidates[0]
	}
	pick := rand.Intn(len(candidates))
	return &candidates[pick]
}

// Field returns the field assigned to the player.
func (p *RandomPlayer) Field() board.Field {
	return p.PlayerField
}
