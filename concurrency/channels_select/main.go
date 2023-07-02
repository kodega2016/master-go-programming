package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, ch chan string) {
	res, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("%s is down", url)
	} else {
		defer res.Body.Close()
		message := fmt.Sprintf("%s -> Status Code: %d \n", url, res.StatusCode)

		if res.StatusCode == 200 {
			bs, err := ioutil.ReadAll(res.Body)

			if err != nil {
				fmt.Println(err)
			} else {

				filename := strings.Split(url, "//")[1]
				filename += ".txt"

				message += fmt.Sprintf("Writing response body to %s\n", filename)
				os.WriteFile(filename, bs, 0644)
				ch <- message

			}
		}
	}

	fmt.Println(strings.Repeat("#", 40))
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}
	ch := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, ch)
	}

	fmt.Println("total number of go routines", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}

}
