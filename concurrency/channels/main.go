package main

import "fmt"

func main() {
	ch := make(chan string)

	go emitValue("Hey i am learning go", ch)

	value := <-ch

	fmt.Println("value:", value)

	close(ch)

	fmt.Println("exiting the main")
}

func emitValue(message string, c chan string) {
	c <- message
}
