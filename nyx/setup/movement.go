package nyx

import (
	"math"
)

func inBounds(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func (p *Piece) IsValidKnightMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if !inBounds(toX, toY) {
		return false
	}
	dx := toX - fromX
	dy := toY - fromY

	absDx := int(math.Abs(float64(dx)))
	absDy := int(math.Abs(float64(dy)))

	target := board[toX][toY]
	if target != nil && target.Colour == p.Colour {
		return false
	}
	if (absDx == 2 && absDy == 1) || (absDx == 1 && absDy == 2) {
		return true
	}
	return false
}

func (p *Piece) IsValidRookMove(fromX, toX, fromY, toY int, board [8][8]*Piece) bool {
	if !inBounds(toX, toY) {
		return false
	}
	dx := toX - fromX
	dy := toY - fromY

	absDx := int(math.Abs(float64(dx)))
	absDy := int(math.Abs(float64(dy)))
}
