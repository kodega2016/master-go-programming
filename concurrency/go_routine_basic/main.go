package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (goroutine) execution")

	for i := 0; i < 4; i++ {
		fmt.Println("f1,i=", i)
	}
}

func f2() {
	fmt.Println("f2 (goroutine) execution")

	for i := 0; i < 4; i++ {
		fmt.Println("f2,i=", i)
	}
}

func main() {
	fmt.Println("main function execution starts here...")
	fmt.Println("=======================================")
	fmt.Println("No.of CPUs:", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("Os:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(8))

	go f1()
	fmt.Println("No of Goroutines after f1()", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)
	fmt.Println("main execution stopped...")
}
