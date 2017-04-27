package app

import "fmt"

// CLIError ...
type CLIError struct {
	Message string `json:"message"`
}

// Error ...
func (e CLIError) Error() string {
	return fmt.Sprintf("error: %v", e.Message)
}

// Empty ...
func (e CLIError) Empty() bool {
	return false
}

// NewCLIError creates an error from a string
func NewCLIError(msg string) CLIError {
	return CLIError{Message: msg}
}
