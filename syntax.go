// used to define an executable program
package main

// Add all modules here
import (
	"fmt"  // This module is used to print something
)

func main(){
	
	// 1. Print something in terminal
	
	fmt.Println("hello world!")

	// 2. Variables and constants (Storing Information)
	
	var name string = "Alice"
	age := 20  // Shorter way to declare a variable (Go figures out the type by itself)
	fmt.Println("Name:", name, "Age:", age)
	const PI = 3.14
	fmt.Println("Value of PI:", PI)

	// 3. Data Types (Numbers, Strings, Booleans)
	var x int = 10  // Integer
	var y float64 = 20.5  // Floating point number
	var am_I_awesome bool = true  // Boolean
	fmt.Println("Integer:", x, "Float:", y, "Boolean:", am_I_awesome)

	// 4. If-else
	if age >= 18 {
		fmt.Println("You are an adult, welcome!")
	} else {
		fmt.Println("GO study for JEE, ain't no way you doing GO right now!")
	}

	// 5. Loops
	for i := 1; i <= 5; i++ {
		fmt.Printf("Hello %d time \n", i)
	}

	// Note that there is nothing called "while" loop in GO, we use "for" loop instead!

	// 6. Arrays
	var numbers = [3]int{10, 20, 30}  // It's fine even if you don't write the '3' in brackets
	fmt.Println("Array elements:", numbers)

	// 7. Functions
	greet("Alice")
}


func greet(name string) {
	fmt.Println("Hello,", name, "Welcome to Go!")
}
