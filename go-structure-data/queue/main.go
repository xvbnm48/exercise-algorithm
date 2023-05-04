package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:len(q.items)]
	return toRemove
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.items)

	q.Dequeue()
	fmt.Println(q.items)

	q.Dequeue()
	fmt.Println(q.items)
}
