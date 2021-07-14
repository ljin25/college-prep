// Write a function that returns the elements on odd positions in a list.

package main

import (
	"fmt"
)

func main() {

	list := [7]int{0, 1, 2, 3, 4, 5, 6}
	
	fmt.Print("The elements on odd positions of this list are: ")

	for i := 0; i < len(list); i++ {
		if i%2 != 0 {
		fmt.Print(list[i], " ")
		}
	
	}

}
