package main

import "fmt"

type lru struct {
	size int
	data map[int]int
	list []int
}

func newLru(size int) *lru {
	return &lru{
		size: size,
		data: make(map[int]int),
		list: make([]int, 0),
	}
}

func (l *lru) get(key int) int {
	v, ok := l.data[key]
	if !ok {
		return -1
	}

	findKeyIndex := 0
	for i, ak := range l.list {
		if ak == key {
			findKeyIndex = i
			break
		}
	}
	l.list = append(l.list[:findKeyIndex], l.list[findKeyIndex+1:]...)
	l.list = append(l.list, key)

	return v
}

func (l *lru) set(key, value int) {
	if len(l.list) < l.size {
		l.list = append(l.list, key)
		l.data[key] = value
		return
	}

	rmKey := l.list[0]
	delete(l.data, rmKey)
	l.list = l.list[1:]

	l.data[key] = value
	l.list = append(l.list, key)
}

func main() {
	l := newLru(10)
	for i := 1; i < 15; i++ {
		l.set(i, i+1)
	}

	fmt.Println(l.get(8))
	l.set(3, 100)
	fmt.Println(l.get(3))
}
