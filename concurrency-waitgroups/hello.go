// Multiprocessing in go using waitgroups

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// gorouting anonymous function, declare and directly run
	go func() {
		count("sheep")
		wg.Done()
	}()

	go func() {
		count("goat")
		wg.Done()
	}()

	// wait till Wgs gets finished

	wg.Wait()
}

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, name)
		time.Sleep(time.Millisecond * 500)
	}
}
