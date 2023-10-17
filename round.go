package main

var turn int
var score int
var direction bool

func startRound() int {
	dealCards()
	game.round += 1
	turn = 0
	score = 0
	direction = true
	end := false
	for !end {
		end = playCard(&game.players[turn])
		if end {
			return turn
		}
		if direction {
			turn += 1
		} else {
			turn -= 1
		}
		if turn >= 3 {
			turn = 0
		} else if turn < 0 {
			turn = 2
		}
	}
	return -1
}

func endRound() {
	for player := range game.players {
		game.dispose = append(game.dispose, game.players[player].cards...)
		game.players[player].cards = []Card{}
	}

}

func dealCards() {
	for i := 0; i < 3; i++ {
		for player := range game.players {
			drawCard(&game.players[player])
		}
	}
}

func discardCard(card int) {
	game.dispose = append(game.dispose, game.players[turn].cards[card])
	game.players[turn].cards = append(game.players[turn].cards[:card], game.players[turn].cards[card+1:]...)
}

func getValidCards() []Card {
	var validCards []Card
	for card := range game.players[turn].cards {
		val := game.players[turn].cards[card].value
		if val > 10 {
			val = 10
		}
		if contains([]int{3, 4, 9}, val) || val <= 99-score {
			validCards = append(validCards, game.players[turn].cards[card])
		}
	}
	return validCards
}

func useCard(card int) {
	val := game.players[turn].cards[card].value
	if contains([]int{3, 4, 9}, val) {
		if val == 9 {
			score = 99
		} else if val == 4 {
			direction = !direction
		}
	} else {
		score += val
	}
	discardCard(card)
}
