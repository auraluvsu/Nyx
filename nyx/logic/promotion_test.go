package nyx

import "testing"

func emptyBoard() [8][8]*Piece {
	var board [8][8]*Piece
	return board
}
func TestPawnMove(t *testing.T) {
	board := emptyBoard()
	pawn := newPiece(Pawn, White)
	board[4][1] = pawn
	q := Queen
	move := &Move{
		Piece:     Pawn,
		Tx:        4,
		Ty:        0,
		PromoteTo: &q,
	}
	board[move.Tx][move.Ty] = pawn
	pawn.Promote(move)
	board[4][1] = nil
	if pawn.Type != Queen {
		t.Errorf("Pawn should promote to queen but is %s", pawn.Type)
	}
}
