package main

import "fmt"


func makeCoffee(coffeeType string,sugar int) string {
	return fmt.Sprintf("Here's your %s coffee with %d sugar!", coffeeType, sugar)
}
// func main() {
// 	myCoffee := makeCoffee("Latte", 2)
// 	fmt.Println(myCoffee)
// }



// Returning multiple values from a function::
func orderMeal(mealName string) (string, string) {
	food := mealName + " with rice"
	drink := "iced tea"

	return food, drink
}

func main() {
	myFood, myDrink := orderMeal("Chicken")

	fmt.Println("Food: ", myFood)
	fmt.Println("Dringk: ", myDrink)
}