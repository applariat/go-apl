package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// ComponentService is the service object for component operations
type ComponentService struct {
	sling    *sling.Sling
	endpoint string
}

// NewComponentsService return a new componentService
func NewComponentsService(sling *sling.Sling) *ComponentService {
	return &ComponentService{
		sling:    sling,
		endpoint: "components",
	}
}

// Component represents a component row
type Component struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Category          string `json:"category"`
	LastModified      string `json:"last_modified"`
	CreatedTime       string `json:"created_time"`
	ComponentVersions interface{} `json:"component_versions,omitempty"`
	MetaData `json:"meta_data,omitempty"`
}

// ComponentParams filter parameters
type ComponentParams struct {
	Category string `url:"category,omitempty"`
	Name     string `url:"name,omitempty"`
}

// listComponentsOutput used to wrap the data for API result
type listComponentsOutput struct {
	Data []Component `json:"data"`
}

// getComponentOutput used to wrap the data for API result
type getComponentOutput struct {
	Data Component `json:"data"`
}

// List gets a list of components with optional filter params
func (c *ComponentService) List(params *ComponentParams) ([]Component, *http.Response, error) {
	output := new(listComponentsOutput)
	apiError := new(APIError)

	resp, err := c.sling.New().Get(c.endpoint).QueryStruct(params).Receive(output, apiError)

	return output.Data, resp, relevantError(err, apiError)
}

// Get get a component for the id specified
func (c *ComponentService) Get(id string) (Component, *http.Response, error) {

	// TODO: Fix the component endpoint to return { "data": {...}}
	//output := new(getComponentOutput)
	output := Component{}
	apiError := new(APIError)

	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := c.sling.New().Get(path).Receive(&output, apiError)

	return output, resp, relevantError(err, apiError)
}