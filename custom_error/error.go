package custom_error

import "fmt"

type CreateError struct {
	Msg string
}

func (ce *CreateError) Error() string {
	return fmt.Sprintf("err %s", ce.Msg)
}
