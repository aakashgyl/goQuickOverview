package main

import "fmt"

// Embedded interfaces behaves same as child inhering parent methods
// Any struct_object can be assigned to interface it implements and can call those methods

type AuthorDetails interface {
	details()
}

type AuthorArticles interface {
	articles()
}

type FinalDetails interface {
	AuthorDetails
	AuthorArticles
}

type author struct {
	a_name       string
	branch       string
	year         int
	articleCount int
}

func (a author) details() {
	fmt.Printf("Author Name: %s\n", a.a_name)
	fmt.Printf("Branch: %s and passing year: %d\n", a.branch, a.year)
}

func (a author) articles() {
	fmt.Println("Article count: ", a.articleCount)
}

func main() {
	values := author{
		a_name: "Mickey",
		branch: "Computer science",
		year:   2012,
	}

	var f FinalDetails = values
	f.details()
	f.articles()

	var a AuthorDetails = values
	a.details()
}
