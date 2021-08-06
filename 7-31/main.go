package main

import (
	"fmt"
	"reflect"
)

type a interface {
	b()
}

type test1 struct {

	name string
}

type test2 struct {

	user string
}

func (d test1) b() {
	fmt.Println("test")

}

func (d test2) b()  {
	fmt.Println("test2")

}

func e(inter a) a{

	return inter
}


func main() {
	switch 	 a := e(test2{user: "jjp"}).(type) {

	default:
		fmt.Printf("%T",a)
	}

	fmt.Println(reflect.TypeOf(e(test2{user: "jjp"})))


}
