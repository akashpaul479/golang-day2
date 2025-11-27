package day3

import "fmt"

func Struct2() {
	type Book struct {
		Title  string
		Author string
		price  float64
	}
	var books []Book
	books = append(books, Book{
		Title:  "go programming",
		Author: "john smith",
		price:  846,
	})
	books = append(books, Book{
		Title:  "java",
		Author: "Herbert Schildt",
		price:  980,
	})
	fmt.Println("List of books are:")
	for _, m := range books {
		fmt.Printf("Title:%s , Author:%s , Price:%.2f\n", m.Title, m.Author, m.price)
	}

}
