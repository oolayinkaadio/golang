package main

import "fmt"

/**
 * Go uses 'for'  as the only looping statement. the 'for' can be used similarly to 'while' or 'do-while' loop as done in other programming languages
 */

func main() {
	// FOR LOOP:
	for i:= 0; i < 5; i++ {
		fmt.Println("FOR LOOP", i)
	}

	/**
	 * WHILE LOOP: 
	 * this is done by using the for loop without specifying initialization [it will run as long as the condition is true]
	 * -- for 'condition' {} -- is an infinite loop that runs indefinitely until a break condition is met
	 */

	a:= 0
	for a < 5 { // a < 5 is the condition
		fmt.Println("WHILE LOOP", a)
		a++;

		if(a >=5) {
			break
		}
	}

/**
 * DO-WHILE LOOP
 * This simulates the do-while loop functionality of other programming langauges, as the loop runs at least once before checking the condition and break if the condition is true/satisfied.
 */
	// DO-WHILE LOOP:
	b:= 0

	for {
		fmt.Println("DO-WHILE LOOP", b)
		b++;

		if(b >=5) { // condition
			break;
		}
	}



}