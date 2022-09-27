// Total Surface Area Of Cube
// https://www.cuemath.com/measurement/surface-area-of-cube/
package main

import (
	"fmt"
)

var sides, area int

func main() {
	fmt.Println("Enter the value of sides:")
	fmt.Scanln(&sides)
	area = 6 * (sides * sides)
	fmt.Println("Total Surface Area Of Cube", area)
}
