package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"s", "b"}
	for i, _ := range a {
		a[i] = "a"
	}
	strings.Split(a,',')
}
