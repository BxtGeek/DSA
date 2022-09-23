// Perimeter Of Equilateral Triangle
package main

import "fmt"

var a, area int

func main() {
	fmt.Println("Enter the Value of a:")
	fmt.Scanln(&a)
	area = 3 * a
	fmt.Println("Perimeter Of Equilateral Triangle:", area)
}
