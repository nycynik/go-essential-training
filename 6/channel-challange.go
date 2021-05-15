package main

import (
	"fmt"
	"net/http"
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

	ch := make(chan string) 

	for _, url := range urls {
		ct, err := contentType(url)
		if err != nil {
			ch <- ct
		}
	}
	close(ch)

	for val := range ch {
		fmt.Printf("received %s\n", val)
	}

}