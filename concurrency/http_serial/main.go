package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down", url)
	} else {
		defer res.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, res.StatusCode)

		if res.StatusCode == 200 {
			bs, err := ioutil.ReadAll(res.Body)

			if err != nil {
				fmt.Println(err)
			} else {

				filename := strings.Split(url, "//")[1]
				filename += ".txt"

				fmt.Printf("Writing response body to %s\n", filename)
				err := os.WriteFile(filename, bs, 0644)

				if err != nil {
					log.Fatal(err)
				}

			}
		}
	}

	fmt.Println(strings.Repeat("#", 40))
	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}

	fmt.Println("total number of go routines", runtime.NumGoroutine())
	wg.Wait()

}
