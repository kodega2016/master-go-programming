package main

import "fmt"

func main() {
	sum := getTotal(10, 20, 30, 4, 1)
	fmt.Println("sum:", sum)

}

func getTotal(args ...int) int {
	sum := 0
	for _, i := range args {
		sum += i
	}
	return sum
}
