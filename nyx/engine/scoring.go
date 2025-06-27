package engine

import nyx "auraluvsu.com/nyx/logic"

var material = map[nyx.PieceType]int{
	nyx.Pawn:   100,
	nyx.Bishop: 300,
	nyx.Knight: 300,
	nyx.Rook:   500,
	nyx.Queen:  900,
	nyx.King:   0,
}

var pawnPST = [8][8]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{50, 50, 50, 50, 50, 50, 50, 50},
	{10, 10, 20, 30, 30, 20, 10, 10},
	{5, 5, 10, 25, 25, 10, 5, 5},
	{0, 0, 0, 20, 20, 0, 0, 0},
	{5, -5, -10, 0, 0, -10, -5, 5},
	{5, 10, 10, -20, -20, 10, 10, 5},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

func Evaluate(board [8][8]*nyx.Piece, side nyx.Colour) {
	score := 0
}
