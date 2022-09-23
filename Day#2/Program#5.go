// Area Of Parallelogram
package main

import "fmt"

var area, height, base int

func main() {
	fmt.Println("Enter the height of Parallelogram:")
	fmt.Scanln(&height)
	fmt.Println("Enter the base of Parallelogram")
	fmt.Scanln(&base)
	area = base * height
	fmt.Println("Area of Parallelogram:", area)
}
