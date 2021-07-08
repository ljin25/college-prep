// Write a program that prints all prime numbers up to 2,000,000,000

package main

import "fmt"

func main() {
   
   var i, j int

   for i = 2; i < 2e9; i++ {
      for j = 2; j <= (i/j); j++ {
         if (i%j) == 0 {
            break;
         }
      }
      if j > (i/j) {
         fmt.Printf("%d\n", i);
      }
   }  
}
