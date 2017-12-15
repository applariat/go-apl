package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// StackService is the service object for stack operations
type StackService struct {
	sling    *sling.Sling
	endpoint string
}

// NewStacksService return a new StackService
func NewStacksService(sling *sling.Sling) *StackService {
	return &StackService{
		sling:    sling,
		endpoint: "stacks",
	}
}

// Stack represents a stack row
type Stack struct {
	ID             string      `json:"id,omitempty"`
	Name           string      `json:"name"`
	//VersionNumber  int         `json:"version_number,omitempty"`
	ReleaseNumber  int         `json:"release_number,omitempty"`
	Project        interface{} `json:"project"`
	StackVersions  interface{} `json:"stack_versions"`
	StackArtifacts interface{} `json:"stack_artifacts"`

	CreatedByUser `json:"created_by_user"`
	LastModified  string `json:"last_modified"`
	CreatedTime   string `json:"created_time"`
}

// StackCreateInput is used for the create of stacks
type StackCreateInput struct {
	ID         string      `json:"id,omitempty"`
	Name       string      `json:"name"`
	MetaData   interface{} `json:"meta_data"`
	ProjectID  string      `json:"project_id"`
	UseVersion int         `json:"use_version"`
	Components interface{} `json:"components"`
}

// StackUpdateInput is used for the update of stacks
type StackUpdateInput struct {
	Name           string      `json:"name"`
	VersionNumber  int         `json:"version_number,omitempty"`
	ReleaseNumber  int         `json:"release_number,omitempty"`
	Project        interface{} `json:"project,omitempty"`
	StackVersions  interface{} `json:"stack_versions,omitempty"`
	StackArtifacts interface{} `json:"stack_artifacts,omitempty"`
}

// StackParams filter parameters used in list operations
type StackParams struct {
	Name          string `url:"name,omitempty"`
	VersionNumber int    `url:"version_number,omitempty"`
	ReleaseNumber int    `url:"release_number,omitempty"`
}

// List gets a list of stacks with optional filter params
func (c *StackService) List(params *StackParams) ([]Stack, *http.Response, error) {
	output := &struct {
		Data []Stack `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a stack for the id specified
func (c *StackService) Get(id string) (Stack, *http.Response, error) {
	output := &struct {
		Data Stack `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a stack
func (c *StackService) Create(input *StackCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a stack for the id specified
func (c *StackService) Update(id string, input *StackUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the stack for the id specified
func (c *StackService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
