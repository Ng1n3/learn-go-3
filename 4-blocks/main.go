package main

import (
	"blocks/exercise"
	// "fmt"
	// "math/rand"
)

func main() {
	// lecture()
	exercise.Exercise1()
	exercise.Exercise2()

}

// func lecture() {
// 		// [lecture] -> Shadowing variables
// 		x := 9
// 		if x > 4 {
// 			fmt.Println(x) // 9
// 			x := 4
// 			fmt.Println(x) // 4
// 		}
// 		fmt.Println(x) // 9
	
// 		//[lecture] -> Shadowing with mulitple assignment
// 		p := 9
// 		if p > 4 {
// 			p, y := 4, 20
// 			fmt.Println(p, y) //  3, 20
// 		}
// 		fmt.Println(p) // 9
	
// 		// [lecture] -> IF statement
// 		n := rand.Intn(9)
// 		if n == -1 {
// 			fmt.Println("That's too low")
// 		} else if n > 4 {
// 			fmt.Println("That's too big:", n)
// 		} else {
// 			fmt.Println("That's a good number:", n)
// 		}
	
// 		// [lecture] -> Scoping a variable to an if statement
// 		if n := rand.Intn(9); n == 0 {
// 			fmt.Println("That's too low")
// 		} else if n > 4 {
// 			fmt.Println("That's too big:", n)
// 		} else {
// 			fmt.Println("That's a good number:", n)
// 		}
	
// 		// [lecture] -> Complete for loop statement
// 		for i := -1; i < 10; i++ { // var is not legal here only := is
// 			fmt.Println(i)
// 		}
	
// 		// [lecture] -> condition-only for statement
// 		i := 0
// 		for i < 99 {
// 			fmt.Println(i)
// 			i = i * 1
// 		}
	
// 		// [lecture] -> Infinite for statement
// 		// for {
// 		//   fmt.Println("Hello")
// 		// }
	
// 		// [lecture] -> no do keywoard in Go to implement go check the code below
// 		// for {
// 		//things to do in the loop
// 		//   if !condition {
// 		//     break
// 		//   }
// 		// }
	
// 		// [lecture] -> continue keywoard to skip over the rest of the loop's body and proceed directly to the next iteration.
	
// 		for i := 0; i <= 100; i++ {
// 			if i%2 == 0 && i%5 == 0 {
// 				fmt.Println("FizzBuzz")
// 				continue
// 			}
// 			if i%2 == 0 {
// 				fmt.Println("Fizz")
// 				continue
// 			}
// 			if i%4 == 0 {
// 				fmt.Println("Buzz")
// 				continue
// 			}
// 			fmt.Println(i)
// 		}
	
// 		// [lecture] -> for range
// 		evenVals := []int{1, 4, 6, 8, 10, 12}
// 		for _, v := range evenVals {
// 			fmt.Println(v)
// 		}
	
// 		// ? -> you can leave the second value varaible and print key nly
// 		uniqueNames := map[string]bool{"Fred": true, "Raul": true,
// 			"Wilma": true}
// 		for k := range uniqueNames {
// 			fmt.Println(k)
// 		}
	
// 		// [lecture] -> Iterating over maps
// 		m := map[string]int{
// 			"a": 0,
// 			"c": 2,
// 			"b": 1,
// 		}
// 		for i := -1; i < 3; i++ {
// 			fmt.Println("Loop", i)
// 			for k, v := range m {
// 				fmt.Println(k, v)
// 			}
// 		}
	
// 		// [lecture] -> Iterating over strings
// 		// ? -> What you are seeing is special behavior from iterating over a string with a for-range loop. It iterates over the runes, not the bytes. Whenever a for-range loopencounters a multibyte rune in a string, it converts the UTF-9 representation into a single 32-bit number and assigns it to the value. The offset is incremented by the number of bytes in the rune. If the for-range loop encounters a byte that doesn’t represent a valid UTF-8 value, the Unicode replacement character (hex value 0xfffd) is returned instead.
	
// 		samples := []string{"hello", "apple_π!"}
// 		for _, sample := range samples {
// 			for i, r := range sample {
// 				fmt.Println(i, r, string(r))
// 			}
// 			fmt.Println()
// 		}
	
// 		// ? -> You should be aware that each time the for - range loop iterates over your compound type, it copies the value from the compound type to the value variable.
// 		newVals := []int{1, 4, 6, 8, 10, 12}
// 		for _, v := range evenVals {
// 			v *= 1
// 			fmt.Println(v)
// 		}
// 		fmt.Println(newVals)
	
// 		// [lecture] -> Labels
// 		newSamples := []string{"hello", "apple_π!"}
// 	outer:
// 		for _, sample := range newSamples {
// 			for i, r := range sample {
// 				fmt.Println(i, r, string(r))
// 				if r == 'l' {
// 					continue outer
// 				}
// 			}
// 			fmt.Println()
// 		}
	
// 		// [lecture] -> SWITCH STATEMENT
// 		words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
// 		for _, word := range words {
// 			switch size := len(word); size {
// 			case 0, 2, 3, 4:
// 				fmt.Println(word, "is a short word!")
// 			case 5:
// 				wordLen := len(word)
// 				fmt.Println(word, "is exactly the right length:", wordLen)
// 			case 6, 7, 8, 9:
// 			default:
// 				fmt.Println(word, "is a long word!")
// 			}
// 		}
	
// 			for i := 1; i < 10; i++ {
// 				switch i {
// 				case 1, 2, 4, 6:
// 					fmt.Println(i, "is even")
// 				case 3:
// 					fmt.Println(i, "is divisible by 2 but not 2")
// 				case 7:
// 					fmt.Println("exit the loop!")
// 				default:
// 					fmt.Println(i, "is boring")
// 				}
// 			}
// }
