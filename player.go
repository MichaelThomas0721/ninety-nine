package main

type Strategy func([]*Card, int) int

type DrawCard func(Player)

type Player struct {
	cards    []Card
	lives    int
	strategy StrategyList
	status   bool
	id string
}

func createPlayer(lives int, strategy StrategyList, id string) Player {
	return Player{[]Card{}, lives, strategy, true, id}
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
		useCard(findCard(*validCards[0], player.cards), player)
	} else {
		card := findCard(*validCards[playPriority(validCards, score, player.strategy)], player.cards)
		useCard(card, player)
	}
	drawCard(player)
	return false
}

func getValidCards() []*Card {
	var validCards []*Card
	for card := range game.players[turn].cards {
		val := game.players[turn].cards[card].value
		if val > 11 {
			val = 10
		}
		if contains([]int{3, 4, 9, 11}, val) || val <= 99-score {
			validCards = append(validCards, &game.players[turn].cards[card])
		}
	}
	return validCards
}

func discardCard(card int, player *Player) {
	game.dispose = append(game.dispose, player.cards[card])
	player.cards = append(player.cards[:card], player.cards[card+1:]...)
}

func useCard(card int, player *Player) {
	val := player.cards[card].value
	if val > 11 {
		val = 10
	}
	if contains([]int{3, 4, 9, 11}, val) {
		if val == 9 {
			score = 99
		} else if val == 4 {
			direction = !direction
		} else if val == 11 {
			score -= 10
		}
	} else {
		score += val
	}
	// fmt.Println("CARD", player.cards[card].value, "SCORE", score)
	discardCard(card, player)
}
