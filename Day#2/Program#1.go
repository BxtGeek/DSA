// Area Of Circle Java Program
package main

import "fmt"

var area, radius float32

func main() {
	fmt.Println("Enter the Radius of circle:")
	fmt.Scanln(&radius)
	area = 3.14 * (radius * radius)
	fmt.Println("Area of Circle is:", area)
}
