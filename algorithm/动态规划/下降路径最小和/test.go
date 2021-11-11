package main

import "fmt"

func slice() {
	a := []int{1,2,3}
	b := a[:2]
	fmt.Println(b)
	var c int8 = -1
	var d int8 = -128 / c
	fmt.Println(d)
}