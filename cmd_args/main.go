package main

import (
	"fmt"
	"os"
)

func main() {
	var args = os.Args
	fmt.Println(args[1], args[2])
}
