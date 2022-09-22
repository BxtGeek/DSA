// Input currency in rupees and output in USD.
package main

import "fmt"

var rupee, dollar float32

func main() {
	fmt.Println("Enter Amount in Rupee:")
	fmt.Scanln(&rupee)
	dollar = rupee / 80.22
	fmt.Println("Your Amounr in $ is:", dollar)
}
