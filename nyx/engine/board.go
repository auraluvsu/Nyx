package nyx

import "fmt"

func newPiece(p PieceType, c colour) *Piece {
	return &Piece{
		Type:     p,
		Colour:   c,
		HasMoved: false,
	}
}

func OppositeColour(c colour) colour {
	if c == White {
		return Black
	}
	return White
}

func SetupBoard() [8][8]*Piece {
	var board [8][8]*Piece
	order := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for i := range 8 {
		board[0][i] = newPiece(order[i], Black)
		board[1][i] = newPiece(Pawn, Black)
		board[6][i] = newPiece(Pawn, White)
		board[7][i] = newPiece(order[i], White)
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
	}
	fmt.Println(" a b c d e f g h")
}
