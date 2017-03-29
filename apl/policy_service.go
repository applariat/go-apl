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
	ID               string `json:"id"`
	Name             string `json:"name"`
	PolicyType       string `json:"policy_type"`
	PolicyGroup      string `json:"policy_group"`
	PolicyTemplateID string `json:"policy_template_id"`
	Status           string `json:"status"`
	Return           string `json:"return"`

	Inputs           interface{} `json:"inputs,omitempty"`
	Assets           interface{} `json:"assets,omitempty"`
	Actions          interface{} `json:"actions,omitempty"`
	Operations       interface{} `json:"operations,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
	Constants        interface{} `json:"constants,omitempty"`

	LastModified     string `json:"last_modified"`
	CreatedTime      string `json:"created_time"`
	CreatedByUserID  string `json:"created_by_user_id"`
}

// PolicyCreateInput is used for the create of policies
type PolicyCreateInput struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name"`
	PolicyType       string `json:"policy_type"`
	PolicyGroup      string `json:"policy_group"`
	PolicyTemplateID string `json:"policy_template_id"`
	Return           string `json:"return"`

	Inputs           interface{} `json:"inputs,omitempty"`
	Assets           interface{} `json:"assets,omitempty"`
	Actions          interface{} `json:"actions,omitempty"`
	Operations       interface{} `json:"operations,omitempty"`
	Attributes       interface{} `json:"attributes,omitempty"`
	Constants        interface{} `json:"constants,omitempty"`
}

// PolicyUpdateInput is used for the update of policies
type PolicyUpdateInput struct {
	Name string `json:"name"`
}

// PolicyParams filter parameters used in list operations
type PolicyParams struct {
	Name             string `url:"name,omitempty"`
	PolicyType       string `url:"policy_type,omitempty"`
	PolicyGroup      string `url:"policy_group,omitempty"`
	PolicyTemplateID string `url:"policy_template_id,omitempty"`
	Return           string `url:"return,omitempty"`
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
