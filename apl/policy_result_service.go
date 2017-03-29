package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// PolicyResultService is the service object for policy_result operations
type PolicyResultService struct {
	sling    *sling.Sling
	endpoint string
}

// NewPolicyResultsService return a new PolicyResultService
func NewPolicyResultsService(sling *sling.Sling) *PolicyResultService {
	return &PolicyResultService{
		sling:    sling,
		endpoint: "policy_results",
	}
}

// PolicyResult represents a policy_result row
type PolicyResult struct {
	ID               string `json:"id"`
	PolicyID         string `json:"policy_id"`
	PolicyScheduleID string `json:"policy_schedule_id,omitempty"`
	ProjectID        string `json:"project_id,omitempty"`

	LastModified     string `json:"last_modified"`
	CreatedTime      string `json:"created_time"`
	CreatedByUserID  string `json:"created_by_user_id"`
}

// PolicyResultCreateInput is used for the create of policy_results
type PolicyResultCreateInput struct {
	ID               string `json:"id,omitempty"`
	PolicyID         string `json:"policy_id"`
	PolicyScheduleID string `json:"policy_schedule_id,omitempty"`
	ProjectID        string `json:"project_id,omitempty"`
}

//// PolicyResultUpdateInput is used for the update of policy_results
//type PolicyResultUpdateInput struct {
//	ID string      `json:"id,omitempty"`
//}

// PolicyResultParams filter parameters used in list operations
type PolicyResultParams struct {
	PolicyID         string `url:"policy_id,omitempty"`
	PolicyScheduleID string `url:"policy_schedule_id,omitempty"`
	ProjectID        string `url:"project_id,omitempty"`
}


// List gets a list of policy_results with optional filter params
func (c *PolicyResultService) List(params *PolicyResultParams) ([]PolicyResult, *http.Response, error) {
	output := &struct{ Data []PolicyResult `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a policy_result for the id specified
func (c *PolicyResultService) Get(id string) (PolicyResult, *http.Response, error) {
	output := &struct{ Data PolicyResult `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a policy_result
func (c *PolicyResultService) Create(input *PolicyResultCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

//// Update will update a policy_result for the id specified
//func (c *PolicyResultService) Update(id string, input *PolicyResultUpdateInput) (ModifyResult, *http.Response, error) {
//	path := fmt.Sprintf("%s/%s", c.endpoint, id)
//	return doUpdate(c.sling, path, input)
//}
//
//// Delete will delete the policy_result for the id specified
//func (c *PolicyResultService) Delete(id string) (ModifyResult, *http.Response, error) {
//	path := fmt.Sprintf("%s/%s", c.endpoint, id)
//	return doDelete(c.sling, path)
//}
