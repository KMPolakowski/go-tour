package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for math.Abs(z*z-x) > 0.000000000001 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
