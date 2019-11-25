package player

import (
	"4iar/board"
	"log"
	"math/rand"
	"time"
)

// WinningMovePlayer is a player smart enough to detect and play winning moves.
type WinningMovePlayer struct {
	PlayerField board.Field
}

// NewWinningMovePlayer creates a new winning move player.
func NewWinningMovePlayer(field board.Field) *Player {
	winningMovePlayer := WinningMovePlayer{field}
	p := Player(&winningMovePlayer)
	rand.Seed(time.Now().Unix())
	return &p
}

// Play tries to find a winning move, or picks a random move, if no winning
// move is available.
func (p *WinningMovePlayer) Play(b *board.Board) *board.Move {
	candidates := b.ValidMoves()
	if len(candidates) == 0 {
		return nil
	}
	if len(candidates) == 1 {
		return &candidates[0]
	}
	for _, candidate := range candidates {
		_, outcome, err := b.Play(candidate, p.PlayerField)
		if err != nil {
			log.Printf("play move %v on board %v: %v", candidate, b, err)
			return nil
		}
		if int(outcome) == int(p.PlayerField) {
			return &candidate
		}
	}
	pick := rand.Intn(len(candidates))
	return &candidates[pick]
}

// Field returns the field assigned to the player.
func (p *WinningMovePlayer) Field() board.Field {
	return p.PlayerField
}
