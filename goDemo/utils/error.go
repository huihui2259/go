package utils

import "fmt"

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("my error: %s", e.Message)
}
