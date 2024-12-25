package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
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


