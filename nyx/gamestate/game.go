package gamestate

import (
	"fmt"

	nyx "auraluvsu.com/nyx/logic"
	"auraluvsu.com/nyx/parsing"
)

func Game() {
	fmt.Println("Welcome to Gochess!")
	board := nyx.SetupBoard()
	turn := nyx.White
	for {
		nyx.DebugPrintBoard(board)
		if nyx.IsInCheck(turn, board) {
			fmt.Printf("%s is in check\n", turn)
			if !nyx.HasAnyLegalMoves(turn, board) {
				fmt.Printf("Checkmate! %s wins", nyx.OppositeColour(turn))
				break
			}
		} else if !nyx.HasAnyLegalMoves(turn, board) {
			fmt.Println("Stalemate! It's a draw.")
			break
		}
		fmt.Printf("%s to move: \n", turn)
		var moveStr string
		fmt.Scan(&moveStr)
		move, err := parsing.ParseSAN(moveStr, turn)
		if move.IsCastle {
			kingX := 4
			kingY := 7
			rookFromX, rookToX := 7, 5
			rookY := 7
			if turn == nyx.Black {
				kingY = 0
				rookY = 0
			}
			if move.Tx == 2 {
				rookFromX = 0
				rookToX = 3
			}

			board[move.Tx][kingY] = board[kingX][kingY]
			board[kingX][kingY] = nil

			board[rookToX][rookY] = board[rookFromX][rookY]
			board[rookFromX][rookY] = nil
			board[move.Tx][kingY].HasMoved = true
			board[rookToX][rookY].HasMoved = true

			return
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		found := false
		for fromX := 0; fromX < 8 && !found; fromX++ {
			for fromY := 0; fromY < 8 && !found; fromY++ {
				piece := board[fromX][fromY]
				if piece != nil && piece.Colour == turn && piece.Type == move.Piece {
					fmt.Printf(
						"Checking piece %v at %d,%d for move to %d,%d (turn: %v)\n",
						piece.Type, fromX, fromY, move.Tx, move.Ty, turn)
					val, err := piece.IsValidMove(fromX, fromY, move.Tx, move.Ty, board)
					if err != nil {
						fmt.Println(err)
						continue
					}
					if val {
						board[move.Tx][move.Ty] = piece
						board[fromX][fromY] = nil
						found = true
					}
					if nyx.IsInCheck(turn, board) {
						fmt.Println("Error: You are in check.")
						board[move.Tx][move.Ty] = nil
						board[fromX][fromY] = piece
						found = false
					}
				}

			}
		}
		if !found {
			fmt.Println("No legal piece found that can perform that move.")
			continue
		}
		turn = nyx.OppositeColour(turn)
	}
}
