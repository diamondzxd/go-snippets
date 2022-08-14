package main

import "fmt"

var cache = make(map[int]int)

func main() {
	var a int
	fmt.Scanln(&a)
	fmt.Println(fib(a))

}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if cache[n] != 0 {
		return cache[n]
	}

	result := fib(n-1) + fib(n-2)
	cache[n] = result
	return result
}
