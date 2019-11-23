// Package board represents a board for the game Four in a Row.
package board

const (
	// Rows is the number of rows on the board.
	Rows = 6
	// Cols is the number of columns on the board.
	Cols = 7

	// Empty represents an empty, i.e. unplayed field.
	Empty = 0
	// PlayerOne represents the field value for the first player.
	PlayerOne = 1
	// PlayerTwo represents the field value for the second player.
	PlayerTwo = 2
)

// Board is a two-dimensional array of integer values, representing the fields
// of a game board.
type Board [][]int

// NewBoard creates a new,  empty board, i.e. a board where all fields have the
// value Empty.
func NewBoard() *Board {
	board := Board(make([][]int, Rows))
	for r := 0; r < Rows; r++ {
		board[r] = make([]int, Cols)
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
