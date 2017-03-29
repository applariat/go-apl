package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// OrgService is the service object for org operations
type OrgService struct {
	sling    *sling.Sling
	endpoint string
}

// NewOrgsService return a new OrgService
func NewOrgsService(sling *sling.Sling) *OrgService {
	return &OrgService{
		sling:    sling,
		endpoint: "orgs",
	}
}

// Org represents a org row
type Org struct {
	ID             string `json:"id,omitempty"`
	Country        string `json:"country"`
	NumOfEmployees string `json:"num_of_employees"`
	OrgType        string `json:"org_type"`
	LastModified   string `json:"last_modified"`
	CompanyName    string `json:"company_name"`
	CreatedTime    string `json:"created_time"`
	IsDeleted      bool `json:"is_deleted"`
}

// OrgUpdateInput is used for the update of orgs
type OrgUpdateInput struct {

	CompanyName    string `json:"company_name,omitempty"`
	OrgType        string `json:"org_type,omitempty"`
	Country        string `json:"country,omitempty"`
	NumOfEmployees string `json:"num_of_employees,omitempty"`
}

// List gets a list of orgs with optional filter params
func (c *OrgService) List() (Org, *http.Response, error) {
	output := &struct{ Data Org `json:"data"` }{}
	resp, err := doGet(c.sling, c.endpoint, output)
	return output.Data, resp, err
}

// Get get a org for the id specified
func (c *OrgService) Get(id string) (Org, *http.Response, error) {
	output := &struct{ Data Org `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Update will update a org for the id specified
func (c *OrgService) Update(id string, input *OrgUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

