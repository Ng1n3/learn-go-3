package exercise

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1. Write a for loop that puts 100 random numbers between 0 and 100 into an
int slice.
2. Loop over the slice you created in exercise 1. For each value in the slice, apply
the following rules:
a. If the value is divisible by 2, print “Two!”
b. If the value is divisible by 3, print “Three!”
c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything
else.
d. Otherwise, print “Never mind”.
3. Start a new program. In main, declare an int variable called total. Write a
for loop that uses a variable named i to iterate from 0 (inclusive) to 10
(exclusive). The body of the for loop should be as follows:
total := total + i
fmt.Println(total)
After the for loop, print out the value of total. What is printed out? What is
the likely bug in this code?
*/



func Exercise1() {
  //define a seed using current time
  seed := time.Now().UnixNano()

  //generate a random number using seed
  r := rand.New(rand.NewSource(seed))
  num := make([]int, 100)
  for i := 0; i < 100; i++ {
    num[i] = r.Intn(101)
    switch {
    case num[i] % 2 == 0:
      fmt.Println("Two")
    case num[i] % 3 == 0:
      fmt.Println("Three!")
    case num[i] % 2 == 0 && num[i] % 3 == 0: {
      fmt.Println("Six!")
    }
  default:
    fmt.Println("Never mind.")
    }
  }
  // fmt.Println(num)
}

func Exercise2() {
  var total int

  i := 0
  for i <= 10 {
    total := total + i
    fmt.Println(total)
  }

}