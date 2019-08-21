package main

import "fmt"

var b bool

func main() {
	fmt.Println(b)

	s := "xixi"
	fmt.Println(s)

	//var i = 87
	//fmt.Println(i)

	if 5 > 4 {
		fmt.Println("5要比4大哦!")
	}

	var pointer1 = &s
	fmt.Println(pointer1)
	fmt.Println(*pointer1)

	var pointer2 = &pointer1
	fmt.Println(pointer2)
	fmt.Println(*pointer2)
	fmt.Println(**pointer2)
	fmt.Println(&pointer2)
}
