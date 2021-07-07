package main

import "fmt"

func main() {

	sum := 0
	
	fmt.Print(Insert value here: )
	
	var n int
	fmt.Scanln(&n)


	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			sum += i
		}
		if i%5 == 0 {
			sum += i
		}
	}

fmt.Println("Answer:", sum)
}
