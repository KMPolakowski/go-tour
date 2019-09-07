package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x <= 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z float64 = 1
	for math.Abs(z*z-x) > 0.000000000001 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z, nil
}

func main() {
	result, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
