package main

import (
	"errors"
	"fmt"
	"functions/exercise"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
  exercise.FIleLen("hello.txt")
	// lecture()
}

func lecture() {
	// closure()

	PassingFunctions()
	// ? -> Return function from a function
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
	// [lecture] -> the main function takes in no parameter and returns no value

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator: ", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}

	// ? -> ANONYMOUS FUNCTIONS
	f := func(j int) {
		fmt.Println("Printing", j, "From inside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	var myFuncVariable func(string) int
	myFuncVariable = f1
	result := myFuncVariable("Hello")
	fmt.Println("f1", result)

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println("f2", result)

	newResult, newRemainder, err := divideAndRemainder(8, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(newRemainder, newResult)

	result, remainder, err := divAndRemainder(4, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// [lecture] -> Return multiple values
func divAndRemainder(num, denum int) (result int, remainder int, err error) {
	if denum == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denum, num % denum, nil
}

// [lecture] -> Blank return
// ! -> Never use this!
func divideAndRemainder(num, denum int) (result int, remainder int, err error) {
	if denum == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = num/denum, num%denum
	return
}

// [lecture] -> Functions are Values

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

// [lecture] ->  SIMPLE CALCULATOR
func add(i, j int) (int, error) { return i + j, nil }
func sub(i, j int) (int, error) { return i - j, nil }
func mul(i, j int) (int, error) { return i * j, nil }
func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// [lecture] -> Function type Declaration
// ? -> Jusat as you can use the type keyword to define a struct, you can use it to define a function type too
type opFucType func(int, int) int

// [lecture] -> Anonymous function
// ? -> you can not only assign funcions to variables, but also define new functions within a function and assign them to varaibles.

// var(
//   add = func(i, j int) int {return i + j}
//   sub = func(i, j int) int {return i - j}
//   mul = func(i, j int) int {return i * j}
//   div = func(i, j int) int {return i / j}
// )
// func anonymous() {

//   x := add(2, 3)
//   fmt.Println(x)
//   changeAdd()
//   y := add(2, 3)
//   fmt.Println(y)
// }

// func changeAdd() {
//   add = func (i, j int) int {return i + j + j}
// }

// [lecture] -> CLOSURES
// ? -> functions declared inside functions are special; they are closures, here the functions have access and can modify variables declared in the outer function.

func closure() {
	a := 20
	f := func() {
		fmt.Println(a)
		a := 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
}

func PassingFunctions() {
	type Person struct {
		Firstname string
		Lastname  string
		Age       int
	}

	people := []Person{
		{"pat", "Patterson", 30},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 20},
	}

	// ? -> Sort by lastname
	sort.Slice(people, func(i, j int) bool {
		return people[i].Lastname < people[j].Lastname
	})
	fmt.Println(people)

	// ? -> sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}

// [lecture] -> Return functions from functions
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

// [lecture] -> DEFER
// ? -> In Go cleanup functions are attached to the function with the defer keywoard

func Defer() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			{
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
		}
	}
}
