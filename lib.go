// The cool thing about go is that I can declare these functions anywhere
// and just use them like they are in the same file so I just put a bunch
// of lib functions into this file :)

package main

import (
	"encoding/json"
	"io/ioutil"
)

// Since go doesn't have a built in contains function for arrays, I made my own
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Used to find the index of a card in an array of cards, I use this a lot
func findCard(card Card, cards []Card) int {
	for c := range cards {
		if cards[c] == card {
			return c
		}
	}
	return 0
}

// Used to convert a pointer card array to a normal one since
// I can't do everything with pointers
func pointerconv(cards []*Card) []Card {
	copy := []Card{}
	for card := range cards {
		copy = append(copy, *cards[card])
	}
	return copy
}

// Function for finding the highest priority card in the array
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

// Function for finding if an array has a card, I don't use this anymore
// but I am keeping it for now in case I decide to do functions again
// func hasCard(cards []*Card, value int) *Card {
// 	for c := range cards {
// 		if cards[c].value == value {
// 			return cards[c]
// 		}
// 	}
// 	return nil
// }

// Basically just used for loading the strategy jsons
func loadJson(location string) StrategyList {
	file, _ := ioutil.ReadFile(location)

	data := StrategyList{}

	_ = json.Unmarshal(file, &data)

	return data
}

// Find how many special cards are in an array
func getSpecials(cards []*Card) int {
	specials := 0
	for c := range cards {
		if contains([]int{3, 4, 9, 11}, cards[c].value) {
			specials += 1
		}
	}
	return specials
}
