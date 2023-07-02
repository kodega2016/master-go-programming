package main

import "fmt"

func main() {
	var x *int
	y := 10
	x = &y

	fmt.Println("address:", x, "value:", *x)
}
