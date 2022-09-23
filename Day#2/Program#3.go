// Area Of Rectangle Program
package main

import "fmt"

var area, width, length int

func main() {
	fmt.Println("Enter the length of Rectangle:")
	fmt.Scanln(&length)
	fmt.Println("Enter the width of Rectangle:")
	fmt.Scanln(&width)
	area = width * length
	fmt.Println("Area of Rectangle :", area)
}
