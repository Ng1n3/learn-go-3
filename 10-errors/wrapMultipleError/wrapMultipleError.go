package wrapmultipleerror

//? Example of wrapping Multiple errors
type Person struct {
  FirstName string
  LastName string
  Age int
}

func ValidatePerson(p Person) error {
  var errs []error
  if(len(p.FirstName) == 0) {
    errs = append(errs, errors.New("FirstName cannot be empty"))
  }
  if len(p.LastName) == 0 {
    errs = append(errs, errors.New("LastName cannot be empty"))
  }
  if len(errs) > 0 {
    return errors.Join(errs...)
  }
  return nil
}

func MergeErrors() {
  err1 := errors.New("first error")
  err2 := errors.New("second eror")
  err3 := errors.New("third error")
  err := fmt.Errorf("first: %w, seocnd: %w, third: %w", err1, err2, err3)
}

// ? Implelent your own error type that supports multiple wrapped errors.

type MyError struct {
  Code int
  Errors []error
}

type(m MyError) Error() string {
  return errors.Join(m.Errors..).Error()
}

func (m MyError) Unwrap() []error {
  return m.Errors
}

//? handle erros that may wrap zero, one or multiple errors

var err error
er = funcThatReturnsAnEror()
switch err := err.(type) {
case interface {Unwrap() error}:
  //handle single error
  innerErr := err.Unwrap()
  //process innerErr
case interface {Unwrap() []error}:
  //handle multiple errors
  innerErrs := err.Unwrap()
  for _, innerErr := range innerErrs {
    //process innerErr
  }
default:
  //handle no wrapped error
}