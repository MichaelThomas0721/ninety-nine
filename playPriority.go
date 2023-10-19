// This takes in the strategy information and returns the card to play

package main

func playPriority(cards []*Card, score int, pList StrategyList) int {
	// Check the ranges and the dependencies until finding one that matches the current score and
	// fill all the dependencies then find the top priority card
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

// Function for checking if the dependencies are good
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
