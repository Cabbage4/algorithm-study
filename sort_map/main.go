package main

import (
	"algorithm-study/sort_map/sm"
	"fmt"
)

func main() {
	c := sm.New()
	for i := 0; i < 10; i++ {
		c.Add(i, i)
	}

	fmt.Println(c.Get(5))
	fmt.Println(c.Get(7))
}
