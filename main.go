package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func htmlScrape(url string, urlChan chan string) {
	resp, err := http.Get(url)

	if err != nil {

		s := fmt.Sprintf("%s Currently Down!\n", url)
		s += fmt.Sprintf("Error: %v  \n", err)

		
		urlChan <- s

	} else {

		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)

		if resp.StatusCode == 200 {
			htmlBody, err := ioutil.ReadAll(resp.Body)

			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Scraping Html Body and writing to %s\n", file)

			err = ioutil.WriteFile(file, htmlBody, 0664)
			if err != nil {

				s += "Error writing to file!\n"

				urlChan <- s
			}
		}
		s += fmt.Sprintf("%s Complete\n", url)

		urlChan <- s
	}
}

func main() {
	urls := []string{"https://go.dev", "https://www.google.com", "https://www.github.com"}

	urlChan := make(chan string)

	for _, url := range urls {
		go htmlScrape(url, urlChan)
		fmt.Println(<-urlChan)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) 


}



