package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// ProjectRoleService is the service object for project_role operations
type ProjectRoleService struct {
	sling    *sling.Sling
	endpoint string
}

// NewProjectRolesService return a new ProjectRoleService
func NewProjectRolesService(sling *sling.Sling) *ProjectRoleService {
	return &ProjectRoleService{
		sling:    sling,
		endpoint: "project_roles",
	}
}

// ProjectRole represents a project_role row
type ProjectRole struct {
	ID              string `json:"id,omitempty"`
	UserID          string `json:"user_id"`
	ProjectID       string `json:"project_id"`
	LastModified    string `json:"last_modified"`
	CreatedTime     string `json:"created_time"`
	CreatedByUserID string `json:"created_by_user_id,omitempty"`
}

// ProjectRoleParams filter parameters used in list operations
type ProjectRoleParams struct {
	UserID    string `url:"user_id,omitempty"`
	ProjectID string `url:"project_id,omitempty"`
}

// List gets a list of project_roles with optional filter params
func (c *ProjectRoleService) List(params *ProjectRoleParams) ([]ProjectRole, *http.Response, error) {
	output := &struct {
		Data []ProjectRole `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a project_role for the id specified
func (c *ProjectRoleService) Get(id string) (ProjectRole, *http.Response, error) {
	output := &struct {
		Data ProjectRole `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}
