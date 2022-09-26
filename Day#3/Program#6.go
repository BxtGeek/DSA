// Volume Of Cylinder
package main

import (
	"fmt"
	"math"
)

var radius, height, volume float32

func main() {
	fmt.Println("Enter the value of radius:")
	fmt.Scanln(&radius)
	fmt.Println("Enter the value of height:")
	fmt.Scanln(&height)
	volume = math.Pi * (radius * radius) * height
	fmt.Println("Volume Of Cylinder", volume)
}
