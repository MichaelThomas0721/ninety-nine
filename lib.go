package main

import (
	"encoding/json"
	"io/ioutil"
)

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

func pointerconv(cards []*Card) []Card {
	copy := []Card{}
	for card := range cards {
		copy = append(copy, *cards[card])
	}
	return copy
}

func findPriority(cards []*Card, values []int) *Card {
	for i := range values {
		for c := range cards {
			if cards[c].value == values[i] {
				return cards[c]
			}
		}
	}
	return nil
}

func hasCard(cards []*Card, value int) *Card {
	for c := range cards {
		if cards[c].value == value {
			return cards[c]
		}
	}
	return nil
}

func loadJson(location string) StrategyList {
	file, _ := ioutil.ReadFile(location)

	data := StrategyList{}

	_ = json.Unmarshal(file, &data)

	return data
}
