package main

import (
	"fmt"
	"regexp"
)

type a struct {

	name string
}

func main() {
	str := "上海公司（山海关）"
	reg := regexp.MustCompile(`[^\p{Han}\(\)（）]+`)
	fmt.Println(len(str))
	fmt.Println(reg.MatchString(str))
	fmt.Println(reg.FindString(str))
	fmt.Println(reg.FindAllString(str, 2))
	//reg := regexp.MustCompile(`[\P{Han}]+`)
	b := []int{}
	fmt.Println(len(b))
}
