package engine

import (
	"log"
	"math"

	nyx "auraluvsu.com/nyx/logic"
)

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

func Evaluate(board [8][8]*nyx.Piece, side nyx.Colour) int {
	score := 0
	for file := range 8 {
		for rank := range 8 {
			p := board[file][rank]
			if p == nil {
				continue
			}
			val := material[p.Type]
			switch p.Type {
			case nyx.Pawn:
				if p.Colour == nyx.White {
					val += pawnPST[rank][file]
				} else {
					val += pawnPST[7-rank][file]
				}
			}
			if p.Colour == side {
				score += val
			} else {
				score -= val
			}
		}
	}
	return score
}

func genMoves(c nyx.Colour, board [8][8]*nyx.Piece, yield func(fx, fy, tx, ty int, p *nyx.Piece, cap *nyx.Piece)) {
	for fromX := range 8 {
		for fromY := range 8 {
			p := board[fromX][fromY]
			if p == nil || p.Colour != c {
				continue
			}
			for toX := range 8 {
				for toY := range 8 {
					val, err := p.IsValidMove(fromX, fromY, toX, toY, board)
					if err != nil {
						log.Fatal("Error getting valid moves")
					}
					if !val {
						continue
					}
					temp := board[toX][toY]
					board[toX][toY] = p
					board[fromX][fromY] = nil
					inCheck := nyx.IsInCheck(c, board)
					board[fromX][fromY] = p
					board[toX][toY] = temp
					if !inCheck {
						yield(fromX, fromY, toX, toY, p, temp)
					}
				}
			}
		}
	}
}

func minimax(board [8][8]*nyx.Piece, depth int, alpha, beta int, maxColour, turn nyx.Colour) int {
	if depth == 0 {
		return Evaluate(board, maxColour)
	}

	if maxColour == turn {
		best := math.MinInt32
		genMoves(turn, board, func(fx, fy, tx, ty int, p, cap *nyx.Piece) {
			board[tx][ty], board[fx][fy] = p, nil
			score := minimax(board, depth-1, alpha, beta, maxColour, nyx.OppositeColour(turn))
			board[fx][fy], board[tx][ty] = p, cap
			if score > best {
				best = score
			}
			alpha = max(alpha, best)
		})
		return best
	}
	best := math.MaxInt32
	genMoves(turn, board, func(fx, fy, tx, ty int, p, cap *nyx.Piece) {
		board[tx][ty], board[fx][fy] = p, nil
		score := minimax(board, depth-1, alpha, beta, maxColour, nyx.OppositeColour(turn))
		board[fx][fy], board[tx][ty] = p, cap
		if score > best {
			best = score
		}
		best = min(best, score)
		beta = min(beta, score)
	})
	return best
}
