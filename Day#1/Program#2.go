// Write a program to print whether a number is even or odd, also take input from the user.
package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter a Number:")
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println("Even Digit")
	} else {
		fmt.Println("Odd Digit")
	}

}
