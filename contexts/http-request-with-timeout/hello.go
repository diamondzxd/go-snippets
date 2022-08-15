package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// -------------------------------------------------------------------------------------
	// The previous was a simple HTTP request, now we'll add our own timeout that if it exceeds
	// 100 ms, we'll terminate the request.

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://oldlappy.com", nil)
	if err != nil {
		panic(err)
	}

	// perform HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// print out a header of the http request
	fmt.Println(res.Header["Server"])

	// read body's data from HTTP Response
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(len(body))

}
