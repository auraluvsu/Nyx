package nyx

import (
	"math"
)

func colorCheck(mp *Piece, op [8][8]*Piece, toX, toY int) bool {
	target := op[toX][toY]
	if target != nil && target.Colour == mp.Colour {
		return false
	}
	return true
}

func inBounds(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func (p *Piece) DiagPawnMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if board[toX][toY] == nil {
		return false
	}
	return board[toX][toY].Colour == p.Colour
}
func (p *Piece) IsValidPawnMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if !inBounds(toX, toY) {
		return false
	}
	direction := -1
	startRow := 6
	if p.Colour == Black {
		direction = 1
		startRow = 1
	}
	if board[toX][toY] != nil || board[toX][fromY+direction] != nil {
		return false
	}
	if math.Abs(float64(toX-fromX)) != 1 {
		return false
	}
	if fromY == toX {
		if toY == fromY+direction && board[toX][toY] == nil {
			return true
		}
		if fromY == startRow && toY == fromY+2*direction && board[toX][toY] == nil && board[toX][fromY+direction] == nil {
			return true
		}
	}
	if math.Abs(float64(toX-fromX)) == 1 && toY == fromY+direction {
		if p.DiagPawnMove(fromX, fromY, toX, toY, board) {
			return true
		}
	}
	return false
}

func (p *Piece) IsValidKnightMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if !inBounds(toX, toY) {
		return false
	}
	dx := toX - fromX
	dy := toY - fromY

	absDx := int(math.Abs(float64(dx)))
	absDy := int(math.Abs(float64(dy)))

	if !colorCheck(p, board, toX, toY) {
		return false
	}
	if (absDx == 2 && absDy == 1) || (absDx == 1 && absDy == 2) {
		return true
	}
	return false
}

func (p *Piece) IsValidRookMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if !inBounds(toX, toY) {
		return false
	}
	if !colorCheck(p, board, toX, toY) {
		return false
	}
	if fromX != toX && fromY != toY {
		return false
	}
	if fromY == toY {
		for i := min(fromX, toX) + 1; i < max(fromX, toX); i++ {
			if board[i][fromY] != nil {
				return false
			}
		}
	} else {
		for i := min(fromY, toY) + 1; i < max(fromY, toY); i++ {
			if board[fromX][i] != nil {
				return false
			}
		}
	}
	return true
}

func (p *Piece) IsValidBishopMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if !inBounds(toX, toY) {
		return false
	}
	if !colorCheck(p, board, toX, toY) {
		return false
	}
	if math.Abs(float64(toX-fromX)) != math.Abs(float64(toY-fromY)) {
		return false
	}
	dx := 1
	if toX < fromX {
		dx = -1
	}
	dy := 1
	if toY < fromY {
		dy = -1
	}
	x, y := fromX+dx, fromY+dy
	for x != toX && y != toY {
		if board[x][y] != nil {
			return false
		}
		x += dx
		y += dy
	}
	return true
}

func (p *Piece) IsValidQueenMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	// Check if destination is in bounds
	if !inBounds(toX, toY) {
		return false
	}
	// Check if there is a piece on that destination and if its your piece
	if !colorCheck(p, board, toX, toY) {
		return false
	}
	// Check if your move is either a valid rook move or a valid bishop move
	if p.IsValidRookMove(fromX, fromY, toX, toY, board) ||
		p.IsValidBishopMove(fromX, fromY, toX, toY, board) {
		return true
	}
	// If the move is valid for neither piece, return false so you can't move
	return false
}

func (p *Piece) IsValidKingMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	// Check if destination is in bounds
	if !inBounds(toX, toY) {
		return false
	}
	// Check if there is a piece on that destination and if its your piece
	if !colorCheck(p, board, toX, toY) {
		return false
	}
	if math.Abs(float64(toX-fromX)) <= 1 || math.Abs(float64(toY-fromY)) <= 1 {
		return true
	}
	return false
}
