// Area Of Rhombus
package main

import "fmt"

var Diagonal1, Diagonal2, area int

func main() {
	fmt.Println("Enter the Diagonal1 of Rhombus:")
	fmt.Scanln(&Diagonal1)
	fmt.Println("Enter the Diagonal2 of Rhombus:")
	fmt.Scanln(&Diagonal2)
	area = (Diagonal1 * Diagonal2) / 2
	fmt.Println("Area of Rhombus:", area)
}
