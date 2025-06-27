package parsing

import (
	"fmt"
	"log"
	"os"
	"regexp"

	nyx "auraluvsu.com/nyx/logic"
)

func ParseSAN(move string) (*nyx.Move, error) {
	m := &nyx.Move{}
	if move == "exit" {
		os.Exit(1)
	}
	re := regexp.MustCompile(`^([NBRQK]?)([a-h]?)([1-8]?)[x-]?([a-h][1-8])$`)
	matches := re.FindStringSubmatch(move)
	if matches == nil {
		return nil, fmt.Errorf("Could not parse move: %s", move)
	}
	//if move == "O-O" || move == "O-O-O" {
	// Logic for later...
	//}
	pieceChar := matches[1]
	if pieceChar == "exit" {
		os.Exit(1)
	}
	if pieceChar == "" {
		m.Piece = nyx.Pawn
	} else {
		pieceMap, ok := map[byte]nyx.PieceType{
			'N': nyx.Knight,
			'B': nyx.Bishop,
			'R': nyx.Rook,
			'Q': nyx.Queen,
			'K': nyx.King,
		}[pieceChar[0]]
		if !ok {
			return nil, fmt.Errorf("Invalid piece: %s", pieceChar)
		}
		m.Piece = pieceMap
	}
	if matches[2] != "" {
		file := int(matches[2][0] - 'a')
		m.Fx = &file
	}
	if matches[3] != "" {
		rank := int(matches[3][0] - '1')
		m.Fy = &rank
	}
	x, y, err := ParsePosition(matches[4])
	if err != nil {
		log.Fatal(err)
	}
	if x == -1 || y == -1 {
		return nil, fmt.Errorf("Invalid square: %s", matches[4])
	}
	m.Tx, m.Ty = x, y
	return m, nil
}
