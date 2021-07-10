// Write a program that prints the next 20 leap years.

package main

import "fmt"

func leapYear(y int) bool {
	return y%4 == 0 || y%4 == 0 && y%100 != 0
}

func main() {
   
	var currentYear int
	
	fmt.Print("Enter the current year: ")
	fnt.Scanln(&currentYear)

	for i := n; i <= (n+80); i++ {
		if leapYear(i) {
			fmt.Println(i)
		}	
	}
}
