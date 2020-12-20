package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func drawBoard(board [][]bool, filename string) {
	fmt.Println("Beginning to draw board to png image")

	dc := gg.NewContext(2002, 1000)
	dc.SetRGB(0, 0, 0)

	for y, row := range board {
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
