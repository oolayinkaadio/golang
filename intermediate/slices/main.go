package main

import "fmt"

/**
 * Slices are like array but they can accept more data as they are not limited by size[no size specification during initialization]
 */
func main() {
	// Small shopping list:
	shoppingList:= []string{"milk"}
	fmt.Println("Initial shopping List", shoppingList)

	// Adding more items to the shopping list slice:
	shoppingList = append(shoppingList, "bread")
	shoppingList = append(shoppingList, "eggs")
	fmt.Println("Bigger shopping List", shoppingList)
}