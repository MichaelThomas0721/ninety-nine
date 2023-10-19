// Tested with 250 players and it worked, took long as hell tho so not recommended but it still works!!!

package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"strconv"
	"time"
)

type Strat struct {
	playerAmt int
	strat     StrategyList
	id        string
}

// Game Data
var game Game

const totalGames = 100000

// const totalGames = 1

func main() {
	hs := loadJson("strategies/highest.json")
	hrs := loadJson("strategies/highestRush.json")
	hsrs := loadJson("strategies/highestSemiRush.json")
	ms := loadJson("strategies/middle.json")
	hsars := loadJson("strategies/highestSafeRush.json")
	ls := loadJson("strategies/low.json")
	strat := []Strat{}

	strat = append(strat, Strat{2, hs, "highest"})
	strat = append(strat, Strat{2, hrs, "highestRush"})
	strat = append(strat, Strat{2, hsrs, "highestSemiRush"})
	strat = append(strat, Strat{2, ms, "middle"})
	strat = append(strat, Strat{2, hsars, "highestSafeRush"})
	strat = append(strat, Strat{2, ls, "lowest"})

	games := totalGames
	wins := &Wins{}
	wins.strategies = make(map[string]int)

	bar := progressbar.Default(totalGames)
	start := time.Now()

	for games > 0 {
		wins.strategies[runGame(strat)] += 1
		games -= 1
		bar.Add(1)
	}
	timeTaken := time.Since(start)
	timePerGame := timeTaken / totalGames
	fmt.Println(winString((wins)))
	fmt.Println("Finished in", timeTaken)
	fmt.Println("Per game", timePerGame)
}

func winString(wins *Wins) string {
	wString := ""
	for key, val := range wins.strategies {
		wString += key + " " + strconv.Itoa(int(float64(val)/float64(totalGames)*100)) + "% | "
	}
	return wString
}
