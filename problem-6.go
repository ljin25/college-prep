// Write a function that returns the largest element in a list.

package main

import (
	"fmt"
)

func main() {

	var quantity int
	fmt.Print("How many numbers? ")
	fmt.Scanln(&quantity)
	
  n := make([]int, quantity)

  for i := 0; i < quantity; i++ {
		fmt.Print("Enter a value: ")
		fmt.Scanln(&n[i])
	}
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	fmt.Println(n)
}
