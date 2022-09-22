// Take 2 numbers as input and print the largest number.
package main

import "fmt"

var a, b int

func main() {
	fmt.Println("Enter First Number:")
	fmt.Scanln(&a)
	fmt.Println("Enter Second Number:")
	fmt.Scanln(&b)
	if a < b {
		fmt.Println("B is Bigest Number", b)
	} else {
		fmt.Println("A is Biggest Number", a)
	}
}
