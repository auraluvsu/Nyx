package parsing

import (
	"fmt"
	"os"
	"regexp"

	nyx "auraluvsu.com/nyx/logic"
)

func ParseSAN(move string, colour nyx.Colour) (*nyx.Move, error) {
	m := &nyx.Move{}
	if move == "exit" {
		os.Exit(1)
	}
	if move == "O-O" || move == "0-0" {
		m.IsCastle = true
		m.Piece = nyx.King
		m.Ty = 0
		if colour == nyx.White {
			m.Ty = 7
		}
		m.Tx = 6
		return m, nil
	}
	if move == "O-O-O" || move == "0-0-0" {
		m.IsCastle = true
		m.Piece = nyx.King
		m.Ty = 0
		if colour == nyx.White {
			m.Ty = 7
		}
		m.Tx = 2
		return m, nil
	}

	// Regex groups: 1:Piece, 2:From, 3:Capture, 4:To, 5:Promo
	re := regexp.MustCompile(`^([NBRQK])?([a-h]?[1-8]?)?(x)?([a-h][1-8])(=[NBRQ])?$`)
	matches := re.FindStringSubmatch(move)

	if matches == nil {
		return nil, fmt.Errorf("invalid move format: %s", move)
	}

	// Piece
	pieceChar := matches[1]
	if pieceChar == "" {
		m.Piece = nyx.Pawn
	} else {
		pieceMap := map[byte]nyx.PieceType{
			'N': nyx.Knight, 'B': nyx.Bishop, 'R': nyx.Rook, 'Q': nyx.Queen, 'K': nyx.King,
		}
		m.Piece = pieceMap[pieceChar[0]]
	}

	// To Square
	toSquare := matches[4]
	x, y, err := ParsePosition(toSquare)
	if err != nil {
		return nil, err
	}
	m.Tx, m.Ty = x, y

	// Capture
	if matches[3] == "x" {
		m.IsCapture = true
	}

	// From Square (Disambiguation)
	from := matches[2]
	if m.Piece == nyx.Pawn && m.IsCapture {
		if len(from) == 1 && from[0] >= 'a' && from[0] <= 'h' {
			file := int(from[0] - 'a')
			m.Fx = &file
		} else {
			return nil, fmt.Errorf("invalid pawn capture format: %s", move)
		}
	} else if from != "" {
		if len(from) == 1 {
			if from[0] >= 'a' && from[0] <= 'h' {
				file := int(from[0] - 'a')
				m.Fx = &file
			} else { // is a rank
				rank := 8 - int(from[0]-'0')
				m.Fy = &rank
			}
		} else if len(from) == 2 {
			fx, fy, err := ParsePosition(from)
			if err != nil {
				return nil, fmt.Errorf("invalid from square in move: %s", move)
			}
			m.Fx = &fx
			m.Fy = &fy
		}
	}

	// Promotion
	if matches[5] != "" {
		promoChar := matches[5][1] // skip '='
		promoMap := map[byte]nyx.PieceType{
			'N': nyx.Knight, 'B': nyx.Bishop, 'R': nyx.Rook, 'Q': nyx.Queen,
		}
		promoPiece, ok := promoMap[promoChar]
		if !ok {
			return nil, fmt.Errorf("invalid promotion piece: %c", promoChar)
		}
		m.PromoteTo = &promoPiece
	}

	return m, nil
}

