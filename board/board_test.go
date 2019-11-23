package board

import "testing"

var emptyBoard = Board{
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
}

func TestNewBoard(t *testing.T) {
	expected := &emptyBoard
	got := NewBoard()
	if !expected.Equal(got) {
		t.Errorf("expected \n%v\n, got \n%v\n", expected, got)
	}
}

func TestEqual(t *testing.T) {
	boardOne := NewBoard()
	boardTwo := NewBoard()
	if !boardOne.Equal(boardTwo) {
		t.Error("expected board one and two to be equal, was false")
	}
}

func TestNotEqual(t *testing.T) {
	boardOne := NewBoard()
	boardTwo := NewBoard()
	(*boardOne)[0][0] = PlayerOne
	(*boardTwo)[0][0] = PlayerTwo
	if boardOne.Equal(boardTwo) {
		t.Error("expected board one and two to be not equal, was true")
	}
}

var validMovesTests = []struct {
	board      Board
	validMoves []Move
}{
	{
		Board{
			{1, 0, 0, 0, 0, 0, 2},
			{1, 2, 0, 0, 0, 0, 1},
			{2, 1, 1, 0, 0, 0, 2},
			{1, 2, 2, 2, 0, 0, 1},
			{2, 1, 2, 1, 1, 0, 2},
			{1, 2, 1, 2, 1, 2, 1},
		},
		[]Move{1, 2, 3, 4, 5},
	},
	{
		Board{
			{1, 1, 2, 2, 1, 1, 2},
			{1, 2, 1, 2, 1, 2, 1},
			{2, 1, 1, 1, 2, 2, 2},
			{1, 2, 2, 2, 1, 2, 1},
			{2, 1, 2, 1, 1, 1, 2},
			{1, 2, 1, 2, 1, 2, 1},
		},
		[]Move{},
	},
	{
		Board{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
		[]Move{0, 1, 2, 3, 4, 5, 6},
	},
}

func TestValidMoves(t *testing.T) {
	for _, test := range validMovesTests {
		expected := test.validMoves
		got := test.board.ValidMoves()
		if !equal(got, expected) {
			t.Errorf("expected %v and %v to be equal, was false", got, expected)
		}
	}
}

func equal(a, b []Move) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type playerMove struct {
	player Field
	move   Move
	winner Field
}

var playMoveTests = []struct {
	before      Board
	playerMoves []playerMove
	after       Board
}{
	{
		Board{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
		[]playerMove{
			{PlayerOne, 0, 0},
			{PlayerTwo, 1, 0},
			{PlayerOne, 0, 0},
			{PlayerTwo, 1, 0},
			{PlayerOne, 1, 0},
			{PlayerTwo, 0, 0},
			{PlayerOne, 2, 0},
			{PlayerTwo, 2, 0},
			{PlayerOne, 2, 0},
			{PlayerTwo, 3, 0},
			{PlayerOne, 3, 0},
			{PlayerTwo, 3, 0},
			{PlayerOne, 4, 0},
			{PlayerTwo, 4, 0},
			{PlayerOne, 1, 1},
		},
		Board{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0},
			{2, 1, 1, 2, 0, 0, 0},
			{1, 2, 2, 1, 2, 0, 0},
			{1, 2, 1, 2, 1, 0, 0},
		},
	},
}

func TestPlayMoves(t *testing.T) {
	for _, test := range playMoveTests {
		board := test.before
		for _, move := range test.playerMoves {
			b, winner, err := board.Play(move.move, move.player)
			if err != nil {
				t.Errorf("applied move %v to board \n%v\n: %v", move, board, err)
			}
			if winner != move.winner {
				t.Errorf("expected winner %d for board \n%v\n, got winner %d",
					move.winner, b, winner)
			}
			board = *b
		}
		got := board
		expected := &test.after
		if !got.Equal(expected) {
			t.Errorf("applying moves %v to board \n%v\n, expected \n%v\n, got \n%v\n",
				test.playerMoves, test.before, expected, got)
		}
	}
}
