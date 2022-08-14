// Multiprocessing in go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// gorouting anonymous function, declare and directly run
	go func() {
		count("sheep")
		wg.Done()
	}()

	// working with channels

	wg.Wait()
}

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, name)
		time.Sleep(time.Millisecond * 500)
	}
}

func channelCount(name string, c chan string) {
	for i := 1; i <= 5; i++ {
		// instead of printing on console, we're feeding to a channel here
		c <- name
		time.Sleep(time.Millisecond * 500)
	}
}
