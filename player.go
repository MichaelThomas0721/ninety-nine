package main

type Strategy func([]Card, int) int

type Player struct {
	cards    []Card
	lives    int
	strategy Strategy
	status   bool
}

func createPlayer(lives int, strategy Strategy) Player {
	return Player{[]Card{}, lives, strategy, true}
}

func findCard(card Card, cards []Card) int {
	for c := range cards {
		if cards[c] == card {
			return c
		}
	}
	return 0
}
