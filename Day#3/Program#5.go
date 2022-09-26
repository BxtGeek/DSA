// Volume Of Triangular Prism
// https://byjus.com/maths/volume-of-a-prism/
package main

import (
	"fmt"
)

var base, height, length, volume float32

func main() {
	fmt.Println("Enter the value of base:")
	fmt.Scanln(&base)
	fmt.Println("Enter the value of height:")
	fmt.Scanln(&height)
	fmt.Println("Enter the value of length:")
	fmt.Scanln(&length)
	volume = (base * length * height) / 2
	fmt.Println("Volume Of Triangular Prism:", volume)
}
