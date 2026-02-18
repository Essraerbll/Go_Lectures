package main

import "fmt"

func example() {
	var age int = 30
	var name string = "Alice"
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	height := 1.75
	fmt.Printf("Height: %.2f meters\n", height)

	isStudent := true
	fmt.Printf("Is Student: %t\n", isStudent)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
