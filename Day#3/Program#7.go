// Volume Of Sphere
package main

import (
	"fmt"
	"math"
)

var radius, volume float32

func main() {
	fmt.Println("Enter the valud of raidus")
	fmt.Scanln(&radius)
	volume = (4 * (math.Pi * radius * radius * radius)) / 3
	fmt.Println("The Volume Of Sphere is:", volume)
}
