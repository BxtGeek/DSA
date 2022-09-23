// Perimeter Of Parallelogram
package main

import "fmt"

var base, side, Perimeter int

func main() {
	fmt.Println("Enter the Value of Side:")
	fmt.Scanln(&side)
	fmt.Println("Enter the Value of Base:")
	fmt.Scanln(&base)
	Perimeter = 2 * (side + base)
	fmt.Println("Perimeter Of Parallelogram:", Perimeter)
}
