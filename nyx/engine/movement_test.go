package nyx

import "testing"

func emptyBoard() [8][8]*Piece {
	return [8][8]*Piece{}
}

func TestKnightMove(t *testing.T) {
	board := emptyBoard()
	knight := newPiece(Knight, White)
	board[4][4] = knight

	validMoves := [][2]int{
		{6, 5}, {6, 3}, {2, 5}, {2, 3}, {5, 6}, {5, 2}, {3, 6}, {3, 2},
	}
	for _, move := range validMoves {
		if !knight.IsValidKnightMove(4, 4, move[0], move[1], board) {
			t.Errorf("Knight should be able to move to %v", move)
		}
	}

	invalidMoves := [][2]int{
		{5, 5}, {4, 6}, {4, 3}, {7, 4},
	}
	for _, move := range invalidMoves {
		if knight.IsValidKnightMove(4, 4, move[0], move[1], board) {
			t.Errorf("Knight should NOT be able to move to %v", move)
		}
	}
}

func TestBishopMove(t *testing.T) {
	board := emptyBoard()
	bishop := &Piece{Type: Bishop, Colour: White, HasMoved: false}
	board[3][3] = bishop

	validMoves := [][2]int{
		{0, 0}, {0, 6}, {6, 6}, {6, 0},
	}
	for _, move := range validMoves {
		if !bishop.IsValidKnightMove(4, 4, move[0], move[1], board) {
			t.Errorf("Bishop should be able to move to %v", move)
		}
	}

	invalidMoves := [][2]int{
		{3, 6}, {0, 3}, {4, 6}, {5, 5},
	}
	board[4][4] = &Piece{Type: Knight, Colour: White, HasMoved: false}
	for _, move := range invalidMoves {
		if bishop.IsValidKnightMove(4, 4, move[0], move[1], board) {
			t.Errorf("Bishop should NOT be able to move to %v", move)
		}
	}
}
