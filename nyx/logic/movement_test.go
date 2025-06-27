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
		{6, 5},
		{6, 3},
		{2, 5},
		{2, 3},
		{5, 6},
		{5, 2},
		{3, 6},
		{3, 2},
	}
	for _, move := range validMoves {
		if !knight.IsValidKnightMove(4, 4, move[0], move[1], board) {
			t.Errorf("Knight should be able to move to %v", move)
		}
	}

	invalidMoves := [][2]int{
		{5, 5},
		{4, 6},
		{4, 3},
		{7, 4},
	}
	for _, move := range invalidMoves {
		if knight.IsValidKnightMove(4, 4, move[0], move[1], board) {
			t.Errorf("Knight should NOT be able to move to %v", move)
		}
	}
}

func TestBishopMove(t *testing.T) {
	board := emptyBoard()
	bishop := newPiece(Bishop, White)
	board[3][3] = bishop

	validMoves := [][2]int{
		{0, 0},
		{0, 6},
		{6, 6},
		{6, 0},
	}
	for _, move := range validMoves {
		if !bishop.IsValidBishopMove(3, 3, move[0], move[1], board) {
			t.Errorf("Bishop should be able to move to %v", move)
		}
	}
}
func TestInvalidBishopMove(t *testing.T) {
	board := emptyBoard()
	bishop := newPiece(Bishop, White)
	board[3][3] = bishop
	invalidMoves := [][2]int{
		{3, 6},
		{0, 3},
		{4, 6},
		{5, 5},
	}
	board[4][4] = newPiece(Pawn, White)
	for _, move := range invalidMoves {
		if bishop.IsValidBishopMove(3, 3, move[0], move[1], board) {
			t.Errorf("Bishop should NOT be able to move to %v", move)
		}
	}
}

func TestRookValidMoves(t *testing.T) {
	board := emptyBoard()
	rook := newPiece(Rook, White)
	board[4][4] = rook
	validMoves := [][2]int{
		{4, 0},
		{4, 7},
		{0, 4},
		{7, 4},
	}

	for _, move := range validMoves {
		toX, toY := move[0], move[1]
		if !rook.IsValidRookMove(4, 4, toX, toY, board) {
			t.Errorf("Expected move to (%d, %d) to be valid", toX, toY)
		}
	}
}

func TestRookInvalidMoves(t *testing.T) {
	board := emptyBoard()
	rook := newPiece(Rook, White)
	board[4][4] = rook
	board[4][6] = newPiece(Queen, White)
	validMoves := [][2]int{
		{5, 5},
		{3, 3},
		{4, 7},
		{4, 4},
	}

	for _, move := range validMoves {
		toX, toY := move[0], move[1]
		if rook.IsValidRookMove(4, 4, toX, toY, board) {
			t.Errorf("Expected move to (%d, %d) to be invalid", toX, toY)
		}
	}
}

func TestPawnValidMoves(t *testing.T) {
	board := emptyBoard()

	pawn := newPiece(Pawn, White)
	board[4][6] = pawn

	board[3][5] = newPiece(Bishop, Black)
	if !pawn.IsValidPawnMove(4, 6, 4, 5, board) {
		t.Errorf("Expected white pawn to move forward one square")
	}
	if !pawn.IsValidPawnMove(4, 6, 4, 4, board) {
		t.Errorf("Expected white pawn to move forward two squares from start")
	}
	if !pawn.IsValidPawnMove(4, 6, 3, 5, board) {
		t.Errorf("Expected white pawn to capture diagonally at d3")
	}
}

func TestPawnInvalidMoves(t *testing.T) {
	board := emptyBoard()

	// White pawn at e2 (6,4)
	pawn := newPiece(Pawn, White)
	board[6][4] = pawn

	// Friendly piece in front
	board[5][4] = newPiece(Queen, White)

	invalidMoves := [][2]int{
		{5, 4}, // Blocked straight move
		{6, 4}, // No movement
		{5, 3}, // Diagonal but no enemy
		{4, 4}, // Two squares but path blocked
		{7, 4}, // Backwards
	}

	for _, move := range invalidMoves {
		toX, toY := move[0], move[1]
		if pawn.IsValidPawnMove(6, 4, toX, toY, board) {
			t.Errorf("Expected move to (%d,%d) to be invalid", toX, toY)
		}
	}
}

func TestQueenMoves(t *testing.T) {
	board := emptyBoard()
	queen := newPiece(Queen, White)
	board[3][3] = queen
	validMoves := [][2]int{
		{3, 0},
		{3, 7},
		{0, 3},
		{7, 3},
		{0, 0},
		{6, 6},
		{0, 6},
		{6, 0},
	}

	for _, move := range validMoves {
		if !queen.IsValidQueenMove(3, 3, move[0], move[1], board) {
			t.Errorf("Expected queen move to (%d, %d) to be valid", move[0], move[1])
		}
	}
	board[3][5] = newPiece(Pawn, White)
	if queen.IsValidQueenMove(3, 3, 3, 6, board) {
		t.Errorf("Expected queen move to (%d, %d) to be invalid due to obstruction", 3, 6)
	}
	board[2][2] = newPiece(Bishop, Black)
	if !queen.IsValidQueenMove(3, 3, 2, 2, board) {
		t.Errorf("Expected queen move to (%d, %d) to be valid due to capture laws", 3, 6)
	}
	board[4][4] = newPiece(Bishop, White)
	if queen.IsValidQueenMove(3, 3, 4, 4, board) {
		t.Errorf("Expected queen move to (%d, %d) to be valid due to capture laws", 4, 4)
	}

}

func TestKingMoves(t *testing.T) {
	board := emptyBoard()
	king := newPiece(King, White)
	board[4][4] = king
	validMoves := [][2]int{
		{3, 3},
		{3, 4},
		{3, 5},
		{4, 3},
		{4, 5},
		{5, 3},
		{5, 5},
	}
	for _, move := range validMoves {
		if !king.IsValidKingMove(4, 4, move[0], move[1], board) {
			t.Errorf("Expected king move to (%d, %d) to be valid", move[0], move[1])
		}
	}
	invalidMoves := [][2]int{
		{2, 2},
		{4, 6},
		{6, 4},
		{0, 0},
	}
	for _, move := range invalidMoves {
		if king.IsValidKingMove(4, 4, move[0], move[1], board) {
			t.Errorf("Expected king move to (%d, %d) to be invalid", move[0], move[1])
		}
	}

}
