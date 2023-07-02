package main

import "fmt"

func main() {
	friends := map[string]string{
		"name": "sakar subedi",
		"role": "software developer",
	}

	neighbors := friends
	friends["address"] = "Madhumalla Morang"
	fmt.Println(neighbors)

	delete(neighbors, "name")
	fmt.Println(neighbors, friends)

}
