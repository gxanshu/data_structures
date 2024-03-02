package main

import "fmt"

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println("Queue size:", q.Size())
	fmt.Println("Dequeued item:", q.Dequeue())
	fmt.Println("Queue size after dequeue:", q.Size())
	fmt.Println("Is queue empty?", q.IsEmpty())
}
