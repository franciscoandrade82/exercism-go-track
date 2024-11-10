package main

import "fmt"

// Item interface defines common methods for all items in the library
type Item interface {
	GetTitle() string
	GetAuthor() string
	GetYear() int
}

// CommonFields contains fields common to all items
type CommonFields struct {
	Title  string
	Author string
	Year   int
}

func (c CommonFields) GetTitle() string  { return c.Title }
func (c CommonFields) GetAuthor() string { return c.Author }
func (c CommonFields) GetYear() int      { return c.Year }

// Book struct contains fields specific to books
type Book struct {
	CommonFields
	ISBN string
}

// Magazine struct contains fields specific to magazines
type Magazine struct {
	CommonFields
	IssueNumber int
}

// DVD struct contains fields specific to DVDs
type DVD struct {
	CommonFields
	Duration int // in minutes
}

func main() {
	book := Book{
		CommonFields: CommonFields{
			Title:  "Go Programming",
			Author: "John Doe",
			Year:   2021,
		},
		ISBN: "123-456-789",
	}

	magazine := Magazine{
		CommonFields: CommonFields{
			Title:  "Tech Today",
			Author: "Jane Smith",
			Year:   2022,
		},
		IssueNumber: 42,
	}

	dvd := DVD{
		CommonFields: CommonFields{
			Title:  "Learning Go",
			Author: "Alice Johnson",
			Year:   2023,
		},
		Duration: 120,
	}

	items := []Item{book, magazine, dvd}

	for _, item := range items {
		fmt.Printf("Title: %s, Author: %s, Year: %d\n",
			item.GetTitle(), item.GetAuthor(), item.GetYear())
	}
}
