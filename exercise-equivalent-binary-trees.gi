package main

import (
	"fmt"
	"strconv"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walker(t, ch)
	close(ch)
}

func Walker(t *tree.Tree, ch chan int) {
	if t != nil {
		Walker(t.Left, ch)
		ch <- t.Value
		Walker(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool {
	// Implement as logic in main()
	return true
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go Walk(tree.New(1), ch)
	go Walk(tree.New(1), ch2)

	var st1 string
	var st2 string
	same := true

	for {
		v1 := <-ch
		v2, ok2 := <-ch2
		st1 = st1 + strconv.Itoa(v1) + " -> "
		st2 = st2 + strconv.Itoa(v2) + " -> "
		if v1 != v2 {
			same = false
		}
		if !ok2 {
			break
		}
	}

	fmt.Println("CH1 : ", st1)
	fmt.Println("CH2 : ", st2)
	fmt.Println("SAME : ", same)
}
