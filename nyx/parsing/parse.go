package parsing

import (
	"fmt"

	nyx "auraluvsu.com/nyx/engine"
)

func ParsePosition(s string) (int, int, error) {
	if len(s) != 2 {
		return -1, -1, fmt.Errorf("Invalid move! Not a real move")
	}
	file := s[0]
	rank := s[1]

	x := int(file - 'a')
	y := int(rank - '0')
	if !nyx.InBounds(x, y) {
		return -1, -1, fmt.Errorf("Invalid move! Out of bounds")
	}
	return x, y, nil
}
