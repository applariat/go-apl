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
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Category          string      `json:"category"`
	LastModified      string      `json:"last_modified"`
	CreatedTime       string      `json:"created_time"`
	ComponentVersions interface{} `json:"component_versions,omitempty"`
	MetaData          `json:"meta_data,omitempty"`
}

// ComponentParams filter parameters used in list operations
type ComponentParams struct {
	Category string `url:"category,omitempty"`
	Name     string `url:"name,omitempty"`
}

// List gets a list of components with optional filter params
func (c *ComponentService) List(params *ComponentParams) ([]Component, *http.Response, error) {
	output := &struct {
		Data []Component `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a component for the id specified
func (c *ComponentService) Get(id string) (Component, *http.Response, error) {

	// wrap output data
	// TODO: Fix the component endpoint to return { "data": {...}}
	output := &struct {
		Data Component `json:"data"`
	}{}
	//output := Component{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
	//return output, resp, err

}
