// Perimeter Of Rectangle
package main

import "fmt"

var length, perimeter, width int

func main() {
	fmt.Println("Enter the length of Rectangle:")
	fmt.Scanln(&length)
	fmt.Println("Enter the width of Rectangle:")
	fmt.Scanln(&width)
	perimeter = 2 * (length + width)
	fmt.Println("Perimeter Of Rectangle is:", perimeter)
}
