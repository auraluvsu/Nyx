package engine

import (
	"math"

	nyx "auraluvsu.com/nyx/logic"
)

type SearchResult struct {
	Move  nyx.Move
	Score int
	PV    []nyx.Move
	Nodes int
	Depth int
}

func BestMove(board [8][8]*nyx.Piece, depth int, turn nyx.Colour) SearchResult {
	best := math.MinInt32
	var bestMove nyx.Move
	nodes := 0

	genMoves(turn, board, func(fx, fy, tx, ty int, p, cap *nyx.Piece) {
		board[tx][ty], board[fx][fy] = p, nil
		score := Minimax(board, depth-1, math.MinInt32, math.MaxInt32, turn, nyx.OppositeColour(turn))
		board[fx][fy], board[tx][ty] = p, cap
		nodes++
		fxCopy, fyCopy := fx, fy
		if score > best {
			best = score
			bestMove = nyx.Move{Fx: &fxCopy, Fy: &fyCopy, Tx: tx, Ty: ty}
		}
	})
	return SearchResult{
		Move:  bestMove,
		Score: best,
		Nodes: nodes,
		Depth: depth,
	}

}
