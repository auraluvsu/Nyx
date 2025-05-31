package nyx

import (
	"fmt"
)

type Piece struct {
	Type     PieceType // Each piece has a specific type assigned to it e.g. Knight
	Colour   colour    //Playing for white or black
	HasMoved bool      //Making sure turn based logic is working
}

type Move struct {
	Piece  PieceType
	Fx, Fy *int
	Tx, Ty int // Tx (To x coordinate), Ty (To y coordinate)
}

type colour string
type PieceType string

const (
	White colour = "White"
	Black colour = "Black"
)

const ( // Piece Types
	Pawn   PieceType = "Pawn"
	Knight PieceType = "Knight"
	Bishop PieceType = "Bishop"
	Rook   PieceType = "Rook"
	King   PieceType = "King"
	Queen  PieceType = "Queen"
)

func makeBoard() [][]uint8 {
	board := make([][]uint8, 8) // Outer Board
	for i := range 8 {
		board[i] = make([]uint8, 8) // Inner Board
	}
	return board
}

func chessToIndex(pos string) (x, y int, err error) {
	if len(pos) != 2 {
		return -1, -1, fmt.Errorf("Error! Invalid position")
	}
	file := pos[0] - 'a'
	rank := pos[1] - '1'
	if file < 0 || file >= 8 || rank < 0 || rank >= 8 {
		return -1, -1, fmt.Errorf("Error: out of bounds position")
	}
	return int(file), 7 - int(rank), nil
}

func indexToChess(x, y int) (string, error) {
	if x >= 8 || y >= 8 || x < 0 || y < 0 {
		return "", fmt.Errorf("Error! Invalid index")
	}
	file := rune('a' + x)
	rank := rune('8' - y)
	return string([]rune{file, rank}), nil
}

func pieceSymbol(p *Piece) string {
	switch p.Type {
	case Rook:
		if p.Colour == White {
			return "♖ "
		}
		return "♜ "

	case Knight:
		if p.Colour == White {
			return "♘ "
		}
		return "♞ "

	case Bishop:
		if p.Colour == White {
			return "♗ "
		}
		return "♝ "

	case Queen:
		if p.Colour == White {
			return "♕ "
		}
		return "♕ "

	case King:
		if p.Colour == White {
			return "♔ "
		}
		return "♕ "

	case Pawn:
		if p.Colour == White {
			return "♙ "
		}
		return "♙ "

	default:
		return "."
	}
}
