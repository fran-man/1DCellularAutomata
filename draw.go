package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func drawBoard(board [][]bool, filename string) {
	fmt.Println("Beginning to draw board to png image")

	dc := gg.NewContext(2002, 1000)

	dc.SetRGB(1, 1, 1)

	dc.DrawRectangle(0, 0, 2002, 1000)
	dc.Fill()

	dc.SetRGB(0, 0, 0)

	for y, row := range board {
		dc.SetRGB(float64(y)/float64(len(board)), 0, 0)
		for x, cell := range row {
			if cell {
				// fmt.Println("Drawing pixel", x, y)
				dc.DrawRectangle(float64(x*4), float64(y*4), float64(4), float64(4))
				dc.Fill()
			}
		}
	}

	dc.SavePNG("output/" + filename + ".png")
}
