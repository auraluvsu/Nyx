package nyx

import (
	"fmt"
	"math"
)

func colorCheck(mp *Piece, op [8][8]*Piece, toX, toY int) bool {
	target := op[toX][toY]
	if target != nil && target.Colour == mp.Colour {
		return false
	}
	return true
}

func InBounds(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func (p *Piece) DiagPawnMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	target := board[toX][toY]
	return target != nil && target.Colour != p.Colour
}

func sign(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func CanCastle(kx, ky int, kingSide bool, side Colour, board [8][8]*Piece) bool {
	var rookX int
	if kingSide {
		rookX = 7
	} else {
		rookX = 0
	}
	rook := board[rookX][ky]
	if rook == nil || rook.Type != Rook || rook.HasMoved || rook.Colour != side {
		return false
	}
	start := min(kx, rookX)
	end := max(kx, rookX)
	for x := start; x <= end; x++ {
		if board[x][ky] != nil {
			return false
		}
	}
	path := []int{kx, kx + sign(rookX-kx)}
	if kingSide {
		path = append(path, kx+2)
	} else {
		path = append(path, kx-2)
	}

	for _, x := range path {
		tmp := board[x][ky]
		board[x][ky], board[kx][ky] = board[kx][ky], nil
		inCheck := IsInCheck(side, board)
		board[kx][ky], board[x][ky] = board[x][ky], tmp
		if inCheck {
			return false
		}
	}
	return true
}

func (p *Piece) IsValidPawnMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if !InBounds(toX, toY) {
		return false
	}
	direction := 1
	startRow := 1
	if p.Colour == White {
		direction = -1
		startRow = 6
	}
	dx := toX - fromX
	dy := toY - fromY
	if dx == 0 {
		if dy == direction && board[toX][toY] == nil {
			return true
		}
		if fromY == startRow && dy == 2*direction &&
			board[toX][toY] == nil && board[toX][fromY+direction] == nil {
			return true
		}
	} else if math.Abs(float64(dx)) == 1 && dy == direction {
		return p.DiagPawnMove(fromX, fromY, toX, toY, board)
	}
	return false
}
func (p *Piece) IsValidKnightMove(fromX, fromY, toX, toY int, board [8][8]*Piece) bool {
	if !InBounds(toX, toY) {
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
	if !InBounds(toX, toY) {
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
	if !InBounds(toX, toY) {
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
	if !InBounds(toX, toY) {
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

func (p *Piece) IsValidKingMove(fromX, fromY, toX, toY int, board [8][8]*Piece) (bool, error) {
	// Check if destination is in bounds
	if !InBounds(toX, toY) {
		return false, nil
	}
	// Check if there is a piece on that destination and if its your piece
	if !colorCheck(p, board, toX, toY) {
		return false, nil
	}
	dx := int(math.Abs(float64(toX - fromX)))
	dy := int(math.Abs(float64(toY - fromY)))
	if dx <= 1 && dy <= 1 && (dx != 0 || dy != 0) {
		return true, nil
	}
	if dy != 0 || p.HasMoved || p.Type != King {
		return false, nil
	}
	if dx == 2 {
		return CanCastle(fromX, fromY, true, p.Colour, board), nil
	}
	if dx == 3 {
		return CanCastle(fromX, fromY, false, p.Colour, board), nil
	}
	return false, nil
}

func (p *Piece) IsValidMove(fromX, fromY, toX, toY int, board [8][8]*Piece) (bool, error) {
	switch p.Type {
	case Rook:
		return p.IsValidRookMove(fromX, fromY, toX, toY, board), nil

	case Bishop:
		return p.IsValidBishopMove(fromX, fromY, toX, toY, board), nil

	case Knight:
		return p.IsValidKnightMove(fromX, fromY, toX, toY, board), nil

	case Queen:
		return p.IsValidQueenMove(fromX, fromY, toX, toY, board), nil

	case King:
		val, err := p.IsValidKingMove(fromX, fromY, toX, toY, board)
		return val, err
	case Pawn:
		return p.IsValidPawnMove(fromX, fromY, toX, toY, board), nil

	default:
		return false, fmt.Errorf("Error! Not a valid piece")
	}
}
