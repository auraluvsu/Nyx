package main

import (
	"log"

	"auraluvsu.com/nyx/engine"
	nyx "auraluvsu.com/nyx/logic"
)

func main() {
	board := nyx.SetupBoard()
	log.Println(engine.Perft(board, 5, nyx.White))
	engine.PerftDivide(board, 5, nyx.White)
}
