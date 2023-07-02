package main

import "fmt"

func main() {
	a := map[string]string{
		"A": "a",
		"B": "b",
	}

	b := map[string]string{
		"A": "A",
		"B": "b",
	}

	fmt.Println("a:", a, "b:", b)
}
