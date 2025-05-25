package nyx

func newPiece(p pieceType, c colour) *Piece {
	return &Piece{
		Type:     p,
		Colour:   c,
		HasMoved: false,
	}
}

func setupBoard() [8][8]*Piece {
	var board [8][8]*Piece
	order := []pieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for i := range 8 {
		board[0][i] = newPiece(order[i], Black)
		board[1][i] = newPiece(Pawn, Black)
	}
	for i := range 8 {
		board[7][i] = newPiece(Pawn, White)
		board[6][i] = newPiece(order[i], White)
	}
	return board

}
