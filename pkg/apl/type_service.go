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
	ID    string      `json:"id"`
	Types interface{} `json:"types"`
}

// List gets a list of types with optional filter params
func (c *TypeService) List() ([]Type, *http.Response, error) {
	output := &struct {
		Data []Type `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, nil, output)
	return output.Data, resp, err
}

// Get get a type for the id specified
func (c *TypeService) Get(id string) (Type, *http.Response, error) {
	output := &struct {
		Data Type `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}
