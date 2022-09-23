// Area Of Equilateral Triangle
package main

import (
	"fmt"
	"math"
)

var a, area float64

func main() {
	fmt.Println("Enter the Side:")
	fmt.Scanln(&a)
	area = (math.Sqrt(3) / 4) * (a * a)
	fmt.Println("Area of Equilateral Triangle:", area)

}
