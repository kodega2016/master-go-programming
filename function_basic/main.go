package main

import "fmt"

func main() {
	greeting()
	sayMessage("Hello from Golang")
	fmt.Println("sum:", sum(10, 20))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	m, s := calc(10, 20)
	fmt.Println("sum:", s, "mul:", m)
}

func greeting() {
	fmt.Println("greeting...")
}
func sayMessage(message string) {
	fmt.Println("message:", message)
}
func sum(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}

func calc(x, y int) (sum int, mul int) {
	sum = x + y
	mul = x * y
	return
}
