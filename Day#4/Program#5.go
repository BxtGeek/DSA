// Find Sum of Digits in a Number
package main

import "fmt"

var dignum, digiSum, digiReminder int

func main() {
	fmt.Println("Enter the value of a:")
	fmt.Scanln(&dignum)
	for digiSum = 0; dignum > 0; dignum = dignum / 10 {
		digiReminder = dignum % 10
		digiSum = digiSum + digiReminder
	}
	fmt.Println("The Sum of Digits in this Number = ", digiSum)
}
