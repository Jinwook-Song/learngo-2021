package main

import "fmt"

func main() {
	a := 2
	b := &a	// a의 메모리 주소

	a = 5;
	fmt.Println(*b)
	
	*b = 10 // *b 메모리 주소가 가르키는 값
	fmt.Println(a)
}
