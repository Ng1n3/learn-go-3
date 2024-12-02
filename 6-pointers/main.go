package main

import (
	// "encoding/json"
	"fmt"

	// "os"
	"pointers/exercises"
)

func main() {
	// lecture()
	// p1 := exercises.MakePerson("Jon", "Doe", 20)
	// fmt.Println("Person(value): ", p1)
	// p2 := exercises.MakePersonPointer("Jane", "Smith", 20)
	// fmt.Println("Person (pointer): ", p2)

  updateSlice :=[]string{"apple", "mango", "banana"}
  fmt.Println("before updateSlice: ", updateSlice)
  exercises.UpdateSlice(updateSlice, "juice")
  fmt.Println("After updateSlice: ", updateSlice)

  fmt.Println()

  growSlice := []string{"apple", "mango", "banana"}
  fmt.Println("before growSlice: ", growSlice)
  exercises.GrowSlice(growSlice, "juice")
  fmt.Println("After growSlice: ", growSlice)


}

// func lecture() {

// 	//[lecture] -> POINTER
// 	// ? -> A pointer is a variable that contains the address where another variable is stored.

// 	x := 10
// 	pointerToX := &x
// 	fmt.Println(pointerToX)  // prints a memory address
// 	fmt.Println(*pointerToX) // prints 10
// 	z := 5 + *pointerToX
// 	fmt.Println(z) // prints 15

// 	// ! -> Your program will panic if you dereference a nil pointer

// 	var p *int
// 	fmt.Println(p == nil) // panics true
// 	// ! fmt.Println(*p) // panics

// 	// ? You can't use an & before a primitive literal(numbers, booleans, ans strings) or a constant because they don't have memory addresses; they exist only at compile time. When you need a pointer to a primitive type, declare a variable and point to it:

// 	type person struct {
// 		Firstname  string
// 		MiddleName *string
// 		LastName   string
// 	}

// 	k := person{
// 		Firstname:  "Jonny",
// 		MiddleName: makePointer("Perry"),
// 		// ! MiddleName: "Perry", this line won't compile because you can't assign a literal directly to the field.
// 		LastName: "Simpsons",
// 	}

// 	fmt.Println(k)

// 	var f *int // f is nil
// 	failedUpdate(f)
// 	fmt.Println(f) // prints nil

// 	// --------------------------

// 	y := 10
// 	failedUpdate2(&y)
// 	fmt.Println(y) // prints 20
// 	update(&y)
// 	fmt.Println(y) // prints 20

	// m := struct {
	// 	Name string `json:name`
	// 	Age  int    `json:age`
	// }{}
	// err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &m)

// }

// ? -> this works because passing in the value into this function gives it an address in memory since a function is just a variable.
// func makePointer[T any](t T) *T {
// 	return &t
// }

// ? -> GO developers use pointers to indidcate that a parameter is mutable.

// ? -> You start with a nil variable f in main. When you call failedUpdate, you copy the value of f, which is nil, into the parameter named g. This means that g is also set to nil. You then declare a new variable x within failedUpdate with the value 10. Next, you change g in failedUpdate to point to x. This does not change the f in main, and when you exit failedUpdate and return to main, f is still nil.
// func failedUpdate(g *int) {
// 	x := 10
// 	g = &x
// }

// ? ->

// func failedUpdate2(px *int) {
// 	x2 := 20
// 	px = &x2
// }

// func update(px *int) {
// 	*px = 20
// }

// ? -> the only time you should use pointer parameter to modify a variable is when thefunction expects an interface. You see this pattern when working with JSON

// [lecture] -> Slices as Buffers
