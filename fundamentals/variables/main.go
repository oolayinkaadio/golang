package main

import "fmt"

func main() {
	// declare a variable 'age' with type int and value 25:
	var age int = 25

	// declare a variable 'name' without specifying type, Go automatically infers the type as string from the value:
	name := "Golang"

	fmt.Println(name, "is", age, "years old")
}