// Package board represents a board for the game Four in a Row.
package board

import "errors"

// Field is an integer representing a field's state.
type Field int

const (
	// Rows is the number of rows on the board.
	Rows = 6
	// Cols is the number of columns on the board.
	Cols = 7

	// Empty represents an empty, i.e. unplayed field.
	Empty = Field(0)
	// PlayerOne represents the field value for the first player.
	PlayerOne = Field(1)
	// PlayerTwo represents the field value for the second player.
	PlayerTwo = Field(2)
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
// with the index indicated by move to the value of player, and returns a new
// board. The original board is not modified in the process. If the move is
// illegal, an ErrorInvalidMove is returned.
func (b *Board) Play(move Move, player Field) (*Board, error) {
	validMoves := b.ValidMoves()
	if !contains(validMoves, move) {
		return nil, ErrorInvalidMove
	}
	newBoard := b.Copy()
	for row := len(*newBoard) - 1; row >= 0; row-- {
		if (*newBoard)[row][move] == Empty {
			(*newBoard)[row][move] = player
			break
		}
	}
	return newBoard, nil
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

func contains(moves []Move, move Move) bool {
	for _, m := range moves {
		if m == move {
			return true
		}
	}
	return false
}
