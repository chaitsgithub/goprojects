package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	items []int
	mutex sync.Mutex
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	} else {
		r := q.items[0]
		q.items = q.items[1:]
		return r
	}
}

func writeOddNumbers(q *Queue) {
	for i := 0; ; i++ {
		if i%2 != 0 {
			q.mutex.Lock()
			q.items = append(q.items, i)
			q.mutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}
}

func writeEvenNumbers(q *Queue) {
	for i := 0; ; i++ {
		if i%2 == 0 {
			q.mutex.Lock()
			q.items = append(q.items, i)
			q.mutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	myQueue := Queue{}
	go writeOddNumbers(&myQueue)
	go writeEvenNumbers(&myQueue)

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(myQueue.items)
	}
}
