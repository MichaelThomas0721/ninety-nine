// 10 Players is the theoretical limit for 1 deck, if you make the -10 card a skip instead, the theoretical limit is 12
// I will later have this auto adjust to add more decks if more than 11 players are in the game

package main

import (
	"fmt"
	"time"
)

type Strat struct {
	playerAmt int
	strat     StrategyList
	id        int
}

// Game Data
var game Game

const totalGames = 10000

// const totalGames = 1

func main() {
	// lives := 3
	hs := loadJson("strategies/highest.json")
	hrs := loadJson("strategies/highestRush.json")
	hsrs := loadJson("strategies/highestSemiRush.json")
	ms := loadJson("strategies/middle.json")
	strat := []Strat{}

	strat = append(strat, Strat{2, hs, 1})
	strat = append(strat, Strat{2, hrs, 2})
	strat = append(strat, Strat{2, hsrs, 3})
	strat = append(strat, Strat{2, ms, 4})

	games := totalGames
	wins := &Wins{}
	wins.player = make(map[int]int)

	start := time.Now()

	for games > 0 {
		wins.player[runGame(strat)] += 1
		games -= 1
		fmt.Println("Game finished: ", totalGames-games)
	}
	timeTaken := time.Now().Sub(start)
	timePerGame := timeTaken / totalGames
	fmt.Println(wins)
	fmt.Println("Finished in", timeTaken)
	fmt.Println("Per game", timePerGame)
}
