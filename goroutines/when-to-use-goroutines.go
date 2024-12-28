package main

import (
	"fmt"
	"time"
)

// Simulating file download
func downloadFile(filename string) {
	fmt.Printf("Starting download: %s\n", filename)

	//  Simulate download time
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished downloading: %s\n", filename)
}

func Downloadfile() {
	// Without goroutines(Sequential downloads)
	fmt.Println("=== Sequential Downloads ======")
	start := time.Now()

	downloadFile("movie.mp4") //Takes 2 seconds
	downloadFile("music.mp3") // Takes 2 more seconds
	downloadFile("image.jpg") // Takes 2 more seconds

	fmt.Printf("Sequential Downloads Time: %v\n\n", time.Since(start))

	// With Goroutine [Parallel downloads]
	fmt.Printf("==== Parallel Downloads =====")
	start = time.Now()

	go downloadFile("movie.mp4") //Starts immediately
	go downloadFile("music.mp3") // Starts immediately
	go downloadFile("image.jpg") // Starts immediately

	time.Sleep(3 * time.Second) //wait for downloads
	fmt.Printf("Parallel Downloads Time: %v\n\n", time.Since(start))
}

// Handling Multiple Requests:::
type Order struct {
	CustomerID int
	OrderID int
	Items []string
}

func execOrder(order Order){
	fmt.Printf("Starting to process Order #%d for customer #%d\n", order.OrderID, order.CustomerID)
	
	// Simulate order processing:
	time.Sleep(2 * time.Second)

	fmt.Printf("Completed order #%d for customer #%d. Items: %v\n", order.OrderID, order.CustomerID, order.Items)
}

func ProcessOrder() {
	orders := []Order {
		{CustomerID: 1, OrderID: 101, Items: []string{"Book", "Pen"}},
		{CustomerID: 2, OrderID: 102, Items: []string{"Laptop", "mouse"}},
		{CustomerID: 3, OrderID: 103, Items: []string{"Coffee", "Tea"}},
	}

	fmt.Println("=== Processing Orders ====")
	start := time.Now()

	// Process all orders concurrently
	for _, order := range orders {
		go execOrder(order)
	}

	// wait for orders  to complete
	time.Sleep(3 * time.Second)
	fmt.Printf("\nTotal processing time: %v\n", time.Since(start))
}