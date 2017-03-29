package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// StackComponentService is the service object for stack_component operations
type StackComponentService struct {
	sling    *sling.Sling
	endpoint string
}

// NewStackComponentsService return a new StackComponentService
func NewStackComponentsService(sling *sling.Sling) *StackComponentService {
	return &StackComponentService{
		sling:    sling,
		endpoint: "stack_components",
	}
}

// StackComponent represents a stack_component row
type StackComponent struct {
	ID string      `json:"id,omitempty"`
	StackID string `json:"stack_id"`
	StackVersionID string `json:"stack_version_id"`
	ComponentID string `json:"component_id"`
	ComponentVersionID string `json:"component_version_id"`
	ProjectID string `json:"project_id"`
	Name string `json:"name"`
	Services interface{} `json:"services"`
	CreatedByUserID string `json:"created_by_user_id"`
	LastModified string `json:"last_modified"`
	CreatedTime string `json:"created_time"`
}

// StackComponentCreateInput is used for the create of stack_components
type StackComponentCreateInput struct {
	ID string      `json:"id,omitempty"`
}

// StackComponentUpdateInput is used for the update of stack_components
type StackComponentUpdateInput struct {
	ID string      `json:"id,omitempty"`
}

// StackComponentParams filter parameters used in list operations
type StackComponentParams struct {
	Name string `url:"name,omitempty"`
}


// List gets a list of stack_components with optional filter params
func (c *StackComponentService) List(params *StackComponentParams) ([]StackComponent, *http.Response, error) {
	output := &struct{ Data []StackComponent `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a stack_component for the id specified
func (c *StackComponentService) Get(id string) (StackComponent, *http.Response, error) {
	output := &struct{ Data StackComponent `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a stack_component
func (c *StackComponentService) Create(input *StackComponentCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a stack_component for the id specified
func (c *StackComponentService) Update(id string, input *StackComponentUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the stack_component for the id specified
func (c *StackComponentService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
