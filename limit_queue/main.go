package main

import (
	"fmt"
	"time"
)

func main() {
	lq := &LimitQueue{
		queue:  make([]queueNode, 0),
		count:  5,
		window: 10 * time.Millisecond,
	}

	for i := 0; i < 50; i++ {
		fmt.Println(lq.Push(func() {}))
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println()
}

type LimitQueue struct {
	queue []queueNode

	count  int
	window time.Duration
}

type queueNode struct {
	f func()
	t time.Time
}

func (l *LimitQueue) Push(f func()) bool {
	t := time.Now()

	if len(l.queue) < l.count {
		l.queue = append(l.queue, queueNode{f: f, t: t})
		return true
	}

	if t.Sub(l.queue[0].t).Milliseconds() <= l.window.Milliseconds() {
		return false
	}

	l.queue = l.queue[1:]
	l.queue = append(l.queue, queueNode{f: f, t: t})
	return true
}
