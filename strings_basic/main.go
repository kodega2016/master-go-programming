package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hi go here"
	fmt.Println(s1)

	fmt.Println("He says:\"Hello buddy to us\"")
	fmt.Println(`He says:"Hello Buddy"`)

	s2 := `I like \n go!`
	fmt.Println(s2)

	fmt.Println(len(s1))

	fmt.Println(strings.ToUpper(s1))

}
