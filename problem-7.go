// Write function that reverses a list, preferably in place.

package main

import (
	"fmt"
)

func main() {
  
	var quantity int
	fmt.Print("How many numbers? ")
	fmt.Scanln(&quantity)
  
	n := make([]float32, quantity)
	for i := 0; i < quantity; i++ {
		fmt.Print("Enter a value: ")
		fmt.Scanln(&n[i])
	}
	for j := 1; j < quantity; j++ {
		if n[0] < n[j] {
			n[0] = n[j]
		}
	}
	fmt.Println("The largest value in your list is", n[0])
}
