package main

import (
	// "log"

	"auraluvsu.com/nyx/engine"
	nyx "auraluvsu.com/nyx/logic"
)

func main() {
	board := nyx.SetupBoard()
	engine.PerftDivide(board, 6, nyx.White, nil)
}
