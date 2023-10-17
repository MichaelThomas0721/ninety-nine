package main

import (
	"fmt"
	"time"
)

// Game Data
var game Game

const totalGames = 5000

func main() {
	// lives := 3
	start := time.Now()

	games := totalGames
	loses := &Loses{}
	loses.player = make(map[int]int)
	for games > 0 {
		loses.player[runGame()] += 1
		games -= 1
	}
	timeTaken := time.Now().Sub(start)
	timePerGame := timeTaken / totalGames
	fmt.Println(loses)
	fmt.Println("Finished in", timeTaken)
	fmt.Println("Per game", timePerGame)
}
