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
	for true {
		loser := startRound()
		game.players[loser].lives -= 1
		if game.players[loser].lives <= 0 {
			return loser
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
		newPlayer := createPlayer(3, highest_strategy)
		players = append(players, newPlayer)
	}

	return players
}

func drawCard(turn int) {
	if len(game.cards) <= 0 {
		game.cards = game.dispose
		game.dispose = []Card{}
	}
	game.players[turn].cards = append(game.players[turn].cards, game.cards[len(game.cards)-1])
	game.cards = game.cards[:len(game.cards)-1]

}
