package main

import "fmt"

func main() {
	// c1 := make(chan string)

	c2 := make(chan string, 3)

	// fmt.Println(<-c2)
	// fmt.Println(<-c1)

	go func() {
		c2 <- "Hello how are you doing?"
		c2 <- "I am doing well, thank you!"
		c2 <- "Good to hear, I am doing well too!"
	}()

	fmt.Println("main go routine starts receiving data")
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	// fmt.Println(<-c2)

}
