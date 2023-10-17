package main

import (
	"fmt"
	"time"
)

// Game Data
var game Game

type Loses struct {
	player map[int]int
}

func main() {
	// lives := 3
	start := time.Now()
	games := 50000
	loses := &Loses{}
	loses.player = make(map[int]int)
	setupGame()
	for games > 0 {
		loser := startRound()
		game.players[loser].lives -= 1
		if game.players[loser].lives <= 0 {
			games -= 1
			loses.player[loser] += 1
			setupGame()
		}
	}
	fmt.Println(loses)
	fmt.Println("Finished in", time.Now().Sub(start))
}
