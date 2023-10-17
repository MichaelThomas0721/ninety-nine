package main

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findCard(card Card, cards []Card) int {
	for c := range cards {
		if cards[c] == card {
			return c
		}
	}
	return 0
}
