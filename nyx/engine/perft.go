package engine

import (
	"fmt"
	"runtime"
	"sync"

	nyx "auraluvsu.com/nyx/logic"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func applyPerftMove(board [8][8]*nyx.Piece, fx, fy, tx, ty int, mover *nyx.Piece, enPassant *nyx.Position) [8][8]*nyx.Piece {
	newBoard := board
	if mover == nil {
		return newBoard
	}

	movedPiece := *mover
	movedPiece.HasMoved = true
	newBoard[fx][fy] = nil
	newBoard[tx][ty] = &movedPiece

	// Handle en passant captures (destination square empty in parent position)
	if enPassant != nil && movedPiece.Type == nyx.Pawn && tx == enPassant.X && ty == enPassant.Y && board[tx][ty] == nil {
		dir := 1
		if movedPiece.Colour == nyx.White {
			dir = -1
		}
		captureY := ty - dir
		newBoard[tx][captureY] = nil
	}

	// Handle rook movement for castling
	if movedPiece.Type == nyx.King && absInt(tx-fx) == 2 {
		var rookFrom, rookTo int
		if tx > fx {
			rookFrom = 7
			rookTo = tx - 1
		} else {
			rookFrom = 0
			rookTo = tx + 1
		}
		if rook := newBoard[rookFrom][fy]; rook != nil {
			rookClone := *rook
			rookClone.HasMoved = true
			newBoard[rookTo][fy] = &rookClone
			newBoard[rookFrom][fy] = nil
		}
	}

	return newBoard
}

// Perft takes the algorithm - in this case, the minimax algorithm - and runs int
// on a set depth to make sure the engine algorithm is working properly
func Perft(board [8][8]*nyx.Piece, depth int, turn nyx.Colour, enPassant *nyx.Position) int {
	if depth == 0 {
		return 1
	}
	workers := runtime.NumCPU()
	sem := make(chan struct{}, workers)
	var wg sync.WaitGroup
	var mu sync.Mutex
	nodes := 0
	genMoves(turn, board, enPassant, func(fx, fy, tx, ty int, mover *nyx.Piece, _ *nyx.Piece, nextEnPassant *nyx.Position) {
		wg.Add(1)
		sem <- struct{}{}
		childBoard := applyPerftMove(board, fx, fy, tx, ty, mover, enPassant)
		go func(b [8][8]*nyx.Piece, nextEP *nyx.Position) {
			defer wg.Done()
			defer func() { <-sem }()

			count := Perft(b, depth-1, nyx.OppositeColour(turn), nextEP)
			mu.Lock()
			nodes += count
			mu.Unlock()
		}(childBoard, nextEnPassant)
	})
	wg.Wait()
	return nodes
}

func PerftDivide(board [8][8]*nyx.Piece, depth int, turn nyx.Colour, enPassant *nyx.Position) {
	total := 0
	genMoves(turn, board, enPassant, func(fx, fy, tx, ty int, mover *nyx.Piece, _ *nyx.Piece, nextEnPassant *nyx.Position) {
		childBoard := applyPerftMove(board, fx, fy, tx, ty, mover, enPassant)
		nodes := Perft(childBoard, depth-1, nyx.OppositeColour(turn), nextEnPassant)
		moveStr := fmt.Sprintf("%c%d%c%d", 'a'+fx, fy+1, 'a'+tx, ty+1)
		fmt.Printf("%s: %d\n", moveStr, nodes)
		total += nodes
	})
	fmt.Printf("Total: %d\n", total)
}
