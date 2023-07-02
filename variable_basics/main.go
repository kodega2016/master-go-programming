package main

import "fmt"

func main() {
	// variable using var keyword
	var name string = "Master Go Programming"
	fmt.Println(name)

	subject := "Exploring Variables in Go"
	fmt.Println(subject)

	// multiple variable declaration
	product, price := "Laptop", 50000
	fmt.Println(product, price)

	var (
		employeeName string = "John"
		age          int    = 30
		salary       int    = 50000
	)

	fmt.Println(employeeName, age, salary)
}
