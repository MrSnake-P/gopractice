package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {	// v已经变成m的副本引用，不是原来的地址了
		fmt.Println(k, "->", *v)
	}
}

// 正确的
//func main() {
//
//	slice := []int{0,1,2,3}
//	m := make(map[int]*int)
//
//	for key,val := range slice {
//		value := val				// <--
//		m[key] = &value
//	}
//
//	for k,v := range m {
//		fmt.Println(k,"===>",*v)
//	}
//}
