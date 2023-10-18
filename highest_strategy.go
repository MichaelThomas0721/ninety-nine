package main

import (
	"encoding/json"
	"io/ioutil"
)

func highest_strategy(cards []*Card, score int) int {
	if len(stratJsons) == 0 {

		file, _ := ioutil.ReadFile("highestPriorities.json")

		data := StrategyList{}

		_ = json.Unmarshal(file, &data)
		stratJsons = append(stratJsons, data)
	}
	play := playPriority(cards, score, stratJsons[0])
	return play
}
