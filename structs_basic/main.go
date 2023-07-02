package main

import "fmt"

type Book struct {
	title    string
	author   string
	year     int
	category Category
}

type Category struct {
	name, logo string
}

func main() {
	b := Book{
		title:  "flutter and firebase for MVP",
		author: "Reso Coders",
		year:   2022,
		category: Category{
			name: "Mobile App Development",
			logo: "mobile-app.png",
		},
	}
	fmt.Println(b.title, b.author, b.year, b.category.name)
}
