package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// WorkloadService is the service object for workload operations
type WorkloadService struct {
	sling    *sling.Sling
	endpoint string
}

// NewWorkloadsService return a new workloadService
func NewWorkloadsService(sling *sling.Sling) *WorkloadService {
	return &WorkloadService{
		sling:    sling,
		endpoint: "workloads",
	}
}

// Workload represents a workload row
type Workload struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	WorkloadType       string `json:"workload_type"`
	LeaseType          string `json:"lease_type"`
	MaxLeasePeriodDays int    `json:"max_lease_period_days"`
	Priority           int    `json:"priority"`
	QualityOfService   string `json:"quality_of_service"`
	LastModified       string `json:"last_modified"`
	CreatedTime        string `json:"created_time"`
}

// WorkloadUpdateInput is used for the update of workloads
type WorkloadUpdateInput struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	WorkloadType       string `json:"workload_type,omitempty"`
	LeaseType          string `json:"lease_type,omitempty"`
	MaxLeasePeriodDays int    `json:"max_lease_period_days,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	QualityOfService   string `json:"quality_of_service,omitempty"`
}

// WorkloadParams filter parameters used in list operations
type WorkloadParams struct {
	Name               string `url:"name,omitempty"`
	WorkloadType       string `url:"workload_type,omitempty"`
	LeaseType          string `url:"lease_type,omitempty"`
	MaxLeasePeriodDays int    `url:"max_lease_period_days,omitempty"`
	Priority           int    `url:"priority,omitempty"`
	QualityOfService   string `url:"quality_of_service,omitempty"`
}

// List gets a list of workloads with optional filter params
func (c *WorkloadService) List(params *WorkloadParams) ([]Workload, *http.Response, error) {
	output := &struct {
		Data []Workload `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a workload for the id specified
func (c *WorkloadService) Get(id string) (Workload, *http.Response, error) {
	output := &struct {
		Data Workload `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Update will update a workload for the id specified
func (c *WorkloadService) Update(id string, input *WorkloadUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}
