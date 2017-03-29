package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// JobService is the service object for job operations
type JobService struct {
	sling    *sling.Sling
	endpoint string
}

// NewJobsService return a new JobService
func NewJobsService(sling *sling.Sling) *JobService {
	return &JobService{
		sling:    sling,
		endpoint: "jobs",
	}
}

// Job represents a job row
type Job struct {
	ID              string      `json:"id,omitempty"`
	Status          interface{} `json:"status"`
	Endpoint        string `json:"endpoint"`
	ResourceID      string `json:"resource_id"`
	ResourceType    string `json:"resource_type"`
	Action          string `json:"action"`
	Payload         interface{} `json:"payload"`
	Result          interface{} `json:"result"`
	CreatedByUserID string `json:"created_by_user_id"`
	CreatedTime     string `json:"created_time"`
	LastModified    string `json:"last_modified"`
}

// JobUpdateInput is used for the update of jobs
type JobUpdateInput struct {
	Status          interface{} `json:"status,omitempty"`
	Result          interface{} `json:"result,omitempty"`
}

// JobParams filter parameters used in list operations
type JobParams struct {
	Endpoint        string `url:"endpoint,omitempty"`
	ResourceID      string `url:"resource_id,omitempty"`
	ResourceType    string `url:"resource_type,omitempty"`
	Action          string `url:"action,omitempty"`
	CreatedByUserID string `url:"created_by_user_id,omitempty"`
}


// List gets a list of jobs with optional filter params
func (c *JobService) List(params *JobParams) ([]Job, *http.Response, error) {
	output := &struct{ Data []Job `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a job for the id specified
func (c *JobService) Get(id string) (Job, *http.Response, error) {
	output := &struct{ Data Job `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Update will update a job for the id specified
func (c *JobService) Update(id string, input *JobUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

