// WHY DOES IT BREAK WHEN THERE ARE MORE THAN 3 PEOPLE!!!!!!!!??????? WTF

package main

import (
	"fmt"
	"time"
)

type Strat struct {
	playerAmt int
	strat     Strategy
}

// Game Data
var game Game

const totalGames = 1000

// const totalGames = 1

func main() {
	// lives := 3

	strat := []Strat{}
	strat = append(strat, Strat{1, highest_strategy})
	strat = append(strat, Strat{2, highestRush_strategy})

	games := totalGames
	wins := &Wins{}
	wins.player = make(map[int]int)

	start := time.Now()

	for games > 0 {
		wins.player[runGame(strat)] += 1
		games -= 1
	}
	timeTaken := time.Now().Sub(start)
	timePerGame := timeTaken / totalGames
	fmt.Println(wins)
	fmt.Println("Finished in", timeTaken)
	fmt.Println("Per game", timePerGame)
}
