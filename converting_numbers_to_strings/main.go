package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strVal string = "12"
	intVal, _ := strconv.Atoi(strVal)
	fmt.Println(intVal)

	var intVal2 int = 12
	strVal2 := strconv.Itoa(intVal2)
	fmt.Printf("%T\n", strVal2)

	var score int = 129
	var strScore = fmt.Sprintf("%d", score)
	fmt.Println("strScore:", strScore)
}
