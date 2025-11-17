package engine

import (
	//	"fmt"
	"math"

	nyx "auraluvsu.com/nyx/logic"
)

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
			case nyx.Knight:
				if p.Colour == nyx.White {
					val += knightPST[rank][file]
				} else {
					val += knightPST[7-rank][file]
				}
			case nyx.Bishop:
				if p.Colour == nyx.White {
					val += bishopPST[rank][file]
				} else {
					val += bishopPST[7-rank][file]
				}
			case nyx.Rook:
				if p.Colour == nyx.White {
					val += rookPST[rank][file]
				} else {
					val += rookPST[7-rank][file]
				}
			case nyx.Queen:
				if p.Colour == nyx.White {
					val += queenPST[rank][file]
				} else {
					val += queenPST[7-rank][file]
				}
			case nyx.King:
				table := kingMidPST
				if p.Colour == nyx.White {
					val += table[rank][file]
				} else {
					val += table[7-rank][file]
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

func genMoves(c nyx.Colour, board [8][8]*nyx.Piece, enPassant *nyx.Position, yield func(fx, fy, tx, ty int, p *nyx.Piece, cap *nyx.Piece, nextEnPassant *nyx.Position)) {
	for fromX := range 8 {
		for fromY := range 8 {
			p := board[fromX][fromY]
			if p == nil || p.Colour != c {
				continue
			}
			for toX := range 8 {
				for toY := range 8 {
					val, err := p.IsValidMove(fromX, fromY, toX, toY, board, enPassant)
					if err != nil || !val {
						continue
					}
					if p.Type == nyx.Pawn && (toY == 0 || toY == 7) {
						promotions := []nyx.PieceType{
							nyx.Queen,
							nyx.Bishop,
							nyx.Rook,
							nyx.Knight,
						}
						captured := board[toX][toY]
						board[toX][toY], board[fromX][fromY] = p, nil
						originalType := p.Type
						for _, promo := range promotions {
							p.Type = promo
							if !nyx.IsInCheck(c, board) {
								yield(fromX, fromY, toX, toY, p, captured, nil)
							}
						}
						p.Type = originalType
						board[fromX][fromY], board[toX][toY] = p, captured
						continue
					}
					temp := board[toX][toY]
					board[toX][toY], board[fromX][fromY] = p, nil
					isEnPassant := p.Type == nyx.Pawn && enPassant != nil &&
						toX == enPassant.X && toY == enPassant.Y && temp == nil &&
						int(math.Abs(float64(toX-fromX))) == 1
					var capturedPawn *nyx.Piece
					var captureY int
					if isEnPassant {
						step := 1
						if p.Colour == nyx.White {
							step = -1
						}
						captureY = toY - step
						capturedPawn = board[toX][captureY]
						board[toX][captureY] = nil
					}
					if !nyx.IsInCheck(c, board) {
						var nextEnPassant *nyx.Position
						if p.Type == nyx.Pawn && int(math.Abs(float64(toY-fromY))) == 2 {
							nextEnPassant = &nyx.Position{X: fromX, Y: (fromY + toY) / 2}
						}
						yield(fromX, fromY, toX, toY, p, temp, nextEnPassant)
					}
					board[fromX][fromY], board[toX][toY] = p, temp
					if isEnPassant {
						board[toX][captureY] = capturedPawn
					}
				}
			}
		}
	}
}

func Minimax(board [8][8]*nyx.Piece, depth int, alpha, beta int, maxColour, turn nyx.Colour, enPassant *nyx.Position) int {
	noMoves := !nyx.HasAnyLegalMoves(turn, board, enPassant)
	if depth == 0 || noMoves {
		if noMoves {
			if nyx.IsInCheck(turn, board) {
				if turn == maxColour {
					return -99999 + depth
				}
				return 99999 - depth
			}
			return 0
		}
		return Evaluate(board, maxColour)
	}

	if maxColour == turn {
		best := math.MinInt32
		genMoves(turn, board, enPassant, func(fx, fy, tx, ty int, p, cap *nyx.Piece, nextEnPassant *nyx.Position) {
			board[tx][ty], board[fx][fy] = p, nil
			score := Minimax(board, depth-1, alpha, beta, maxColour, nyx.OppositeColour(turn), nextEnPassant)
			board[fx][fy], board[tx][ty] = p, cap
			if score > best {
				best = score
			}
			if best > alpha {
				alpha = best
			}
			if alpha >= beta {
				return
			}
		})
		return best
	}
	best := math.MaxInt32
	genMoves(turn, board, enPassant, func(fx, fy, tx, ty int, p, cap *nyx.Piece, nextEnPassant *nyx.Position) {
		board[tx][ty], board[fx][fy] = p, nil
		score := Minimax(board, depth-1, alpha, beta, maxColour, nyx.OppositeColour(turn), nextEnPassant)
		board[fx][fy], board[tx][ty] = p, cap

		if score < best {
			best = score
		}
		if best < beta {
			beta = best
		}
		if beta <= alpha {
			return
		}
	})
	return best
}
