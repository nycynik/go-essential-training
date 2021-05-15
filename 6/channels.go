package main

import (
	"fmt"
	"time"
)


func main()  {
	ch := make(chan int) 

	/*
	// Example: this will block since were not sending in a value
	<-ch 
	fmt.Println("here")
	*/

	go func() {
		ch <- 353 // send a number to channel
	}()

	// receive from the channel
	val := <-ch
	fmt.Printf("got %d\n", val)

	// ==
	fmt.Println("---")
	
	// send multiple
	go func() {
		for i := 0; i<3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	// ==
	fmt.Println("---")

	// send multiple with close
	go func() {
		for i := 0; i<3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Printf("received %d\n", val)
	}
}