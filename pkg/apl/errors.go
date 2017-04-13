package apl

import "fmt"

// APIError ...
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
} // `json:"error"`

// WrappedAPIError ...
type WrappedAPIError struct {
	APIError `json:"error"`
}

// Error ...
func (e WrappedAPIError) Error() string {
	return fmt.Sprintf("error: %v %v", e.APIError.StatusCode, e.APIError.Message)
}

// Empty ...
func (e WrappedAPIError) Empty() bool {
	return false
}

// Error ...
func (e APIError) Error() string {
	return fmt.Sprintf("error: %v %v", e.StatusCode, e.Message)
}

// Empty ...
func (e APIError) Empty() bool {
	return false
}

// from the two errors, return one APLError
func relevantError(httpError error, wrappedError *WrappedAPIError) error {

	if httpError != nil {
		return APIError{
			StatusCode: 400,
			Message:    httpError.Error(),
		}
	}

	if wrappedError.APIError.Message == "" {
		return nil
	}

	return wrappedError.APIError
}
