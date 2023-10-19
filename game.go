package main

import (
	"math"
	"math/rand"
	"time"
)

type StartFunc func() string

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

func runGame(strats []Strat) string {
	setupGame(strats)
	return game.start()

}

func setupGame(strats []Strat) {
	players := generatePlayers(strats)
	cards := shuffleCards(generateCards(len(players)))

	var dispose []Card
	game = Game{players, 0, cards, dispose, startGame}
}

func startGame() string {
	starting := 0
	for true {
		loser := startRound(starting)
		game.players[loser].lives -= 1
		if game.players[loser].lives <= 0 {
			game.players[loser].status = false
			if !checkRemainingPlayers(game.players) {
				for p := range game.players {
					if game.players[p].status == true {

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
	return "NILL"
}

func generateCards(players int) []Card {
	values := [2]int{1, 13}
	suits := [4]string{"hearts", "diamonds", "clubs", "spades"}
	var cards []Card

	for suit := range suits {
		for i := values[0]; i <= values[1]; i++ {
			cards = append(cards, Card{suits[suit], i})
		}
	}

	decks := math.Ceil(float64(players) / 8)
	decksInt := int(decks)
	tempCards := cards
	for i := 0; i < decksInt-1; i++ {
		cards = append(cards, tempCards...)
	}

	return cards
}

func shuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

func shufflePlayers(players []Player) []Player {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(players), func(i, j int) {players[i], players[j] = players[j], players[i]})
	return players
}

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
