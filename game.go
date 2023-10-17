package main

import (
	"math/rand"
	"time"
)

type StartFunc func() int

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

type Round struct {
	score     int
	direction bool
	turn      int
}

func runGame() int {
	setupGame()
	return game.start()

}

func setupGame() {
	cards := shuffleCards(generateCards())
	players := generatePlayers(3)
	var dispose []Card
	game = Game{players, 0, cards, dispose, startGame}
}

func startGame() int {
	starting := 0
	for true {
		loser := startRound(starting)
		game.players[loser].lives -= 1
		if game.players[loser].lives <= 0 {
			game.players[loser].status = false
			if !checkRemainingPlayers(game.players) {
				for p := range game.players {
					if game.players[p].status == true {
						return p
					}
				}
			}
		}
		starting += 1
		if starting >= 3 {
			starting = 0
		}
	}
	return 0
}

func generateCards() []Card {
	values := [2]int{1, 13}
	suits := [4]string{"hearts", "diamonds", "clubs", "spades"}
	var cards []Card

	for suit := range suits {
		for i := values[0]; i <= values[1]; i++ {
			cards = append(cards, Card{suits[suit], i})
		}
	}

	return cards
}

func shuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

func generatePlayers(amt int) []Player {
	var players []Player

	for i := 0; i < amt; i++ {
		newPlayer := createPlayer(3, highest_strategy, i)
		players = append(players, newPlayer)
	}

	return players
}

func checkRemainingPlayers(players []Player) bool {
	in := 0
	for p := range players {
		if players[p].status == true {
			in += 1
			if in > 1 {
				return true
			}
		}
	}
	return false
}
