package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("message.txt")

	if err != nil {
		log.Fatal(err)
	}

	byteSlice := make([]byte, 200)
	nbs, _ := io.ReadFull(f, byteSlice)

	fmt.Println("total bytes:", nbs)
	fmt.Println("data read:", string(byteSlice))

	defer f.Close()

}
