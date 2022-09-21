// Take name as input and print a greeting message for that particular name.
package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Println("Enter Your Name:")
	fmt.Scanln(&a)
	fmt.Printf("Hello %s\n", a)
}
