// To find Armstrong Number between two given number.
package main

import "fmt"

var a, i, sum, temp int

func main() {
	fmt.Println("Enter a Number:")
	fmt.Scanln(&a)
	temp = a
	for a > 0 {
		i = a % 10
		sum = sum + (i * i * i)
		a = a / 10
	}
	if temp == sum {
		fmt.Println("Number is a armstrong number")
	} else {
		fmt.Println("Number is not a armstrong number")
	}
}
