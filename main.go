package main

import "fmt"

func main() {
	names := []string{"nico","lynn","jinwook"}
	// array를 return하지 않으므로 값을 할당해야 함
	names = append(names, "golang") 
	fmt.Println(names)
}
