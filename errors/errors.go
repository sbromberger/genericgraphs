package errors

import "fmt"

type EarlyTerminationError string

func (e EarlyTerminationError) Error() string {
	return fmt.Sprintf("EarlyTerminationError: %s", string(e))
}

type InvalidVertexError string

func (e InvalidVertexError) Error() string {
	return fmt.Sprintf("InvalidVertexError: %s", string(e))
}
