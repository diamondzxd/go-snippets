// Nested Structs and Objects

package main

import (
	"fmt"
)

type name struct {
	first string
	last  string
}

type student struct {
	name   name
	age    int
	gender string
	active bool
}

func main() {
	piyush := student{name: name{first: "Piyush", last: "Dhall"}, age: 21, gender: "M", active: true}
	fmt.Println(piyush.name.first)
}
