package main

import "fmt"

func main() {
	price, inStock := 900, true

	if price > 1000 {
		fmt.Println("Too Expensive !!!")
	}

	if price < 1000 && inStock {
		fmt.Println("Buy it !!!")
	}

	if price < 100 {
		fmt.Println("Too Cheap !!!")
	} else if price == 100 {
		fmt.Println("On the edge !!!")
	} else {
		fmt.Println("Don't Buy !!!")
	}

}
