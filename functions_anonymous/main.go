package main

import "fmt"

func main() {
	func(message string) {
		fmt.Println(message)
	}("Hey how are you?")
}
