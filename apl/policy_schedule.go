package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// PolicyScheduleService is the service object for policy_schedule operations
type PolicyScheduleService struct {
	sling    *sling.Sling
	endpoint string
}

// NewPolicySchedulesService return a new PolicyScheduleService
func NewPolicySchedulesService(sling *sling.Sling) *PolicyScheduleService {
	return &PolicyScheduleService{
		sling:    sling,
		endpoint: "policy_schedules",
	}
}

// PolicySchedule represents a policy_schedule row
type PolicySchedule struct {
	ID string      `json:"id,omitempty"`
	CreatedByUserID string `json:"created_by_user_id"`
	Inputs interface{} `json:"inputs"`
	Name string `json:"name"`
	ResourceID string `json:"resource_id"`
	Schedule interface{} `json:"schedule"`
	OrgID string `json:"org_id"`
	LastModified string `json:"last_modified"`
	Status string `json:"status"`
	CreatedTime string `json:"created_time"`
	ResourceType string `json:"resource_type"`
	PolicyID string `json:"policy_id"`
}

// PolicyScheduleCreateInput is used for the create of policy_schedules
type PolicyScheduleCreateInput struct {
	ID string      `json:"id,omitempty"`
}

// PolicyScheduleUpdateInput is used for the update of policy_schedules
type PolicyScheduleUpdateInput struct {
	ID string      `json:"id,omitempty"`
}

// PolicyScheduleParams filter parameters used in list operations
type PolicyScheduleParams struct {
	Name string `url:"name,omitempty"`
}


// List gets a list of policy_schedules with optional filter params
func (c *PolicyScheduleService) List(params *PolicyScheduleParams) ([]PolicySchedule, *http.Response, error) {
	output := &struct{ Data []PolicySchedule `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a policy_schedule for the id specified
func (c *PolicyScheduleService) Get(id string) (PolicySchedule, *http.Response, error) {
	output := &struct{ Data PolicySchedule `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a policy_schedule
func (c *PolicyScheduleService) Create(input *PolicyScheduleCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a policy_schedule for the id specified
func (c *PolicyScheduleService) Update(id string, input *PolicyScheduleUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the policy_schedule for the id specified
func (c *PolicyScheduleService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
