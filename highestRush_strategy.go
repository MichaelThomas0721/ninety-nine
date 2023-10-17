package main

import (
	"sort"
)

func highestRush_strategy(cards []*Card, score int) int {
	pcCards := pointerconv(cards)
	if hasCard(cards, 9) != nil {
		return findCard(*hasCard(cards, 9), pcCards)
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].value > cards[j].value
	})

	var card *Card
	for c := range cards {
		if !contains([]int{3, 4, 11}, cards[c].value) {
			card = cards[c]
			return findCard(*card, pcCards)
		}
	}
	if score == 99 {
		priority := []int{3, 4, 11}
		fp := findPriority(cards, priority)
		if fp != nil {
			return findCard(*fp, pcCards)
		}
	}
	return 0
}
