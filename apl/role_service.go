package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// RoleService is the service object for role operations
type RoleService struct {
	sling    *sling.Sling
	endpoint string
}

// NewRolesService return a new roleService
func NewRolesService(sling *sling.Sling) *RoleService {
	return &RoleService{
		sling:    sling,
		endpoint: "roles",
	}
}

// Role represents a role row
type Role struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	AccessLevel  int    `json:"access_level"`
	Workloads    interface{} `json:"workloads,omitempty"`
	Permissions  interface{} `json:"permissions,omitempty"`
	LastModified string `json:"last_modified"`
	CreatedTime  string `json:"created_time"`
}

// RoleParams filter parameters used in list operations
type RoleParams struct {
	Name        string `url:"name,omitempty"`
	Role        string `url:"role,omitempty"`
	AccessLevel int    `url:"access_level,omitempty"`
}

// RoleUpdateInput ...
type RoleUpdateInput struct {
	Workloads []string `json:"workloads"`
}


// List gets a list of roles with optional filter params
func (c *RoleService) List(params *RoleParams) ([]Role, *http.Response, error) {
	output := &struct{ Data []Role `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a role for the id specified
func (c *RoleService) Get(id string) (Role, *http.Response, error) {
	output := &struct{ Data Role `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Update will update a role for the id specified
func (c *RoleService) Update(id string, input *RoleUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}
