package apl

import "fmt"

// APIError ...
type APIError struct {
	Data struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	} `json:"error"`
}

// Error ...
func (e APIError) Error() string {
	return fmt.Sprintf("error: %v %v", e.Data.StatusCode, e.Data.Message)
}

// Empty ...
func (e APIError) Empty() bool {
	return false
}

func relevantError(httpError error, apiError *APIError) error {

	if httpError != nil {
		return httpError
	}

	if apiError.Data.Message == "" {
		return nil
	}

	return apiError
}
