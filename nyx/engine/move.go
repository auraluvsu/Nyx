package engine

import nyx "auraluvsu.com/nyx/logic"

type Undo struct {
	fromX, fromY, toX, toY int
	moved, captured        *nyx.Piece

	// Special cases
	wasPromotion bool
	prevType     nyx.PieceType

	wasCastle                 bool
	rookFromX, rookToX, rookY int
	rookPiece                 *nyx.Piece

	wasEnPassant     bool
	epCaptX, epCaptY int

	// State
	prevEnPassant *nyx.Position // copy of previous EP square (or nil)
	// If you rely on HasMoved:
	prevMovedFlag bool
	prevRookMoved bool
}

func makeMove(b *[8][8]*nyx.Piece, m nyx.Move, ep *nyx.Position) (Undo, *nyx.Position) {
	u := Undo{fromX: *m.Fx, fromY: *m.Fy, toX: m.Tx, toY: m.Ty, moved: b[*m.Fx][*m.Fy], captured: b[m.Tx][m.Ty], prevEnPassant: ep}
	piece := u.moved
	// Clear EP by default; set if double pawn push
	nextEP := (*nyx.Position)(nil)

	// En passant capture removal
	if piece.Type == nyx.Pawn && m.IsEnPassant {
		u.wasEnPassant = true
		if piece.Colour == nyx.White {
			u.epCaptX, u.epCaptY = m.Tx, m.Ty+1
		} else {
			u.epCaptX, u.epCaptY = m.Tx, m.Ty-1
		}
		u.captured = b[u.epCaptX][u.epCaptY]
		b[u.epCaptX][u.epCaptY] = nil
	}

	// Move piece
	b[m.Tx][m.Ty], b[*m.Fx][*m.Fy] = piece, nil

	// Promotion
	if piece.Type == nyx.Pawn && (m.Ty == 0 || m.Ty == 7) && m.PromoteTo != nil {
		u.wasPromotion = true
		u.prevType = piece.Type
		piece.Type = *m.PromoteTo
	}

	// Castling: also move rook
	if piece.Type == nyx.King && abs(m.Tx-m.Fx) == 2 {
		u.wasCastle = true
		u.rookY = m.Fy
		if m.Tx > m.Fx { // O-O
			u.rookFromX, u.rookToX = 7, 5
		} else { // O-O-O
			u.rookFromX, u.rookToX = 0, 3
		}
		u.rookPiece = b[u.rookFromX][u.rookY]
		b[u.rookToX][u.rookY], b[u.rookFromX][u.rookY] = u.rookPiece, nil
	}

	// Track HasMoved flags if you use them for castling legality
	u.prevMovedFlag = piece.HasMoved
	piece.HasMoved = true
	if u.rookPiece != nil {
		u.prevRookMoved = u.rookPiece.HasMoved
		u.rookPiece.HasMoved = true
	}

	// En passant target after double pawn push
	if piece.Type == Pawn && abs(m.Ty-m.Fy) == 2 {
		mid := (m.Ty + m.Fy) / 2
		nextEP = &Position{X: m.Fx, Y: mid}
	}

	return u, nextEP
}

func unmakeMove(b *[8][8]*Piece, u Undo) *Position {
	piece := u.moved

	// Undo promotion
	if u.wasPromotion {
		piece.Type = u.prevType
	}

	// Undo rook move for castling
	if u.wasCastle {
		b[u.rookFromX][u.rookY], b[u.rookToX][u.rookY] = u.rookPiece, nil
		if u.rookPiece != nil {
			u.rookPiece.HasMoved = u.prevRookMoved
		}
	}

	// Move piece back
	b[u.fromX][u.fromY], b[u.toX][u.toY] = piece, u.captured

	// Restore EP-captured pawn
	if u.wasEnPassant {
		b[u.epCaptX][u.epCaptY] = u.captured
		b[u.toX][u.toY] = nil
	}

	// Restore flags
	piece.HasMoved = u.prevMovedFlag

	return u.prevEnPassant
}
