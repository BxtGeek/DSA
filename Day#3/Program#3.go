// Perimeter Of Rhombus
package main

import "fmt"

var perimeter, side int

func main() {
	fmt.Println("Enter the side:")
	fmt.Scanln(&side)
	perimeter = 4 * side
	fmt.Println("Perimeter Of Rhombus:", perimeter)
}
