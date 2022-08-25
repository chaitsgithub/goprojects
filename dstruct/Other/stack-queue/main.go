package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	} else {
		l := len(s.items) - 1
		popNum := s.items[l]
		s.items = s.items[:l]
		return popNum
	}
}

type Queue struct {
	items []int
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

func main() {
	myStack := Stack{}
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(50)
	myStack.Push(30)
	fmt.Println(myStack)
	myStack.Pop()
	myStack.Pop()
	fmt.Println(myStack)

	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(50)
	myQueue.Enqueue(30)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
