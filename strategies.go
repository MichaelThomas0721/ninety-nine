package main

import (
	"sort"
)

func highest_strategy(cards []*Card, score int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].value > cards[j].value
	})
	var card *Card
	for c := range cards {
		if !contains([]int{3,4,9}, cards[c].value) {
			card = cards[c]
			return findCard(*card, pointerconv(cards))
		}
	}
	return 0

	// var regulars []Card
	// for card := range cards {
	// 	if !contains([]int{3, 4, 9}, cards[card].value) {
	// 		regulars = append(regulars, *cards[card])
	// 	}
	// }
	// if len(regulars) >= 1 {
	// 	sort.Slice(regulars, func(i, j int) bool {
	// 		return regulars[i].value < regulars[j].value
	// 	})
	// 	return findCard(regulars[0], pointerconv(cards))
	// }
	// return 0
}

func pointerconv(cards []*Card) []Card {
	copy := []Card{}
	for card := range cards {
		copy = append(copy, *cards[card])
	}
	return copy
}

// func findMax(values []int, exclusions []int) int {
// 	min := values[0]
// 	for _, v := range values {
// 		if v < min {
// 			min = v
// 		}
// 	}

// }
