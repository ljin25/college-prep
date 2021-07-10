/*
Write a program that asks the user for a number n and prints the sum of the numbers 1 to n. 
Modify the previous program such that only multiples of three or five are considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n=17.
*/

package main

import "fmt"

func main() {

//part 1: sum of all numbers 1 to n
	fmt.Print("Insert value here: ")

	var n uint8
	fmt.Scanln(&n)

	var sum = (n * (n + 1)) / 2
	fmt.Println("Answer:", sum)

//part 2: sum of all multiples of 3 & 5 to n (x in this case)
	threesAndFives := 0

	fmt.Print("Insert value here: ")

	var x int
	fmt.Scanln(&x)

	for i := 1; i <= x; i++ {
		if i%3 == 0 {
			threesAndFives += i
		}
		if i%5 == 0 {
			threesAndFives += i
		}
	}
	
	fmt.Println("Answer:", threesAndFives)
}
