// Curved Surface Area Of Cylinder
package main

import (
	"fmt"
	"math"
)

var radius, height, area float32

func main() {
	fmt.Println("Enter the value of radius:")
	fmt.Scanln(&radius)
	fmt.Println("Enter the value of height:")
	fmt.Scanln(&height)
	area = 2 * math.Pi * radius * height
	fmt.Println("Curved Surface Area Of Cylinder:", area)
}
