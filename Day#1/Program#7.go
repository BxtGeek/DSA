// To calculate Fibonacci Series up to n numbers
package main

import "fmt"

var steps, i int
var a, b, c int

func main() {
	fmt.Println("How many Number your want:")
	fmt.Scanln(&steps)
	a = 0
	b = 1
	for i = 0; i <= steps; i++ {
		c = a + b
		a = b
		b = c
	}
	fmt.Printf("The %dth number of series is:%d", steps, b)
}
