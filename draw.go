package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func drawBoard(board [][]*Cell, filename string) {
	fmt.Println("Beginning to draw board to png image")

	dc := gg.NewContext(2002, 1000)

	dc.SetRGB(1, 1, 1)

	dc.DrawRectangle(0, 0, 2002, 1000)
	dc.Fill()

	dc.SetRGB(0, 0, 0)

	for y, row := range board {
		dc.SetRGB(float64(y)/float64(len(board)), 0, 0)
		for x, cell := range row {
			if cell.state {
				dc.SetRGB(float64(cappedStreak(cell.trueStreak))/15, 0, 0)
				dc.DrawRectangle(float64(x*8), float64(y*8), float64(8), float64(8))
				dc.Fill()
			}
		}
	}

	dc.SavePNG("output/" + filename + ".png")
}

func cappedStreak(streak int) int {
	if streak > 15 {
		return 15
	}
	return streak
}
