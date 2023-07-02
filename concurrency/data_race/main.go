package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var gr = 100
	var n int = 0

	var m sync.Mutex

	wg.Add(gr * 2)

	go func() {
		for i := 0; i < gr; i++ {
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}
	}()

	go func() {
		for i := 0; i < gr; i++ {
			m.Lock()
			n--
			m.Unlock()
			wg.Done()
		}
	}()

	wg.Wait()

	fmt.Println("final value of n", n)

}
