package main

import (
	"algorithm-study/avl_tree/avl"
	"fmt"
	"strconv"
)

func main() {
	t := new(avl.Tree)
	for i := 0; i < 100; i++ {
		t.Add(fmt.Sprintf("jerry%d", i), strconv.Itoa(i))
	}

	fmt.Println(t.Get("jerry50"))
	t.Remove("jerry50")
	fmt.Println(t.Get("jerry50"))
}
