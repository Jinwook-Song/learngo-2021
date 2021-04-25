package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood:= []string{"kimchi","ramen"}
	nico:= person{
		name:"nico",
		age:29,
		favFood: favFood}
	fmt.Println(nico)
}
