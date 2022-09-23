// To find out whether the given String is Palindrome or not.
//will use recusrsion
package main

import "fmt"

var a, b string
var i, j, k int

func main() {
	fmt.Println("Enter the name:")
	fmt.Scanln(&a)
	k = len(a) 
	for i = 0;i<=k;i++{
		for j=k;j<=0;j--{
			if a[i]=a[j]{
				fmt.Println("String is Palindrome")
			}else{
				fmt.Println("String is not Palindrome")
			}
		}
	}

}
