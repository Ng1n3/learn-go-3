package wraperror

import "fmt"

  //[Lecture] -> When you preserve an error while adding information, it is called WRAPPING an error. When yyou have a series of wrapped errors, it is called an error tree. The convention is to write :%w at the end of the error format string and to make the error to be wrapped the last parameter passed to fmt.Errorf. Not all errors need to be wrapped. A library can return an error that means processing cannot continue, but the error message contains implementation details that aren’t needed in other parts of your program. In this situation, it is perfectly acceptable to create a brand-new error and return that instead. Understand the situation and determine what needs to be returned.

type Status int

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}
func (se StatusErr) Error() string {
  return se.Message
}

func (se StatusErr) Unwrap() error {
  return se.Err
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
  token, err := login(uid, pwd)
  if err != nil {
    return nil, StatusErr{
      Status: invalidLogin,
      Message: fmt.Sprintf("invalid credentials for user %s", uid),
    }
  }
  data, err := getData(token, file)
  if err != nil {
    return nil, StatusErr{
      Status: NotFound,
      Message: fmt.Sprintf("file %s not found", file),
      Err: err,
    }
  }
  return data, nil
}
