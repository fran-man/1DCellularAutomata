package main

func main() {
	cellsN := 1001
	rowsN := 500
	row := make([]bool, cellsN)
	board := make([][]bool, rowsN)

	for i := 0; i < cellsN; i++ {
		row[i] = false
	}

	row[(cellsN-1)/2] = true
	board[0] = row

	ruleset := ruleset()

	for r := 1; r < rowsN; r++ {
		row = generateNextRow(row, ruleset)
		board[r] = row
	}

	drawBoard(board)
}

func generateNextRow(row []bool, ruleset []bool) []bool {
	nextRow := make([]bool, len(row))
	for i := 1; i < len(row)-1; i++ {
		ruleToApply := stateToInt(row[i+1]) + stateToInt(row[i])*2 + stateToInt(row[i-1])*4
		nextRow[i] = ruleset[ruleToApply]
	}
	nextRow[0] = false
	nextRow[len(row)-1] = false

	return nextRow
}

func ruleset() []bool {
	ruleset := make([]bool, 8)
	for i := 0; i < 8; i++ {
		ruleset[i] = false
	}
	ruleset[1] = true
	ruleset[2] = true
	return ruleset
}

func stateToInt(state bool) int {
	if state {
		return 1
	}
	return 0
}
