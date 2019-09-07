package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value

	if t.Left != nil {
		go Walk(t.Left, ch)
	}

	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)

	var valuesT1 []int
	var valuesT2 []int

	go Walk(t1, ch)

	for i := 0; i < 10; i++ {
		valuesT1 = append(valuesT1, <-ch)
	}

	go Walk(t2, ch)

	for i := 0; i < 10; i++ {
		valuesT2 = append(valuesT2, <-ch)
	}

	sort.Ints(valuesT1)
	sort.Ints(valuesT2)

	var t1Arr [10]int
	var t2Arr [10]int

	copy(t1Arr[:], valuesT1)
	copy(t2Arr[:], valuesT2)

	return t1Arr == t2Arr
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
