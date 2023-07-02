package main

import "fmt"

func main() {
	var a = 12
	var b = 12.0

	fmt.Println(a, b)
	// a = b
	a = int(b)
	fmt.Println(a, b)

	var name string
	var isAtive bool
	var price float64
	var qtd int

	fmt.Println(name, isAtive, price, qtd)

}
