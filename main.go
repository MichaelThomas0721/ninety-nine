package main

import (
	"fmt"
	"time"
)

// Game Data
var game Game

func main() {
	// lives := 3
	start := time.Now()
	games := 50
	loses := &Loses{}
	loses.player = make(map[int]int)
	for games > 0 {
		loses.player[runGame()] += 1
		games -= 1
	}
	fmt.Println(loses)
	fmt.Println("Finished in", time.Now().Sub(start))
}
