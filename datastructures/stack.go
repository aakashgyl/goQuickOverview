package main

import (
	"errors"
	"fmt"
)

var stackOverflowException = errors.New("stack overflow")
var stackUnderflowException = errors.New("stack underflow")

type stack struct {
	arr [10]int
	top int
}

type StackOps interface {
	Push(int) error
	Pop() (int, error)
	Print()
	Empty() bool
	Size() int
	Top() int
}

func GetStack() StackOps {
	return &stack{
		arr: [10]int{},
		top: -1,
	}
}

func (s *stack) Push(i int) error {
	if s.top+1 == len(s.arr) {
		return stackOverflowException
	}
	s.top++
	s.arr[s.top] = i
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.Empty() {
		return 0, stackUnderflowException
	}
	popVal := s.Top()
	s.top--
	return popVal, nil
}

func (s *stack) Print() {
	i := s.top
	for i > -1 {
		fmt.Printf("%d -> ", s.arr[i])
		i--
	}
	fmt.Println()
}

func (s *stack) Empty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *stack) Size() int {
	return s.top + 1
}

func (s *stack) Top() int {
	return s.arr[s.top]
}

func main() {
	s := GetStack()
	fmt.Println(s.Pop())
	fmt.Println("isEmpty -> ", s.Empty())
	fmt.Println("print stack: ")
	s.Print()
	s.Push(1)
	s.Push(2)
	fmt.Println("top: ", s.Top())
	fmt.Println("print stack: ")
	s.Print()
	fmt.Printf("pop: ")
	fmt.Println(s.Pop())
	fmt.Println("print stack: ")
	s.Print()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
