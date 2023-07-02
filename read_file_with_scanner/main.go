package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("message.txt")
	sc := bufio.NewScanner(f)

	defer f.Close()

	// success := sc.Scan()

	// if !success {
	// 	err := sc.Err()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// fmt.Println(sc.Text())

	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	vs := bufio.NewScanner(os.Stdin)
	vs.Scan()

	fmt.Println(vs.Text())

}
