package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	i := -1
	sequence := []int{0, 1}
	return func() int {
		len := len(sequence)
		next := sequence[len-1] + sequence[len-2]
		sequence = append(sequence, next)
		i++

		return sequence[i]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
