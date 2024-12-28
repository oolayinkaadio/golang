package main

import "fmt"

/**
 * Arrays have specific length and type in Go
 */
func main(){
	fruitBox := [3]string {"apple", "orange", "banana"} //length = 3, type = string

	fmt.Println("Fruits in my box:", fruitBox)

	// Change Orange to Mango:
	fruitBox[1] = "mango"
	fmt.Println("Updated fruits:", fruitBox)
}