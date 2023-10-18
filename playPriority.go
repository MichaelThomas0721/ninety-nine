package main

func playPriority(cards []*Card, score int, pList StrategyList) int {

	var priority []int
	for r := range pList.Ranges {
		if pList.Ranges[r].Start <= score && pList.Ranges[r].End >= score {
			priority = pList.Ranges[r].Priority
			break

		}
	}
	pcCards := pointerconv(cards)
	fp := findPriority(cards, priority)
	if fp != nil {
		return findCard(*fp, pcCards)
	}
	return 0
}
