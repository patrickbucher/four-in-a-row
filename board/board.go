// Package board represents a board for the game Four in a Row.
package board

import (
	"bytes"
	"errors"
)

// Field is an integer representing a field's state.
type Field int

// Outcome is an integer representing the outcome of a game.
type Outcome int

const (
	// Rows is the number of rows on the board.
	Rows = 6
	// Cols is the number of columns on the board.
	Cols = 7
	// Goal is the length of a row needed to win the game.
	Goal = 4

	// Empty represents an empty, i.e. unplayed field.
	Empty = Field(0)
	// PlayerOne represents the field value for the first player.
	PlayerOne = Field(1)
	// PlayerTwo represents the field value for the second player.
	PlayerTwo = Field(2)

	// Undecided is the outcome of a running game.
	Undecided = Outcome(-1)
	// Tie is the outcome of a game where with all fields filled, but no player won.
	Tie = Outcome(0)
	// PlayerOneWins is the outcome when player one won the game.
	PlayerOneWins = Outcome(1)
	// PlayerTwoWins is the outcome when player two won the game.
	PlayerTwoWins = Outcome(2)
)

// Board is a two-dimensional array of fields, representing the fields of a
// game board.
type Board [][]Field

// NewBoard creates a new,  empty board, i.e. a board where all fields have the
// value Empty.
func NewBoard() *Board {
	board := Board(make([][]Field, Rows))
	for r := 0; r < Rows; r++ {
		board[r] = make([]Field, Cols)
		for c := 0; c < Cols; c++ {
			board[r][c] = Empty
		}
	}
	return &board
}

// Equal compares two boards and returns true if both boards have the same
// dimensions and field values, and false otherwise.
func (b *Board) Equal(other *Board) bool {
	if len(*b) != len(*other) {
		return false
	}
	for r := 0; r < len(*b); r++ {
		if len((*b)[r]) != len((*other)[r]) {
			return false
		}
		for c := 0; c < len((*b)[r]); c++ {
			if (*b)[r][c] != (*other)[r][c] {
				return false
			}
		}
	}
	return true
}

// Move represents a column to be picked by a player in the range of [0;Cols).
type Move int

// ValidMoves returns a slice of moves that can be played, i.e. columns with an
// Empty field.
func (b *Board) ValidMoves() []Move {
	validMoves := make([]Move, 0)
	for col := 0; col < Cols; col++ {
		for row := 0; row < Rows; row++ {
			if (*b)[row][col] == Empty {
				validMoves = append(validMoves, Move(col))
				break
			}
		}
	}
	return validMoves
}

// ErrorInvalidMove indicates that a move has been tried to apply to a board,
// which doesn't allow for that move.
var ErrorInvalidMove = errors.New("illegal move")

// Play applies move of player, i.e. sets the topmost empty field in the column
// with the index indicated by move to the value of player, returns a new
// board, with the move applied, and a outcome value indicating the winner of
// the game. If the game is not over yet, Undecided is returned for the winner.
// If the game is over, but no player won, Tie is the outcome. The original
// board is not modified in the process. If the move is illegal, an
// ErrorInvalidMove is returned.
func (b *Board) Play(move Move, player Field) (*Board, Outcome, error) {
	validMoves := b.ValidMoves()
	if !Contains(validMoves, move) {
		return nil, -1, ErrorInvalidMove
	}
	newBoard := b.Copy()
	finalRow := -1
	for row := len(*newBoard) - 1; row >= 0; row-- {
		if (*newBoard)[row][move] == Empty {
			(*newBoard)[row][move] = player
			finalRow = row
			break
		}
	}
	return newBoard, newBoard.winner(finalRow, int(move)), nil
}

// Copy creates a copy B of the initial board A, so that A.Equal(B) holds true,
// but A == B doesn't.
func (b *Board) Copy() *Board {
	cpy := NewBoard()
	for r := 0; r < len(*b); r++ {
		for c := 0; c < len((*b)[r]); c++ {
			(*cpy)[r][c] = (*b)[r][c]
		}
	}
	return cpy
}

// String returns a string representation of the board.
func (b *Board) String() string {
	buf := bytes.NewBufferString("")
	for r := 0; r < len(*b); r++ {
		for c := 0; c < len((*b)[r]); c++ {
			buf.WriteRune(rune((*b)[r][c] + '0'))
			buf.WriteRune(' ')
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

type shift struct {
	v int
	h int
}

type direction int

const (
	north direction = iota
	northEast
	east
	southEast
	south
	southWest
	west
	northWest
)

var shifts = map[direction]shift{
	north:     shift{-1, 0},
	northEast: shift{-1, 1},
	east:      shift{0, 1},
	southEast: shift{1, 1},
	south:     shift{1, 0},
	southWest: shift{1, -1},
	west:      shift{0, -1},
	northWest: shift{-1, -1},
}

type coord struct {
	row int
	col int
}

func (c *coord) apply(s shift) {
	c.row += s.v
	c.col += s.h
}
func (c coord) inRange() bool {
	return c.row >= 0 && c.row < Rows && c.col >= 0 && c.col < Cols
}

// winner starts at the field (*b)[setRow][setCol], checks the board in all
// directions for fields of the same player, and returns the player's Field
// value, if found four fields in a row of that player.
func (b *Board) winner(setRow, setCol int) Outcome {
	chains := make(map[direction]int)
	playerValue := (*b)[setRow][setCol]
	for dir, sft := range shifts {
		f := &coord{row: setRow, col: setCol}
		var count int
		for count = 0; f.inRange(); f.apply(sft) {
			if (*b)[f.row][f.col] != playerValue {
				break
			}
			count++
		}
		chains[dir] = count
	}
	// origin was counted in both directions, remove one
	vertical := chains[north] + chains[south] - 1
	horizontal := chains[west] + chains[east] - 1
	upwards := chains[northEast] + chains[southWest] - 1
	downwards := chains[southEast] + chains[northWest] - 1
	if vertical >= Goal || horizontal >= Goal ||
		upwards >= Goal || downwards >= Goal {
		return Outcome(playerValue)
	}
	if b.hasEmptyFields() {
		return Undecided
	}
	return Tie
}

func (b *Board) hasEmptyFields() bool {
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if (*b)[r][c] == Empty {
				return true
			}
		}
	}
	return false
}

// Contains checks if move is contained in moves, returns true if so, and else
// otherwise.
func Contains(moves []Move, move Move) bool {
	for _, m := range moves {
		if m == move {
			return true
		}
	}
	return false
}
