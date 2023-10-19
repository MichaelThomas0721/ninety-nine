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
	hs := loadJson("strategies/highest.json")
	hrs := loadJson("strategies/highestRush.json")
	hsrs := loadJson("strategies/highestSemiRush.json")
	ms := loadJson("strategies/middle.json")
	hsars := loadJson("strategies/highestSafeRush.json")
	strat := []Strat{}

	strat = append(strat, Strat{5, hs, 1})
	strat = append(strat, Strat{5, hrs, 2})
	strat = append(strat, Strat{5, hsrs, 3})
	strat = append(strat, Strat{5, ms, 4})
	strat = append(strat, Strat{5, hsars, 5})

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
