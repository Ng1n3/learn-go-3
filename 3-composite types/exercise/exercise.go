package exercise

import "fmt"

/*
1. Write a program that defines a variable named greetings of type slice of
strings with the following values: "Hello", "Hola", "नमस् का र", "こんにちは",
and "Привіт". Create a subslice containing the first two values; a second
subslice with the second, third, and fourth values; and a third subslice with
the fourth and fifth values. Print out all four slices.
2. Write a program that defines a string variable called message with the value
"Hi
and
" and prints the fourth rune in it as a character, not a number.
3. Write a program that defines a struct called Employee with three fields:
firstName, lastName, and id. The first two fields are of type string, and the
last field (id) is of type int. Create three instances of this struct using
whatever values you’d like. Initialize the first one using the struct literal style
without names, the second using the struct literal style with names, and the
third with a var declaration. Use dot notation to populate the fields in the
third struct. Print out all three structs.
*/

func Exercise3() {
	// Third exercise
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	john := Employee{
		"John",
		"Fagbenga",
		25,
	}

	thomas := Employee{
		firstName: "Thomas",
		lastName:  "Federighi",
		id:        11,
	}

	var ederson struct {
		firstName string
		lastName  string
		id        int
	}

	ederson = thomas
	ederson.firstName = "Ederson"
	ederson.lastName = "Frenkie"
	ederson.id = 10

	fmt.Println("John Struct: ", john)
  fmt.Println("Thomas struct: ", thomas)
  fmt.Println("Ederson struct",  ederson)
}

func Exercise2() {
  var message string = "Hi 👩‍🦰 and 🧔‍♂️"
  var x []rune  = []rune(message)
  fmt.Printf("%c\n", x[3])
}

func Exercise1() {
  greetings := []string{"Hello", "Hola", "नमस् का र", "こんにちは", "Привіт"}
  greeting1 := greetings[:2]
  greeting2 := greetings[1: 4]
  greetings3 := greetings[3: 5]

  fmt.Println("greeting: ", greetings)
  fmt.Println("greeting1: ", greeting1)
  fmt.Println("greeting2: ", greeting2)
  fmt.Println("greeting3: ", greetings3)

}


