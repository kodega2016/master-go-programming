package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	ch := make(chan int)
	defer close(ch)

	go factorial(n, ch)
	fmt.Printf("factorial of %d:%d\n", n, <-ch)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		fmt.Printf("fact of %d is %d\n", i, <-ch)
	}

}

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	c <- f
}
