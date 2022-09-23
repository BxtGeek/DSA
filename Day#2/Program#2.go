// Area Of Triangle
package main

import "fmt"

var height, base, area int

func main() {
	fmt.Println("Enter the height of Triange:")
	fmt.Scanln(&height)
	fmt.Println("Enter the base of Triange:")
	fmt.Scanln(&base)
	area = (height * base) / 2
	fmt.Println("Area Of Triangle is:", area)
}
