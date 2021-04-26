package main

import (
	"fmt"

	"github.com/jinwook-song/learngo-2021/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first":"First word"}

	baseWord := "hello"

	// Search
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// Add none exist word
	err2 := dictionary.Add(baseWord, "greeting")
	if err2 != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search(baseWord)
	fmt.Println(hello)

	// Add exist word -> error
	err3 := dictionary.Add(baseWord, "greeting")
	if err3 != nil {
		fmt.Println(err3)
	}

	// Update word
	err4 := dictionary.Update(baseWord, "edit")
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(dictionary[baseWord])

	// Delete word
	dictionary.Delete(baseWord)
	fmt.Println(dictionary)
}