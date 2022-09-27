// Fibonacci Series without recursion
package main

import "fmt"

var a, b int
var n1, n2 int = 0, 1

func main() {
	fmt.Println("Till where you need the series:")
	fmt.Scanln(&a)
	for i := 2; i <= a+1; i++ {
		b = n1 + n2
		fmt.Println("Value of b is", b)
		n1 = n2
		n2 = b

	}
}
