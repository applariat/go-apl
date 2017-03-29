package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// PolicyService is the service object for policy operations
type PolicyService struct {
	sling    *sling.Sling
	endpoint string
}

// NewPolicyService return a new PolicyService
func NewPolicyService(sling *sling.Sling) *PolicyService {
	return &PolicyService{
		sling:    sling,
		endpoint: "policies",
	}
}

// Policy represents a policy row
type Policy struct {
	ID               string      `json:"id,omitempty"`

	CreatedByUserID  string `json:"created_by_user_id"`
	Inputs           interface{} `json:"inputs,omitempty"`
	Return           string `json:"return"`
	Assets           interface{} `json:"assets"`
	Actions          interface{} `json:"actions"`
	PolicyType       string `json:"policy_type"`
	LastModified     string `json:"last_modified"`
	Status           string `json:"status"`
	PolicyGroup      string `json:"policy_group"`
	Operations       interface{} `json:"operations"`
	CreatedTime      string `json:"created_time"`
	Attributes       interface{} `json:"attributes"`
	PolicyTemplateID string `json:"policy_template_id"`
	Name             string `json:"name"`
	Constants        interface{} `json:"constants,omitempty"`
}

// PolicyCreateInput is used for the create of policies
type PolicyCreateInput struct {
	ID string      `json:"id,omitempty"`
}

// PolicyUpdateInput is used for the update of policies
type PolicyUpdateInput struct {
	ID string      `json:"id,omitempty"`
}

// PolicyParams filter parameters used in list operations
type PolicyParams struct {
	Name string `url:"name,omitempty"`
}


// List gets a list of policies with optional filter params
func (c *PolicyService) List(params *PolicyParams) ([]Policy, *http.Response, error) {
	output := &struct{ Data []Policy `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a policy for the id specified
func (c *PolicyService) Get(id string) (Policy, *http.Response, error) {
	output := &struct{ Data Policy `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a policy
func (c *PolicyService) Create(input *PolicyCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a policy for the id specified
func (c *PolicyService) Update(id string, input *PolicyUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the policy for the id specified
func (c *PolicyService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
