package main

import "fmt"

/**
 * Map can be seens as a dictionary.
 */

func main() {
// Creating a simple student map
 students := map[string]string{
		"S001": "John",
		"S002": "Alice",
		"S003": "Bob",
 }

//  Looking up a student with their id in the students map above:
fmt.Println("Student with id S001: " + students["S001"])

// Looking up a student with their id stored in a variable:
 alice := "S002"
 fmt.Println("Get student with id stored in variable:", students[alice], "\n\n\n")
 
 
	//  

	// Create an empty map
	contacts := make(map[string]string)

	// Adding new contacts
	contacts["John"] = "123-457"
	contacts["Alice"] = "789-013"

	// changing a number:
 	contacts["John"] = "999-999"

	// Removing a contact:
	// delete(contacts, "Alice")

	fmt.Println(contacts)



	fmt.Println("Checking if a variable exists in a map\n")
	// Checking if a value exists in a map:
	scores := map[string]int{
		"John": 85,
		"Alice": 92,
	}

	// checking if Bob's score exists:
	scoreForBob,exists := scores["Bob"]
	if exists {
		fmt.Println("Bob's score exists: ", scoreForBob)
	} else {
		fmt.Println("Bob's score does not exist!!!")
	}


	// Looping through a map:
fmt.Println("Looping through a map:\n")
	menu := map[string]float64{
		"Coffee": 2.50,
		"Tea": 2.00,
		"Cake": 3.50,
	}

	for item,price := range menu {
		fmt.Printf("%s consts $%.2f\n", item, price)
	}


	
}