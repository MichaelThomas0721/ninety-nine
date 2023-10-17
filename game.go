package main

import (
	// "fmt"
	"math/rand"
	"time"
)

type Game struct {
	players []Player
	round   int
	cards   []Card
	dispose []Card
}

type Card struct {
	suit  string
	value int
}

func setupGame() {
	cards := shuffleCards(generateCards())
	players := generatePlayers(3)
	var dispose []Card
	game = Game{players, 0, cards, dispose}
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
		newPlayer := Player{[]Card{}, 3}
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
