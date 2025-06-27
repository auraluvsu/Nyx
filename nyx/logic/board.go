package nyx

import "fmt"

func newPiece(p PieceType, c Colour) *Piece {
	return &Piece{
		Type:     p,
		Colour:   c,
		HasMoved: false,
	}
}

func OppositeColour(c Colour) Colour {
	if c == White {
		return Black
	}
	return White
}

func SetupBoard() [8][8]*Piece {
	var board [8][8]*Piece
	order := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for i := range 8 {
		board[i][0] = newPiece(order[i], Black)
		board[i][1] = newPiece(Pawn, Black)
		board[i][7] = newPiece(order[i], White)
		board[i][6] = newPiece(Pawn, White)
	}
	return board
}

func DebugPrintBoard(board [8][8]*Piece) {
	for y := 7; y >= 0; y-- {
		fmt.Printf("%d ", y+1)
		for x := range 8 {
			p := board[x][y]
			if p == nil {
				fmt.Print("  .  ")
			} else {
				sym := pieceSymbol(p)
				fmt.Print(sym + " ")
			}
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("  a  b  c  d  e  f  g  h  ")
}
