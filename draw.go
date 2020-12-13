package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

func drawBoard(board [][]bool) {
	fmt.Println("Beginning to draw board to png image")

	dc := gg.NewContext(5005, 2500)
	dc.SetRGB(0, 0, 0)

	for y, row := range board {
		fmt.Println(row)
		for x, cell := range row {
			if cell {
				fmt.Println("Drawing pixel", x, y)
				dc.DrawRectangle(float64(x*5), float64(y*5), float64(5), float64(5))
				dc.Fill()
			}
		}
	}

	dc.SavePNG("board.png")
}
