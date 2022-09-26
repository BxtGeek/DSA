// Perimeter Of Square
package main

import "fmt"

var perimeter, side int

func main() {
	fmt.Println("Enter the Side:")
	fmt.Scanln(&side)
	perimeter = 4 * side
	fmt.Println("Perimeter Of Square:", perimeter)
}
