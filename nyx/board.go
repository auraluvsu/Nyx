package nyx

var board [8][8]*Piece

func newPiece(p pieceType, c colour) *Piece {
	return &Piece{
		Type:     p,
		Colour:   c,
		HasMoved: false,
	}
}
func setupBoard()
