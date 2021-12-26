package main

import (
	"errors"
	"fmt"
)

type queue []string

var queueUnderflowErr = errors.New("Queue underflow exception")

func (q *queue) Enqueue(s string) {
	*q = append(*q, s)
}

func (q *queue) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "", queueUnderflowErr
	}
	front := (*q)[0]
	*q = (*q)[1:len(*q)]
	return front, nil
}

func (q *queue) Front() string {
	return (*q)[0]
}

func (q *queue) Rear() string {
	if len(*q) == 0 {
		return ""
	}
	return (*q)[len(*q)-1]
}

type QueueOps interface {
	Enqueue(string)
	Dequeue() (string, error)
	Front() string
	Rear() string
}

func GetQueue() QueueOps {
	return &queue{}
}

func main() {
	q := GetQueue()
	q.Enqueue("Hello")
	q.Enqueue("World")
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
}
