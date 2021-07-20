package main

import "fmt"

type Role struct {
	name string
	age int
}

func main() {
	a:= []Role{{name: "name", age: 22}, {name: "w", age: 23}}
	fmt.Println(a)
	for _, b := range a {
		fmt.Println(b.name)
	}
}
