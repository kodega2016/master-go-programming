package main

import "fmt"

func main() {

	// int and float
	var i int = 10
	var salary float32 = 1000.0
	fmt.Printf("%T %T \n", i, salary)

	// runes types
	var r rune = 'a'
	fmt.Printf("%T/%v\n", r, r)

	// bool types
	var b bool = true
	fmt.Printf("%T\n", b)

	// string types
	var s string = "Hello World"
	fmt.Printf("%T\n", s)

	// arrays
	var users [2]string = [2]string{"John", "Doe"}
	fmt.Printf("%T\n", users)

	// slices
	var cities []string = []string{"London", "Paris", "New York"}
	fmt.Printf("%T\n", cities)

	// maps
	var m map[string]int = map[string]int{"foo": 1, "bar": 2}
	fmt.Printf("%T\n", m)

	// structs
	type User struct {
		ID        int
		FirstName string
	}

	var u User = User{ID: 1, FirstName: "John"}
	fmt.Printf("%T\n", u)

	// pointers
	var p *int = &i
	fmt.Printf("%T\n", p)

	// functions
	fmt.Printf("%T\n", f)

}

func f() {}
