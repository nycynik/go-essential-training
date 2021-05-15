package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User 	string 	`json:"user"`
	Action 	string 	`json:"action"`
	Count 	int 	`json:"count"`
}

func postHTTPBin() {
	// POST req
	job := &Job{
		User: "saitama",
		Action: "punch",
		Count: 1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: Cant encode job - %s", err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}

func getHTTPBin()  {

	// GET Request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

func main()  {
	fmt.Println("posting")
	postHTTPBin()
	
	fmt.Println("getting")
	getHTTPBin()
}