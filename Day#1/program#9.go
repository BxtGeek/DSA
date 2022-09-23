// To find out whether the given Number is Palindrome or not.
package main

import "fmt"

var a, b, rem, i int
var r int = 0

func main() {
	fmt.Println("Enter a Number:")
	fmt.Scanln(&a)
	b = rev(a)
	if a == b {
		fmt.Println("Number is Palindrome")
	} else {
		fmt.Println("Number is not Palindrome")
	}
}
func rev(a int) int {
	for i = 0; i < 0; i++ {
		rem = a % 10
		r = r*10 + rem
		a = a / 10
	}
	return (r)
}
