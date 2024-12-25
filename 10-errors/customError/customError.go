package customerror


// ? Example of a custom error
type Status int

const (
	invalidLogin Status = iota + 1
	NotFound
)

type StatusError struct {
	Status Status
	Msg    string
}

func (se StatusError) Error() string {
	return se.Msg
}

func(se StatusError) Unwrap() error {
  return se.Err
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
  token, err := login(uid, pwd)
  if err != nil {
  return nil, StatusErr{
  Status: InvalidLogin,
  Message: fmt.Sprintf("invalid credentials for user
  %s", uid),
  }
  }
  data, err := getData(token, file)
  if err != nil {
  return nil, StatusErr{
  Status: NotFound,
  Message: fmt.Sprintf("file %s not found", file),
  }
  }
  return data, nil
  }