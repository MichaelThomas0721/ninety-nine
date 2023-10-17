package main

type Strategy func([]Card, int) int

type DrawCard func(Player)

type Player struct {
	cards    []Card
	lives    int
	strategy Strategy
	status   bool
}

func createPlayer(lives int, strategy Strategy, id int) Player {
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

func drawCard(player *Player) {
	if len(game.cards) <= 0 {
		game.cards = game.dispose
		game.dispose = []Card{}
	}
	player.cards = append(player.cards, game.cards[len(game.cards)-1])
	game.cards = game.cards[:len(game.cards)-1]

}

func playCard(player *Player) bool {
	validCards := getValidCards()
	if len(validCards) == 0 {
		endRound()
		return true
	} else if len(validCards) == 1 {
		useCard(findCard(validCards[0], player.cards))
	} else {
		card := findCard(validCards[player.strategy(validCards, score)], player.cards)
		useCard(card)
	}
	drawCard(player)
	return false
}
