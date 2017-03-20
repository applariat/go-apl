package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// TypeService is the service object for type operations
type TypeService struct {
	sling    *sling.Sling
	endpoint string
}

// NewTypesService return a new typeService
func NewTypesService(sling *sling.Sling) *TypeService {
	return &TypeService{
		sling:    sling,
		endpoint: "types",
	}
}

// Type represents a type row
type Type struct {
	ID    string   `json:"id"`
	Types []string `json:"types"`
}

// listTypesOutput used to wrap the data for API result
type listTypesOutput struct {
	Data []Type `json:"data"`
}

// getTypeOutput used to wrap the data for API result
type getTypeOutput struct {
	Data Type `json:"data"`
}

// List gets a list of types with optional filter params
func (c *TypeService) List() (*[]Type, *http.Response, error) {
	output := new(listTypesOutput)
	apiError := new(APIError)

	resp, err := c.sling.New().Get(c.endpoint).Receive(output, apiError)

	return &output.Data, resp, relevantError(err, apiError)
}

// Get get a type for the id specified
func (c *TypeService) Get(id string) (*Type, *http.Response, error) {
	output := new(getTypeOutput)
	apiError := new(APIError)

	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := c.sling.New().Get(path).Receive(output, apiError)

	return &output.Data, resp, relevantError(err, apiError)
}
