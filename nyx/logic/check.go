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
				check, err := p.IsValidMove(x, y, kingX, kingY, board)
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

func HasAnyLegalMoves(c Colour, board [8][8]*Piece) bool {
	for fromX := range 8 {
		for fromY := range 8 {
			p := board[fromX][fromY]
			if p != nil && p.Colour == c {
				for toX := range 8 {
					for toY := range 8 {
						val, err := p.IsValidMove(fromX, fromY, toX, toY, board)
						if err != nil {
							log.Fatal("Error getting valid moves")
						}
						if val {
							temp := board[toX][toY]
							board[toX][toY] = p
							board[fromX][fromY] = nil
							inCheck := IsInCheck(c, board)
							board[fromX][fromY] = p
							board[toX][toY] = temp
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

func GetGameState(c Colour, board [8][8]*Piece) GameState {
	inCheck := IsInCheck(c, board)
	hasMoves := HasAnyLegalMoves(c, board)
	if inCheck && !hasMoves {
		return Checkmate
	}

	if !inCheck && !hasMoves {
		return Stalemate
	}

	return Ongoing
}
