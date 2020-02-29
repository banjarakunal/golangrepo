package main

import "fmt"

//Queue represents FIFO datastructure
type Queue struct {
	slice []int
}

//Enqueue adds element in queue
func (q *Queue) Enqueue(e int) {
	q.slice = append(q.slice, e)
}

//Dequeue returns last in element in queue
func (q *Queue) Dequeue() int {
	ret := q.slice[0]
	q.slice = q.slice[1:]
	return ret
}

func main() {
	q := Queue{
		slice: []int{},
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println("After enqueuing...", q.slice)

	e := q.Dequeue()
	fmt.Println("After dequeue element = %v  from queue = %v...", e, q.slice)

}
