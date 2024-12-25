package main

import (
	"errors"
	panicrecover "errors/panicRecover"
	"fmt"

	"github.com/pkg/errors"
)

type error interface {
	Error() string
}

func Lecture() {
	//[lecture] -> Error is a built-in interface that defines a single method
	//[lecture] -> The reason you return nill from a function is to indicate that no error occurred is that nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.
	//[Lecture] -> You can call errors using error.New() or fmt.Errorf()

	//[lecture] -> Since error is an interface, you can define your own errors that include additinal information for logging or error handling.

	//[Lecture] -> If ou need to merge multiple errors into a single error. That's what the errors.Join function is for:

}

func main() {

	//? Sentinel error solution
	ids := []int{1, 0, -3}

	for _, id := range ids {
		err := panicrecover.Solution2(id)
		if errors.Is(err, panicrecover.ErrInvalidID) {
			fmt.Println("Invalid ID found")
		} else {
			fmt.Printf("Id %d is valid.\n", id)
		}
	}

	//? Custom error solution
	name := ""
	err := panicrecover.Solution3(name)
	if err != nil {
		var emptyFieldErr *panicrecover.EmptyFieldError
		if errors.As(err, &emptyFieldErr) {
			fmt.Printf("Error: %v\n", emptyFieldErr)
			fmt.Printf("Field causing error: %s\n", emptyFieldErr.FieldName)
		} else {
			fmt.Println("No empty field error found")
		}
	} else {
		fmt.Println("No error found")
	}
}


