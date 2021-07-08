package main

import "fmt"

func main() {
   
var i, n int


for i = n; i <= (n+80); i++ {
	if (i%4) != 0 {
		break;
	}
	
	fmt.Print("Enter the current year: ")
	n = 2021
	fmt.Println("Answer:", i)
}

  
}
