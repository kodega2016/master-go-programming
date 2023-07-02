package main

import "fmt"

func main() {
	defer clear()
	foo()
	bar()

}
func foo() {
	fmt.Println("this is foo")
}

func bar() {
	fmt.Println("this is bar")
}

func clear() {
	fmt.Println("clear resources...")
}
