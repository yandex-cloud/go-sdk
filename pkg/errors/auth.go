package errors

import "fmt"

type AuthError struct {
	Err error
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("authentication error: %v", e.Err)
}
