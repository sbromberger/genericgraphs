package errors

import "fmt"

// EarlyTerminationError is returned when a visitor function returns `false`.
type EarlyTerminationError string

func (e EarlyTerminationError) Error() string {
	return fmt.Sprintf("EarlyTerminationError: %s", string(e))
}

// InvalidVertexError is returned when a referenced vertex is not valid for a
// given graph.
type InvalidVertexError string

func (e InvalidVertexError) Error() string {
	return fmt.Sprintf("InvalidVertexError: %s", string(e))
}
