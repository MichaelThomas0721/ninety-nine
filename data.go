package main

type Wins struct {
	player map[int]int
}

type StrategyList struct {
	Ranges []PriorityRange `json:"ranges"`
}

type PriorityRange struct {
	Start    int   `json:"start"`
	End      int   `json:"end"`
	Priority []int `json:"priority"`
}
