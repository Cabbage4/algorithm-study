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

	quickSort(testList)
	fmt.Println(testList)
}

func quickSort(list []int) {
	quickSortWithRange(0, len(list), list)
}

func quickSortWithRange(left, right int, list []int) {
	if left >= right {
		return
	}

	splitIndex := splitAndFindSplitIndex(left, right, list)
	quickSortWithRange(left, splitIndex, list)
	quickSortWithRange(splitIndex+1, right, list)
}

func splitAndFindSplitIndex(left, right int, list []int) int {
	ans := left
	target := list[left]
	left++

	for left < right {
		if list[left] > target {
			list[left], list[right-1] = list[right-1], list[left]
			right--
		} else {
			list[left], list[ans] = list[ans], list[left]
			ans = left
			left++
		}
	}

	return ans
}
