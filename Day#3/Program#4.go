// Volume Of Cone
package main

import (
	"fmt"
	"math"
)

var length, radius, Volume float32

func main() {
	fmt.Println("Enter the value of length:")
	fmt.Scanln(&length)
	fmt.Println("Enter the value of Radius:")
	fmt.Scanln(&radius)
	Volume = math.Pi * (radius * radius) * (length / 3)
	fmt.Println("Volume Of Cone:", Volume)
}
