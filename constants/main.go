package main

import "fmt"

func main() {
	const PI = 3.14
	res := PI * 12
	fmt.Println("result:", res)

	const a, b = 10, 20
	// a = 12
	// b = 21
	fmt.Println("a:", a, "b:", b)

	const y = 10
	const x = y * 10
	fmt.Println("x:", x, "y:", y)
}
