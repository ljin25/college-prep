// Write a program that prints a multiplication table for numbers up to 12.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Insert value here: ")
	fmt.Scanln(&n)
	i := 1

	for {
		if i > 12 {
		break;
		}

		fmt.Println(n, "*", i, "=", n*i)
		i++

	}
}
