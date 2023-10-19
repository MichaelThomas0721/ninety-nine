package main

type Wins struct {
	strategies map[string]int
}

type StrategyList struct {
	Ranges []PriorityRange `json:"ranges"`
}

type Dependency struct {
	Type string `json:"type"`
	Amt  int    `json:"amt"`
}

type PriorityRange struct {
	Start        int          `json:"start"`
	End          int          `json:"end"`
	Dependencies []Dependency `json:"dependencies"`
	Priority     []int        `json:"priority"`
}
