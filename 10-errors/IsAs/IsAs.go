package isas

import (
	"errors"
	"fmt"
	"os"
)

// [lecture] -> To check whether the returned error or any errors that it wraps match a specific sentinel error instance, use errors.Is. It takes in two parameters: the error being checked and the instance you are comparing it against. By default, errors.is uses == to compare each wrapped error with a specified error.

// [lecture] -> errors.As function alows you to check whether a returned error(or any error it wraps) matches a specific type. It takes in two parameters: the error being checked and a pointer to a variable of the type you are comparing it against. If the error matches the type, errors.As assigns the error to the variable and returns true. Otherwise, it returns false.


func fileChecker(name string) error {
  f, err := os.Open(name)
  if err != nil { 
    return fmt.Errorf("in fileChecker: %w", err)
  }
  f.Close()
  return nil
}

type MyErr struct {
  Codes []int
}

func (me MyErr) Error() string {
  return fmt.Sprintf("codes: %v", me.Codes)
}

func fakeMain() {
  err := fileChecker("not_here.txt")
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      fmt.Println("fileChecker failed because the file does not exist")
    }
  }
  
  err := AFuctionThatReturnsAnError()
  var myErr MyErr
  if errors.As(err, &myErr) {
    fmt.Println(MyErr.Codes)
  }
}
