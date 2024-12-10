/*
q1. Create a slice with 3 random urls
    Create a function doGetRequest()
    doGetRequest:
        It spins up 2 goroutines
        1st goroutines do get request and
			put the url, resp, err inside one single channel
        //1st goroutine spins up n number of goroutines for n number of urls (fanout pattern)
        2nd goroutines fetch the values from the channel and perform following operations
            -check err
            -read body
            -check if status code above 299
            -and print url resp.Status
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type result struct {
	url        string
	resp *http.Response
	err        error
}

func main() {
	urls := []string{
		"https://example.com",
		"https://randomuser.me/api/test",
		"https://httpbin.org/status/404",
	}
	doGetRequest(urls)

}

func doGetRequest(urls []string) {
	ch := make(chan result)

	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)

	wg.Add(1)
	go func() { // 1 goroutine
		defer wg.Done()

		for _, url := range urls {

			wgWorker.Add(1)
			go func(u string) {
				defer wgWorker.Done()
				res, err := http.Get(u)
				ch <- result{url: u, resp: res, err: err}
			}(url)

		}

		wgWorker.Wait()
		close(ch) // close the channel in the sender goroutine
	}()

	wg.Add(1)
	go func() { // 2 goroutine

		defer wg.Done()

		for response := range ch {
			if response.err != nil {
				fmt.Println(response.err)
				continue 
			}
			defer response.resp.Body.Close()
			_, err = io.ReadAll(response.resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			if response.err != nil {
				fmt.Printf("Error fetching URL %s: *****%v\n", response.url, response.err)
				continue
			}

			if response.resp.statusCode > 299 {
				fmt.Printf("Status greater than 299 returned for URL %s and status is ***** %d\n", response.url, response.resp.statusCode)
				continue
			} else {
				fmt.Printf("URL %s returned status: ***** %d\n", response.url, response.resp.statusCode)
			}
		}
	}()

	wg.Wait()
}
