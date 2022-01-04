package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	lock *sync.Mutex
	Values []int
}

func Init() Queue {
	return Queue{&sync.Mutex{}, make([]int, 0)}
}

func (q *Queue) Enqueue(x int) {
	for {
		q.lock.Lock()
		q.Values = append(q.Values, x)
		q.lock.Unlock()
		return
	}
}

func (q *Queue) Dequeue() *int {
	for {
		if (len(q.Values) > 0) {
			q.lock.Lock()
			x := q.Values[0]
			q.Values = q.Values[1:]
			q.lock.Unlock()
			return &x
		}
		return nil
	}
	return nil
}

func main() {
	graph := [][]int {
		{1, 2},
		{3},
		{4},
		{5},
		{},
		{},
	}
	queue := Init()

	fmt.Print("Iterative -> ")
	queue.Enqueue(0)
	bfsIter(graph, &queue)

	fmt.Print("\nRecursive -> ")
	bfsRec(graph, 0)
}

func bfsRec(graph [][]int, start int) {

}

func bfsIter(graph [][]int, queue *Queue) {
	for {
		if front := queue.Dequeue(); front != nil {
			fmt.Printf("%d -> ", *front)
			for _, val := range graph[*front] {
				queue.Enqueue(val)
			}
		}
	}
}
