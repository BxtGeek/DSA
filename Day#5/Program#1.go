// Input a number and print all the factors of that number (use loops).
package main

import "fmt"

var a, i int

func main() {
	fmt.Println("Enter a Number:")
	fmt.Scanln(&a)
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Println(i, "is a factor of", a)
		}
	}
}
