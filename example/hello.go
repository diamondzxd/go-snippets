package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello")

	x := 5
	y := 1

	sum := x + y

	fmt.Println(sum)

	if sum > 10 {
		fmt.Println("Congrats! You have got double digit package! :D")
	} else if sum < 10 {
		fmt.Println("Keep up with the good work! one day :)")
	}

	// Arrays

	var a [5]int
	a[3] = 10

	//alternatively
	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(b)

	// however, arrays are of fixed length, defined in the type itself.
	// To make dynamic arrays, we can use slices...

	aSlice := []int{1, 2, 3, 4, 5}
	aSlice = append(aSlice, 6)
	// append creates a new slice and returns the copy.

	fmt.Println(aSlice)

	// Maps

	// Syntax :
	// map[key_type]value_type

	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["pentagon"] = 5

	delete(vertices, "square")

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	// Loops
	fmt.Println("---------------Loops------------------")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// As a while Loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Iterating over an array

	arr := []int{100, 200, 300, 400, 500}

	for index, value := range arr {
		println(index, " -> ", value)
	}

	// Iterating over a map

	store := map[string]int{"aa": 10, "bb": 20, "cc": 123}
	for key, value := range store {
		println(key, " ->", value)
	}

	fmt.Println(add(10, 20))

	// calling function and checking for errors

	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// using structs

	akhil := person{name: "Akhil", age: 31}
	fmt.Println(akhil)
	fmt.Println(akhil.name)

	// Pointers

	xyz := 45
	fmt.Println(xyz)

	inc(&xyz) // passing address of the int to function
	fmt.Println("after inc -> ", xyz)
}

type person struct {
	name string
	age  int
}

// Functions
func add(x int, y int) int {
	return x + y
}

// Function using pointers
// note how there is no return type. its similar to a void function
func inc(x *int) {
	// De-referencing the pointer
	*x++
}

// Functions with multiple return values

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
