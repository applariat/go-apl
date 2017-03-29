package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// StackVersionService is the service object for stack_version operations
type StackVersionService struct {
	sling    *sling.Sling
	endpoint string
}

// NewStackVersionsService return a new StackVersionService
func NewStackVersionsService(sling *sling.Sling) *StackVersionService {
	return &StackVersionService{
		sling:    sling,
		endpoint: "stack_versions",
	}
}

type StackVersionList struct {
	StackID       string `json:"stack_id"`
	StackVersions []StackVersion `json:"stack_versions"`
}

// StackVersion represents a stack_version row
type StackVersion struct {
	ID              string `json:"id"`
	StackID         string `json:"stack_id"`
	ProjectID       string `json:"project_id"`
	Version         int `json:"version"`
	VersionSub      int `json:"version_sub"`
	Releases        interface{} `json:"releases"`
	MetaData        interface{} `json:"meta_data"`
	StackComponents interface{} `json:"stack_components"`
	Stack           interface{} `json:"stack"`
	StackArtifacts  interface{} `json:"stack_artifacts"`

	LastModified    string `json:"last_modified"`
	CreatedTime     string `json:"created_time"`
	CreatedByUser    `json:"created_by_user"`
}

// StackVersionCreateInput is used for the create of stack_versions
type StackVersionCreateInput struct {
	//ID            string      `json:"id,omitempty"`
	StackID       string `json:"stack_id"`
	StackVersions []StackVersion `json:"stack_versions"`

	//ProjectID       string `json:"project_id"`
	//Version         int `json:"version"`
	//VersionSub      int `json:"version_sub"`
	//Releases        interface{} `json:"releases"`
	//MetaData        interface{} `json:"meta_data"`
	//StackComponents interface{} `json:"stack_components"`
	//Stack           interface{} `json:"stack"`
	//StackArtifacts  interface{} `json:"stack_artifacts"`
}

// List gets a list of stack_versions with optional filter params
func (c *StackVersionService) List() ([]StackVersionList, *http.Response, error) {
	output := &struct{ Data []StackVersionList `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, nil, output)
	return output.Data, resp, err
}

// Get get a stack_version for the id specified
func (c *StackVersionService) Get(id string) (StackVersion, *http.Response, error) {
	output := &struct{ Data StackVersion `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a stack_version
func (c *StackVersionService) Create(input *StackVersionCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}


// Delete will delete the stack_version for the id specified
func (c *StackVersionService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
