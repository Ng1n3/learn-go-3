package exercises

import "fmt"

/*
1. Create a struct named Person with three fields: FirstName and LastName of
type string and Age of type int. Write a function called MakePerson that
takes in firstName, lastName, and age and returns a Person. Write a second
function MakePersonPointer that takes in firstName, lastName, and age and
returns a *Person. Call both from main. Compile your program with go build
-gcflags="-m". This both compiles your code and prints out which values
escape to the heap. Are you surprised about what escapes?
2. Write two functions. The UpdateSlice function takes in a []string and a
string. It sets the last position in the passed-in slice to the passed-in string.
At the end of UpdateSlice, print the slice after making the change. The
GrowSlice function also takes in a []string and a string. It appends the
string onto the slice. At the end of GrowSlice, print the slice after making the
change. Call these functions from main. Print out the slice before each
function is called and after each function is called. Do you understand why
some changes are visible in main and why some changes are not?
3. Write a program that builds a []Person with 10,000,000 entries (they could
all be the same names and ages). See how long it takes to run. Change the
value of GOGC and see how that affects the time it takes for the program to
complete. Set the environment variable GODEBUG=gctrace=1 to see when
garbage collections happen and see how changing GOGC changes the number
of garbage collections. What happens if you create the slice with a capacity of
10,000,000?
*/

type Person struct {
	FirstName string
	Lastname  string
	Age       int
}

func MakePerson(firstname, lastname string, age int) Person {
	return Person{
		FirstName: firstname,
		Lastname:  lastname,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		Lastname:  lastName,
		Age:       age,
	}
}

func UpdateSlice(s []string, newValue string) {
	if len(s) > 0 {
		s[len(s)-1] = newValue
	}
	fmt.Println("inside updateSLice: ", s)
}

func GrowSlice(s []string, newValue string) []string {
	s = append(s, newValue)
	fmt.Println("Inside GrowSice: ", s)
	return s
}
