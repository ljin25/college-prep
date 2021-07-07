package main

import "fmt"

func main() {

	sum := 0
	
	fmt.Print(Insert value here: )
	
	var n uint8
	fmt.Scanln(&n)


	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			sum += i
		}
		if i%5 == 0 {
			sum += i
		}
	}

fmt.Print
fmt.Println("Answer:", sum)
}
