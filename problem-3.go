/*
Write a program that asks the user for a number n and gives them the possibility to choose between 
computing the sum and computing the product of 1,…,n.
*/

package main

import "fmt"

func sum(n int) uint16 {

	var resultSum uint16 = 0
	for i := 1; i <= n; i++ {
		resultSum += uint16(i)
	}
	return resultSum
}

func product(n int) uint64 {

	var resultProduct uint64 = 1
	if n < 0 {

	} else {
		for i := 1; i <= n; i++ {
			resultProduct *= uint64(i)
		}
	}
	return resultProduct
}

func main() {

	var n int
	fmt.Print("Insert value here: ")
	fmt.Scanln(&n)
	
	var choose uint8
	fmt.Println("Sum (1) or Product (2)?")
	fmt.Scanln(&choose)
	
	if choose == 1 {
		fmt.Printf("Sum of all numbers from 1 to %d is %d", n, sum(n))
	} else if choose == 2 {
		fmt.Printf("Product of all numbers from 1 to %d (factorial) is %d", n, product(n))
	} else {
		fmt.Println("Please choose 1 for the sum or 2 for the product.")
	}
	
}
