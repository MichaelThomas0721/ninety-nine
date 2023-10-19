package main

func playPriority(cards []*Card, score int, pList StrategyList) int {

	var priority []int
	for r := range pList.Ranges {
		if pList.Ranges[r].Start <= score && pList.Ranges[r].End >= score {
			dep := true
			if pList.Ranges[r].Dependencies != nil {
				for d := range pList.Ranges[r].Dependencies {
					if !checkDependecies(cards, pList.Ranges[r].Dependencies[d]) {
						dep = false
						break
					}
				}
			}
			if dep {
				priority = pList.Ranges[r].Priority
				break
			}
		}
	}
	pcCards := pointerconv(cards)
	fp := findPriority(cards, priority)
	if fp != nil {
		return findCard(*fp, pcCards)
	}
	return 0
}

func checkDependecies(cards []*Card, dependency Dependency) bool {
	if dependency.Type == "specials" {
		if getSpecials(cards) >= dependency.Amt {
			return true
		} else {
			return false
		}
	}
	return false
}
