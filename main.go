package main

import "fmt"

func main() {
	jinwook := map[string]string{"name":"jw", "age":"29"}
	fmt.Println(jinwook)

	for key, value := range jinwook {
		fmt.Println(key, value)
	}
}
