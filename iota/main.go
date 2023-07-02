package main

import "fmt"

func main() {
	const (
		a = iota + 2
		_
		c
	)

	fmt.Println(a, c)

	const (
		North = iota
		East
		South
		West
	)

	fmt.Println(North, East, South, West)
}
