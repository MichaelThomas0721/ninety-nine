// 10 Players is the theoretical limit for 1 deck, if you make the -10 card a skip instead, the theoretical limit is 12
// I will later have this auto adjust to add more decks if more than 11 players are in the game

package main

import (
	"fmt"
	"time"
)

type Strat struct {
	playerAmt int
	strat     Strategy
	id        int
}

// Game Data
var game Game
var stratJsons []StrategyList

const totalGames = 1000000

// const totalGames = 1

func main() {
	// lives := 3

	strat := []Strat{}
	// strat = append(strat, Strat{1, highestSafeRush_strategy, 0})
	strat = append(strat, Strat{10, highest_strategy, 1})
	// strat = append(strat, Strat{1, highestRush_strategy, 2})
	// strat = append(strat, Strat{1, highestSemiRush_strategy, 3})
	// strat = append(strat, Strat{1, middle_strategy, 4})

	stratJsons = []StrategyList{}

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
