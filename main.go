package main

import "strconv"

func main() {
	for idx, rule := range rulesToPrint {
		drawAndPrintBoardForRule(idx, rule, true)
	}
	//drawAndPrintBoardForRule(110, true)
}

func drawAndPrintBoardForRule(ruleIdx int, rule int, drawAllFrames bool) {
	row := make([]*Cell, cellsN)
	board := make([][]*Cell, rowsN)

	for i := 0; i < cellsN; i++ {
		row[i] = blankCell()
	}

	row[(cellsN-1)/2].state = true
	row[(cellsN-1)/2].trueStreak = 1
	board[0] = row

	ruleset := rulesFromInt(rule)

	for r := 1; r < rowsN; r++ {
		row = generateNextRow(row, ruleset)
		board[r] = row

		if drawAllFrames {
			drawBoard(board, "InterestingRules-frame"+strconv.Itoa(calculateOverallFrame(ruleIdx, r)))
		}
	}
}

func generateNextRow(row []*Cell, ruleset []bool) []*Cell {
	nextRow := make([]*Cell, len(row))
	for i := 0; i < len(nextRow); i++ {
		nextRow[i] = blankCell()
	}

	for i := 1; i < len(row)-1; i++ {
		ruleToApply := stateToInt(row[i+1].state) + stateToInt(row[i].state)*2 + stateToInt(row[i-1].state)*4
		nextRow[i].state = ruleset[ruleToApply]
		if nextRow[i].state {
			nextRow[i].trueStreak = row[i].trueStreak + 1
		}
	}

	nextRow[0].state = false
	nextRow[len(row)-1].state = false

	return nextRow
}

func stateToInt(state bool) int {
	if state {
		return 1
	}
	return 0
}

func calculateOverallFrame(ruleIdx int, relativeFrame int) int {
	result := ruleIdx*(rowsN-1) + relativeFrame
	return result
}
