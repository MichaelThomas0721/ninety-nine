package main

var turn int
var score int
var direction bool

func startRound(starting int) int {
	dealCards()
	game.round += 1
	turn = starting
	score = 0
	direction = true
	end := false
	for !end {
		if game.players[turn].status {
			end = playCard(&game.players[turn])
		}

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
