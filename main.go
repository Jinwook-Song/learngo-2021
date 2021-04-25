package main

import (
	"fmt"
	"strings"
)

func lenAndUpper (name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // return 이후 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}


func main() {
	totalLength, uppercase := lenAndUpper("jinwook")
	fmt.Println(totalLength, uppercase)
}