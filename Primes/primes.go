package main

import "fmt"

func main() {
  var i, j, numRange int
  fmt.Scan(&numRange)

  for i=2; i < numRange; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break; // if factor found, not prime
         }
      }
      if(j > (i/j)) {
         fmt.Println(i, " is a prime number");
      }
   }

}
