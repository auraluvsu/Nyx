package parsing

import (
	"fmt"
)

func ParsePosition(s string) (int, int, error) {
	if len(s) != 2 {
		return -1, -1, fmt.Errorf("Invalid move! Not a real move")
	}
	file := int(s[0] - 'a')
	rank := 8 - int(s[1]-'0')

	if file < 0 || file > 7 || rank < 0 || rank > 7 {
		return -1, -1, fmt.Errorf("Invalid move! Out of bounds")
	}
	return file, rank, nil
}
