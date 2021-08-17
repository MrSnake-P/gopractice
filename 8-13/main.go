package main

import "fmt"

///*
//#include <stdio.h>
//
//void printint(int v) {
//    printf("printint: %d\n", v);
//}
//*/
//import "C"
//
//func main() {
//	v := 42
//	C.printint(C.int(v))
//}
type Foo interface {
	Say()
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

func main() {
	var _ Foo = Dog{}
	var _ Foo = (*Dog)(nil)

	fmt.Printf("%T", (*Dog)(nil))
	fmt.Println((*Dog)(nil))
	fmt.Println(&Dog{}==(*Dog)(nil))
}