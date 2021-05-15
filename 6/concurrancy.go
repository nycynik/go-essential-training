package main

import (
	"fmt"
	"net/http"
	"sync"
)

func contentType(url string) (string, error) {
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error %s\n", err)
		return "", err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	fmt.Printf("%s is %s\n", url, ct)
	return ct, err
}


func main()  {
	fmt.Println("starting")

	urls := []string{
		"https://mikelynchgames.com",
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
		
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			contentType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}