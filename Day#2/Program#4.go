// Area Of Isosceles Triangle
package main

import "fmt"

var base, height, area int

func main() {
	fmt.Println("Enter the Base of Isosceles Triangle:")
	fmt.Scanln(&base)
	fmt.Println("Enter the height of Isosceles Triangle:")
	fmt.Scanln(&height)
	area = (base * height) / 2
	fmt.Println("Area of Isosceles Triangle:", area)
}
