package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// ProjectService is the service object for project operations
type ProjectService struct {
	sling    *sling.Sling
	endpoint string
}

// NewProjectsService return a new ProjectService
func NewProjectsService(sling *sling.Sling) *ProjectService {
	return &ProjectService{
		sling:    sling,
		endpoint: "projects",
	}
}

// Project represents a project row
type Project struct {
	ID            string            `json:"id,omitempty"`
	Name          string            `json:"name"`
	Settings      interface{}       `json:"settings"`
	Users         interface{}       `json:"users"`
	MetaData      map[string]string `json:"meta_data,omitempty"`
	LastModified  string            `json:"last_modified"`
	CreatedTime   string            `json:"created_time"`
	CreatedByUser `json:"created_by_user"`
}

// ProjectCreateInput is used for the create of projects
type ProjectCreateInput struct {
	ID       string      `json:"id,omitempty"`
	Name     string      `json:"name"`
	Settings interface{} `json:"settings,omitempty"`
	Users    interface{} `json:"users,omitempty"`
}

// ProjectUpdateInput is used for the update of projects
type ProjectUpdateInput struct {
	Name string `json:"name"`
}

// ProjectParams filter parameters used in list operations
type ProjectParams struct {
	Name string `url:"name,omitempty"`
}

// List gets a list of projects with optional filter params
func (c *ProjectService) List(params *ProjectParams) ([]Project, *http.Response, error) {
	output := &struct {
		Data []Project `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a project for the id specified
func (c *ProjectService) Get(id string) (Project, *http.Response, error) {
	output := &struct {
		Data Project `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a project
func (c *ProjectService) Create(input *ProjectCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a project for the id specified
func (c *ProjectService) Update(id string, input *ProjectUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the project for the id specified
func (c *ProjectService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
