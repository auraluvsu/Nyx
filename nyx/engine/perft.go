package engine

import (
	"fmt"
	"runtime"
	"sync"

	nyx "auraluvsu.com/nyx/logic"
)

// Perft takes the algorithm - in this case, the minimax algorithm - and runs int
// on a set depth to make sure the engine algorithm is working properly
func Perft(board [8][8]*nyx.Piece, depth int, turn nyx.Colour) int {
	if depth == 0 {
		return 1
	}
	workers := runtime.NumCPU()
	sem := make(chan struct{}, workers)
	var wg sync.WaitGroup
	var mu sync.Mutex
	nodes := 0
	genMoves(turn, board, func(fx, fy, tx, ty int, p, cap *nyx.Piece) {
		wg.Add(1)
		sem <- struct{}{}
		go func(fx, fy, tx, ty int, p *nyx.Piece) {
			defer wg.Done()
			defer func() { <-sem }()

			var newBoard [8][8]*nyx.Piece
			for i, row := range board {
				for j, piece := range row {
					newBoard[i][j] = piece
				}
			}
			newBoard[tx][ty], newBoard[fx][fy] = p, nil

			count := Perft(newBoard, depth-1, nyx.OppositeColour(turn))
			mu.Lock()
			nodes += count
			mu.Unlock()
		}(fx, fy, tx, ty, p)
	})
	wg.Wait()
	return nodes
}

func PerftDivide(board [8][8]*nyx.Piece, depth int, turn nyx.Colour) {
	total := 0
	genMoves(turn, board, func(fx, fy, tx, ty int, p, cap *nyx.Piece) {
		var newBoard [8][8]*nyx.Piece
		for i, row := range board {
			for j, piece := range row {
				newBoard[i][j] = piece
			}
		}
		newBoard[tx][ty], newBoard[fx][fy] = p, nil
		nodes := Perft(newBoard, depth-1, nyx.OppositeColour(turn))
		moveStr := fmt.Sprintf("%c%d%c%d", 'a'+fx, fy+1, 'a'+tx, ty+1)
		fmt.Printf("%s: %d\n", moveStr, nodes)
		total += nodes
	})
	fmt.Printf("Total: %d\n", total)
}
