package main

import "fmt"

func main() {
	amit := Author{
		"Amit",
		32,
	}

	book := Book{
		name:   "hello world",
		Author: amit,
	}

	publisher := Publication{
		name:         "Govind",
		publishCount: 1,
		book:         book,
	}

	// using parent struct, can call methods of parent only
	fmt.Println(amit.getName())
	fmt.Println(amit.getAge())

	// using child struct, can call methods of parent + child
	fmt.Println(book.getName())
	fmt.Println(book.getAge())
	fmt.Println(book.getBookName())
	fmt.Println(book.Author.getName()) // if both have same field name, access like this

	// using composition, cant access composed data directly
	fmt.Println(publisher.getName())
	fmt.Println(publisher.getPublishCount())
	fmt.Println(publisher.getBook())
}

type Author struct {
	name string
	age  int
}

type Book struct {
	name string
	Author
}

type Publication struct {
	name         string
	publishCount int
	book         Book
}

func (a Author) getName() string {
	return a.name
}

func (a Author) getAge() int {
	return a.age
}

func (b Book) getBookName() string {
	return b.name
}

func (p Publication) getName() string {
	return p.name
}

func (p Publication) getPublishCount() int {
	return p.publishCount
}

func (p Publication) getBook() Book {
	return p.book
}
