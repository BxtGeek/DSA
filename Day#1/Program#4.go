// Write a program to input principal, time, and rate (P, T, R) from the user and find Simple Interest.
package main

import "fmt"

func main() {
	var principal, time, rate, total int
	fmt.Println("Enter the Principal Amount:")
	fmt.Scanln(&principal)
	fmt.Println("Enter the Time:")
	fmt.Scanln(&time)
	fmt.Println("Enter the Rate:")
	fmt.Scanln(&rate)
	total = (principal * time * rate) / 100
	fmt.Printf("Total Amount After %d Year will be:%d", time, total)
}
