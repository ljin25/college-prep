/* Write a program that asks the user for a number n and gives them the possibility to choose between 
computing the sum and computing the product of 1,…,n. */

package main

import "fmt"

var n int



func main() {
 

	sumN := 1
	productN := 1
	
	for i := 1; i <= n; i++ {
		sumN += i + 1
		return sumN;
	}
	
	 

	for i := 1; i <= n; i++ {
		productN *= (i + 1)
		return productN;
	}


	fmt.Print("Insert value here: ")
	fmt.Scanln(&n)
	
	fmt.Println("Sum (1) or Product (2)?")
	var choose uint8
	fmt.Scanln(&choose)
	
	if choose == 1 {
		fmt.Println("Answer:", sumN)
	} else if choose == 2 {
		fmt.Println("Answer:", productN)
	


}

}
