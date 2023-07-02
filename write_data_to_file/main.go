package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var message string = "Hello i am khadga bahadur shrestha,working as a flutter developer at PortPro"
	var bs []byte = []byte(message)

	// fmt.Println("bytes:", bs)

	// write this byte to a file
	os.WriteFile("messages.txt", bs, 0666)

	// open file
	f, err := os.OpenFile("info.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("error:", err)
	}

	newInfo := "application is running..."
	f.Write([]byte(newInfo))

	defer f.Close()

	ff, _ := os.OpenFile("data.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	buffWriter := bufio.NewWriter(ff)

	bss := []byte("Nice application")

	bytes, err := buffWriter.Write(bss)
	// if err != nil {
	// 	log.Fatal("error:", err)
	// }
	bytesAvailable := buffWriter.AvailableBuffer()
	fmt.Println("available buffer:", bytesAvailable)
	fmt.Println("data written:", bytes)

	buffWriter.WriteString("hey")

}
