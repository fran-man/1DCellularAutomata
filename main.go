package main

import "fmt"

func main() {
	cells := 101
	row := make([]bool, cells)

	for i := 0; i < cells; i++ {
		row[i] = false
	}

	row[(cells-1)/2] = true
	fmt.Println(row)
}
