package main

import (
	"fmt"
)

func main() {

	var n int
	var x string

	
	fmt.Printf("Please enter a positive integer, followed by a string: ")
	fmt.Scan(&n, &x)

	fmt.Println("\nYour int:", n,  "\nYour string:", x)
	fmt.Println("\nNow outputting", x, n, "times\n\n")

	i := 1
	for i <= n{
		fmt.Println(x)
		i = i + 1
	}


}

