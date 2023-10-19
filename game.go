package main

import (
	"math"
	"math/rand"
	"time"
)

// This was from when I was using functions instead of jsons, I am keeping it like
// this for now in case I want to switch back, I am not sure if I prefer the jsons just yet
type StartFunc func() string

// For holding the information about the current game
type Game struct {
	players []Player
	round   int
	cards   []Card
	dispose []Card
	start   StartFunc
}

type Card struct {
	suit  string
	value int
}

// Setup and start
func runGame(strats []Strat) string {
	setupGame(strats)
	return game.start()

}

// Make the players and cards then add them to a game
func setupGame(strats []Strat) {
	players := generatePlayers(strats)
	cards := shuffleCards(generateCards(len(players)))

	var dispose []Card
	game = Game{players, 0, cards, dispose, startGame}
}

// Run each round until there isn't anyone left, I might change it from "for" to "for checkRemainingPlayers"
// Undecided
func startGame() string {
	// For rotating the starting player at the beginning of each round
	starting := 0
	for {
		loser := startRound(starting)
		game.players[loser].lives -= 1
		if game.players[loser].lives <= 0 {
			game.players[loser].status = false
			if !checkRemainingPlayers(game.players) {
				for p := range game.players {
					if game.players[p].status {
						// Return the winning player's strategy, used for stat tracking
						return game.players[p].id
					}
				}
			}
		}
		starting += 1
		if starting >= len(game.players) {
			starting = 0
		}
	}
	// This should never happen but I added it just in case
	// I mean the only way out of that for loop is the return statement so either it goes forever
	// or an error happens but just to be safe
	return "NILL"
}

// Make cards 1-13 (Ace - 10 + face cards) in all four suits
func generateCards(players int) []Card {
	values := [2]int{1, 13}
	suits := [4]string{"hearts", "diamonds", "clubs", "spades"}
	var cards []Card

	for suit := range suits {
		for i := values[0]; i <= values[1]; i++ {
			cards = append(cards, Card{suits[suit], i})
		}
	}

	// I found that adding 10 or more players would either break the game or cause it to go forever since
	// the players were just putting down special cards infinitely so I have made it automatically add
	// more decks if there are 8 or more players
	// I chose 8 because I felt that was a more realistic number than 10
	decks := math.Ceil(float64(players) / 8)
	decksInt := int(decks)
	tempCards := cards
	for i := 0; i < decksInt-1; i++ {
		cards = append(cards, tempCards...)
	}

	return cards
}

// Function for shuffling the cards so they aren't in order
func shuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

// Function for shuffling the players, I found that having each strategy next to their own was unrealistic
// and scewed results, by shuffling them, more risky strategies lose less since they aren't fighting with
// each other essentially
func shufflePlayers(players []Player) []Player {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
	return players
}

// Function for making the players and tieing them to their strategy (id)
func generatePlayers(strats []Strat) []Player {
	var players []Player

	for i := 0; i < len(strats); i++ {
		for j := 0; j < strats[i].playerAmt; j++ {
			newPlayer := createPlayer(3, strats[i].strat, strats[i].id)
			players = append(players, newPlayer)
		}
	}
	players = shufflePlayers(players)

	return players
}

// Function for checking if there are two or more players in the game,
// if there isn't than it would just be someone playing against themselves
// and we can't have that
func checkRemainingPlayers(players []Player) bool {
	in := 0
	for p := range players {

		if players[p].status {
			in += 1
			if in > 1 {
				return true
			}
		}
	}
	return false
}
