
package main

import (
	"fmt"
	"strings"

)

func main() {
	var p int
	var x string

	fmt.Println("Please enter a string, followed by an integer, p, which is LESS THAN or EQUAL TO the total number of characters in the string:")
	fmt.Scan(&x, &p)
	fmt.Println("\nYour string:", x, "\nYour int:", p, "\nNow outputting the character at position", p, "in string", x, "using 0-based indexing\n\n")

	s := fmt.Sprintf(string(x[p]))
	fmt.Println(s)

	fmt.Println("\nNow outputting your original string,", x, "without the character at position", p, "\n")
	
	t := strings.Replace(x, s, "", -1)
	fmt.Println(t)


}
