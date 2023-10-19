// Tested with 250 players and it worked, took long as hell tho so not recommended but it still works!!!

package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"
)

// For holding all the strategy information
type Strat struct {
	playerAmt int
	strat     StrategyList
	id        string
}

// Gloabl Game Data and Constants
var game Game

const totalGames = 100000

// In case I want to just simulate 1 game
// const totalGames = 1

func main() {
	// Loading all the strategies from their json file
	hs := loadJson("strategies/highest.json")
	hrs := loadJson("strategies/highestRush.json")
	hsrs := loadJson("strategies/highestSemiRush.json")
	ms := loadJson("strategies/middle.json")
	hsars := loadJson("strategies/highestSafeRush.json")
	ls := loadJson("strategies/low.json")

	// Put the strategies into an array for easy player creation
	strat := []Strat{}
	strat = append(strat, Strat{2, hs, "highest"})
	strat = append(strat, Strat{2, hrs, "highestRush"})
	strat = append(strat, Strat{2, hsrs, "highestSemiRush"})
	strat = append(strat, Strat{2, ms, "middle"})
	strat = append(strat, Strat{2, hsars, "highestSafeRush"})
	strat = append(strat, Strat{2, ls, "lowest"})

	// For tracking the games progress and wins
	games := totalGames
	wins := &Wins{}
	wins.strategies = make(map[string]int)

	// For stat tracking, tells me current progress and how long the test took
	bar := progressbar.Default(totalGames)
	start := time.Now()

	// Simulate all the games, track the wins and update the progress
	for games > 0 {
		wins.strategies[runGame(strat)] += 1
		games -= 1
		bar.Add(1)
	}

	// For post test stats, tells me the win percentage of each strategy and how long it took to run the test, it also get the average test time
	// which is useful for when I am trying to make it more efficient
	timeTaken := time.Since(start)
	timePerGame := timeTaken / totalGames
	fmt.Println(winString((wins)))
	fmt.Println("Finished in", timeTaken)
	fmt.Println("Per game", timePerGame)
}

// Function that just takes the wins and turns it into a percentage
func winString(wins *Wins) string {
	wString := ""
	for key, val := range wins.strategies {
		wString += key + " " + strconv.Itoa(int(float64(val)/float64(totalGames)*100)) + "% | "
	}
	return wString
}
