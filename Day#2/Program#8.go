// Perimeter Of Circle
package main

import (
	"fmt"
	"math"
)

var radius, area float32

func main() {
	fmt.Println("Enter the radius of circle:")
	fmt.Scanln(&radius)
	area = 2 * math.Pi * radius
	fmt.Println("Perimeter Of Circle:", area)
}
