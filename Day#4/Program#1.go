// Volume Of Pyramid
package main

import (
	"fmt"
)

var base, width, height, volume float32

func main() {
	fmt.Println("Enter the value of Base:")
	fmt.Scanln(&base)
	fmt.Println("Enter the value of width:")
	fmt.Scanln(&width)
	fmt.Println("Enter the value of height:")
	fmt.Scanln(&height)
	volume = (base * width * height) / 3
	fmt.Println("Volume Of Pyramid", volume)
}
