// Package game implements a game of two players against one another.
package game

import (
	"4iar/board"
	"4iar/player"
	"fmt"
)

// Game represents a game of two players against one another.
type Game struct {
	PlayerOne *player.Player
	PlayerTwo *player.Player
}

// NewGame creates a new game with the two players in the given order playing
// against one another.
func NewGame(playerOne, playerTwo *player.Player) *Game {
	g := Game{
		PlayerOne: playerOne,
		PlayerTwo: playerTwo,
	}
	return &g
}

// Play plays through the game until a winner is found, or the board has been
// filled without a player winning, in which case the game is tied. The outcome
// is returned.
func (g *Game) Play() (board.Outcome, error) {
	b := board.NewBoard()
	activePlayer := g.PlayerTwo
	finished := false
	for !finished {
		if activePlayer == g.PlayerOne {
			activePlayer = g.PlayerTwo
		} else {
			activePlayer = g.PlayerOne
		}
		validMoves := b.ValidMoves()
		move := (*activePlayer).Play(b)
		if !board.Contains(validMoves, *move) {
			return board.Undecided, fmt.Errorf("illgal move %v from player %v", move, activePlayer)
		}
		brd, outcome, err := b.Play(*move, (*activePlayer).Field())
		b = brd
		if err != nil {
			return board.Undecided, fmt.Errorf("apply move %v to board %v: %v", move, b, err)
		}
		fmt.Println(brd)
		if outcome != board.Undecided {
			return outcome, nil
		}
	}
	return board.Undecided, nil
}
