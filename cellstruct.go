package main

type Cell struct {
	state      bool
	trueStreak int
}

func blankCell() *Cell {
	c := Cell{state: false, trueStreak: 0}
	return &c
}
