package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	var matrix = make([][]uint8, dx)

	for x := range matrix {
		matrix[x] = make([]uint8, dy)

		for y := range matrix[x] {
			matrix[x][y] = uint8(x ^ y)
		}
	}

	return matrix

}

func main() {
	pic.Show(Pic)
}
