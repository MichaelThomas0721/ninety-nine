package main

import (
	"fmt"
	"time"
)

// Game Data
var game Game

const totalGames = 100000

// const totalGames = 1

func main() {
	// lives := 3
	start := time.Now()

	games := totalGames
	wins := &Wins{}
	wins.player = make(map[int]int)
	for games > 0 {
		wins.player[runGame()] += 1
		games -= 1
	}
	timeTaken := time.Now().Sub(start)
	timePerGame := timeTaken / totalGames
	fmt.Println(wins)
	fmt.Println("Finished in", timeTaken)
	fmt.Println("Per game", timePerGame)
}
