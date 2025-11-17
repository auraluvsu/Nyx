package nyx

import "log"

func IsInCheck(c Colour, board [8][8]*Piece) bool {
	var kingX, kingY int
	found := false
	for x := range 8 {
		for y := range 8 {
			p := board[x][y]
			if p != nil && p.Type == King && p.Colour == c {
				kingX, kingY = x, y
				found = true
				break
			}
		}
	}
	if !found {
		return false
	}
	for x := range 8 {
		for y := range 8 {
			p := board[x][y]
			if p != nil && p.Colour != c {
				check, err := p.IsValidMove(x, y, kingX, kingY, board, nil)
				if err != nil {
					log.Fatal("Error processing check")
				}
				if check {
					return true
				}
			}
		}
	}
	return false
}

func HasAnyLegalMoves(c Colour, board [8][8]*Piece, enPassant *Position) bool {
	for fromX := range 8 {
		for fromY := range 8 {
			p := board[fromX][fromY]
			if p != nil && p.Colour == c {
				for toX := range 8 {
					for toY := range 8 {
						val, err := p.IsValidMove(fromX, fromY, toX, toY, board, enPassant)
						if err != nil {
							log.Fatal("Error getting valid moves")
						}
						if val {
							temp := board[toX][toY]
							board[toX][toY], board[fromX][fromY] = p, nil
							isEnPassant := p.Type == Pawn && enPassant != nil &&
								toX == enPassant.X && toY == enPassant.Y && temp == nil
							var capturedPawn *Piece
							var captureX, captureY int
							if isEnPassant {
								step := 1
								if p.Colour == White {
									step = -1
								}
								captureX = toX
								captureY = toY - step
								capturedPawn = board[captureX][captureY]
								board[captureX][captureY] = nil
							}
							inCheck := IsInCheck(c, board)
							board[fromX][fromY], board[toX][toY] = p, temp
							if isEnPassant {
								board[captureX][captureY] = capturedPawn
							}
							if !inCheck {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

type GameState string

const (
	Ongoing   GameState = "ongoing"
	Checkmate GameState = "checkmate"
	Stalemate GameState = "stalemate"
)

func GetGameState(c Colour, board [8][8]*Piece, enPassant *Position) GameState {
	inCheck := IsInCheck(c, board)
	hasMoves := HasAnyLegalMoves(c, board, enPassant)
	if inCheck && !hasMoves {
		return Checkmate
	}

	if !inCheck && !hasMoves {
		return Stalemate
	}

	return Ongoing
}
