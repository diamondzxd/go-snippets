package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// create HTTP request

	req, err := http.NewRequest(http.MethodGet, "https://place-hold.it/2000x2000", nil)
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

	// -------------------------------------------------------------------------------------
	// The above was a simple HTTP request, now we'll add our own timeout that if it exceeds
	// 100 ms, we'll terminate the request.

}
