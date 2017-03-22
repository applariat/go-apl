package apl

import "fmt"

// APLError ...
type APLError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
} // `json:"error"`

// APIError ...
type APIError struct {
	APLError `json:"error"`
}

// Error ...
func (e APIError) Error() string {
	return fmt.Sprintf("error: %v %v", e.APLError.StatusCode, e.APLError.Message)
}

// Empty ...
func (e APIError) Empty() bool {
	return false
}

// Error ...
func (e APLError) Error() string {
	return fmt.Sprintf("error: %v %v", e.StatusCode, e.Message)
}

// Empty ...
func (e APLError) Empty() bool {
	return false
}


// from the two errors, return one APLError
func relevantError(httpError error, apiError *APIError) error {

	if httpError != nil {
		return APLError{
			StatusCode: 400,
			Message: httpError.Error(),
		}
	}

	if apiError.APLError.Message == "" {
		return nil
	}

	return apiError.APLError
}
