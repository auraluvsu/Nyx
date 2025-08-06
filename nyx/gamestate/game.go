package gamestate

import (
	"fmt"
	"math"

	nyx "auraluvsu.com/nyx/logic"
	"auraluvsu.com/nyx/parsing"
)

var enPassantPos *nyx.Position

func Game() {
	fmt.Println("Welcome to Gochess!")
	board := nyx.SetupBoard()
	turn := nyx.White
	for {
		fmt.Print("\033[2J\033[H")
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
		if err != nil {
			fmt.Println(err)
			continue
		}
		if move.IsCastle {
			kingX := 4
			kingY := 7 // White's back rank is 7
			if turn == nyx.Black {
				kingY = 0 // Black's back rank is 0
			}

			rookFromX, rookToX := 7, 5 // King-side castle
			if move.Tx == 2 {          // Queen-side castle
				rookFromX = 0
				rookToX = 3
			}

			// Move king
			board[move.Tx][kingY] = board[kingX][kingY]
			board[kingX][kingY] = nil
			board[move.Tx][kingY].HasMoved = true

			// Move rook
			rookY := kingY // Rook is on the same rank
			board[rookToX][rookY] = board[rookFromX][rookY]
			board[rookFromX][rookY] = nil
			board[rookToX][rookY].HasMoved = true

			turn = nyx.OppositeColour(turn)
			continue
		}
		found := false
		for fromX := 0; fromX < 8 && !found; fromX++ {
			for fromY := 0; fromY < 8 && !found; fromY++ {
				piece := board[fromX][fromY]
				if piece != nil && piece.Type == nyx.Pawn {
					finalRank := 0
					if piece.Colour == nyx.White {
						finalRank = 7
					}
					if move.Ty == finalRank && move.PromoteTo != nil {
						piece.Type = *move.PromoteTo
					}
				}
				if piece != nil && piece.Colour == turn && piece.Type == move.Piece {
					val, err := piece.IsValidMove(fromX, fromY, move.Tx, move.Ty, board, enPassantPos)
					if err != nil {
						fmt.Println(err)
						continue
					}
					if val {
						// Validate capture move
						targetPiece := board[move.Tx][move.Ty]
						isEnPassant := piece.Type == nyx.Pawn && enPassantPos != nil && move.Tx == enPassantPos.X && move.Ty == enPassantPos.Y
						if move.IsCapture && targetPiece == nil && !isEnPassant {
							fmt.Println("Invalid move: capture notation used but no piece to capture.")
							continue
						} else if !move.IsCapture && targetPiece != nil {
							fmt.Println("Invalid move: must use capture notation 'x'.")
							continue
						}

						// Handle en passant capture
						if piece.Type == nyx.Pawn && enPassantPos != nil && move.Tx == enPassantPos.X && move.Ty == enPassantPos.Y {
							var capturedPawnY int
							if piece.Colour == nyx.White {
								capturedPawnY = move.Ty + 1
							} else {
								capturedPawnY = move.Ty - 1
							}
							board[move.Tx][capturedPawnY] = nil
						}

						// Set enPassantPos for the next turn
						if piece.Type == nyx.Pawn && math.Abs(float64(fromY-move.Ty)) == 2 {
							enPassantPos = &nyx.Position{X: fromX, Y: (fromY + move.Ty) / 2}
						} else {
							enPassantPos = nil
						}

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
