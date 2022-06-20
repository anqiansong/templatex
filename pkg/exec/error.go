package exec

import "fmt"

// ExitCodeError is an error type that represents an exit code.
type ExitCodeError struct {
	code int
	err  error
}

// Error returns the error message.
func (e ExitCodeError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("exitCode: %d, err: %s", e.code, e.err.Error())
	}
	return fmt.Sprintf("exitCode: %d", e.code)
}

// Code returns the exit code.
func (e ExitCodeError) Code() int {
	return e.code
}
