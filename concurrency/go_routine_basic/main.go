package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution")

	for i := 0; i < 4; i++ {
		fmt.Println("f1,i=", i)
		time.Sleep(time.Second)
	}
	wg.Done()
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

	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)
	fmt.Println("No of Goroutines after f1()", runtime.NumGoroutine())

	f2()

	wg.Wait()

	// time.Sleep(time.Second * 2)
	fmt.Println("main execution stopped...")
}
