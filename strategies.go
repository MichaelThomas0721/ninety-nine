package main

import (
	"sort"
)

func highest_strategy(cards []Card, score int) int {
	var regulars []Card
	for card := range cards {
		if !contains([]int{3, 4, 9}, cards[card].value) {
			regulars = append(regulars, cards[card])
		}
	}
	if len(regulars) >= 1 {
		sort.Slice(regulars, func(i, j int) bool {
			return regulars[i].value < regulars[j].value
		})
		return findCard(regulars[0], cards)
	}
	return 0
}
