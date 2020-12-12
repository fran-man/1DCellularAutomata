package main

import "fmt"

func main() {
	cellsN := 101
	row := make([]bool, cellsN)

	for i := 0; i < cellsN; i++ {
		row[i] = false
	}

	row[(cellsN-1)/2] = true
	fmt.Println(row)

	ruleset := ruleset()

	nextRow := make([]bool, cellsN)
	for i := 1; i < cellsN-1; i++ {
		ruleToApply := stateToInt(row[i+1]) + stateToInt(row[i])*2 + stateToInt(row[i-1])*4
		nextRow[i] = ruleset[ruleToApply]
	}
	nextRow[0] = false
	nextRow[cellsN-1] = false
	fmt.Println(nextRow)
}

func ruleset() []bool {
	ruleset := make([]bool, 8)
	for i := 0; i < 8; i++ {
		ruleset[i] = false
	}
	ruleset[1] = true
	return ruleset
}

func stateToInt(state bool) int {
	if state {
		return 1
	} else {
		return 0
	}
}
