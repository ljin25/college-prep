// Write a function that checks whether an element occurs in a list.

package main

import (
	"fmt"
)

func main() {

	list := [7]int{0, 1, 2, 3, 4, 5, 6}

	for i := 0; i < len(list); i++ {
		if i == 0 { // As long as i == any value in 'list', the below line will print, otherwise nothing
			fmt.Println("Element occurs in list")
		}
	}

}
