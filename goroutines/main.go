package main

import (
	"fmt"
	"time"
)

func makeBreakfast(item string) {
	fmt.Printf("Starting to make %s\n", item)
	time.Sleep(2 * time.Second) // cooking time
	fmt.Printf("Finished making %s\n", item)
}

func main() {
	// Without goroutines (sequential and slow way):
	makeBreakfast("eggs") // wait 2 minutes
	makeBreakfast("coffee") // wait another 2 minutes
	// Total: 4 mins

	// With goroutines (fast way):
	go makeBreakfast("eggs") // start cooking eggs
	go makeBreakfast("coffee") // start making coffee immediately
	// Total: 2mins

	time.Sleep(3 * time.Second)

	fmt.Println('\n')
	Downloadfile()
	fmt.Println('\n')

	ProcessOrder()

}