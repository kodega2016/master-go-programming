package main

import "fmt"

func main() {
	var users []string = []string{
		"khadga", "sakar", "sabin",
	}
	fmt.Println(users)
	fmt.Println(users[0])
	fmt.Println(users[len(users)-1])
	fmt.Println("total users:", len(users))

	for index, user := range users {
		fmt.Println(index, user)
	}
}
