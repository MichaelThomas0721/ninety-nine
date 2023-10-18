package main

import (
	"sort"
)

func middle_strategy(cards []*Card, score int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].value > cards[j].value
	})
	pcCards := pointerconv(cards)
	var card *Card
	priority := []int{8, 7, 6, 5, 13, 12, 10, 2, 1}
	fp := findPriority(cards, priority)
	if fp != nil {
		return findCard(*fp, pcCards)
	}
	for c := range cards {
		if !contains([]int{3, 4, 9, 11}, cards[c].value) {
			card = cards[c]
			return findCard(*card, pcCards)
		}
	}
	if score == 99 {
		priority := []int{3, 4, 9, 11}
		fp := findPriority(cards, priority)
		if fp != nil {
			return findCard(*fp, pcCards)
		}
	}
	return 0
}
