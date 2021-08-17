package main

import "fmt"

//func main() {
//	array := []int{1,2,3}
//	for _, r := range array {
//		fmt.Println(r)
//	}
//
//	fmt.Println(time.Now().UTC().AddDate(0, 1, 0).Format(""))
//	start := time.Now()
//	var sum int
//	for i:=0; i <= 10000000; i++ {
//		sum += i
//	}
//	end := time.Now().Sub(start)
//	fmt.Println(end.Microseconds())
//
//}
//type ConfigOne struct {
//	Daemon string
//}
//
//func (c *ConfigOne) String() string {
//	return fmt.Sprintf("print: %v", c)
//}

func main() {
	//var a int = 100
	//var b int = 200
	//b, a = a, b
	//fmt.Println(a,b)
	a := map[byte]byte{
		'{': '}',
	}
	fmt.Println(a['}'])
	fmt.Printf("%T", a['}'])
}
