package main

import (
	"errors"
	"fmt"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val int) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack underflow")
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, nil
	}
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
	stack := Stack{}

	fmt.Print("Iterative -> ")
	stack.Push(0)
	dfsIter(graph, &stack)

	fmt.Print("\nRecursive -> ")
	dfsRec(graph, 0)
}

func dfsRec(graph [][]int, start int) {
	fmt.Printf("%d -> ", start)
	for _, val := range graph[start] {
		dfsRec(graph, val)
	}
}

func dfsIter(graph [][]int, stack *Stack) {
	for !stack.IsEmpty() {
		top, _ := stack.Pop()
		fmt.Printf("%d -> ", top)
		for _, val := range graph[top] {
			stack.Push(val)
		}
	}
}
