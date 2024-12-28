package main

import "fmt"

// Creating Custom Types with Structs:

type GameCharacter struct {
	Name			string
	Health		int
	IsPowered	bool
}

func main() {
	// Createing a new character
	hero := GameCharacter{
		Name: "Super Nova",
		Health: 100,
		IsPowered: true,
	}

	fmt.Printf("Hero Name: %s\n", hero.Name)
	fmt.Printf("Hero Health: %d\n", hero.Health)
}