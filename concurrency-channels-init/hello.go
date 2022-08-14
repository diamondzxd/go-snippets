package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	// for {
	// 	msg := <-c
	// 	fmt.Println(msg)
	// }
	// we'll receive a fatal error during runtime if the channel doesn't closes in the count function
	// because this goroutine will keep waiting for,
	// but there will be no one sending it,
	// hence, it'll go to sleep and cause a deadlock

	// to fix these issues -->

	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}

	// syntactically, a better way

	cNew := make(chan string)
	go count("goat", cNew)
	for msg := range cNew {
		// iterating over a for loop on channel will keep sending out data
		// until the channel is closed
		fmt.Println(msg)
	}

}

func count(name string, c chan string) {
	for i := 1; i <= 5; i++ {
		// instead of printing on console, we're feeding to a channel here
		c <- name
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
	// closing the channel after we're done sending the data
}
