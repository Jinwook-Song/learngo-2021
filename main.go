package main

import (
	"fmt"
	"time"
)

func main() {
	// create channel
	channel := make(chan string)
	people := [2]string{"nico", "jw"}
	for _, person := range people {
		// connect channel
		go isNice(person, channel)
	}
	fmt.Println("Waiting for messages...")
	// block main func until receive something from channel
	results := []string{}
	for i:=0; i<len(people); i++ {
		results = append(results, <- channel)
	}

	for idx := range results {
		fmt.Println(results[idx])
	}
}

// connect channel
func isNice(person string, channel chan string) {
	time.Sleep(time.Second * 3)
	// return(send) to main
	channel <- person + " is nice"
}