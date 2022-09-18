package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	testList := make([]int, 0)
	for i := 0; i < 1000; i++ {
		v, _ := rand.Int(rand.Reader, big.NewInt(1000))
		testList = append(testList, (int)(v.Int64()))
	}

	heapSort(&testList)
	fmt.Println(testList)
	fmt.Println("done")
}

func heapSort(list *[]int) {
	for i := len(*list) / 2; i >= 0; i-- {
		initHeap(i, *list)
	}

	r := make([]int, 0)
	for len(*list) != 0 {
		v := popHeap(list)
		r = append(r, v)
	}

	*list = r
}

func initHeap(index int, list []int) {
	if index >= len(list) {
		return
	}

	left := 2*index + 1
	right := 2*index + 2
	if left < len(list) && list[left] < list[index] {
		list[index], list[left] = list[left], list[index]
	}
	if right < len(list) && list[right] < list[index] {
		list[index], list[right] = list[right], list[index]
	}

	initHeap(left, list)
	initHeap(right, list)
}

func popHeap(list *[]int) int {
	r := (*list)[0]
	(*list)[0], (*list)[len(*list)-1] = (*list)[len(*list)-1], (*list)[0]
	*list = (*list)[:len(*list)-1]

	if len(*list) >= 2 && (*list)[0] > (*list)[1] {
		(*list)[0], (*list)[1] = (*list)[1], (*list)[0]
		initHeap(1, *list)
	}
	if len(*list) >= 3 && (*list)[0] > (*list)[2] {
		(*list)[0], (*list)[2] = (*list)[2], (*list)[0]
		initHeap(2, *list)
	}

	return r
}
