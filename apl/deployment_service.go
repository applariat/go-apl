package apl

import (
	"github.com/dghubble/sling"
	"fmt"
	"net/http"
)

// DeploymentService is the service object for deployment operations
type DeploymentService struct {
	sling    *sling.Sling
	endpoint string
}

// NewDeploymentsService return a new DeploymentService
func NewDeploymentsService(sling *sling.Sling) *DeploymentService {
	return &DeploymentService{
		sling:    sling,
		endpoint: "deployments",
	}
}

// Deployment represents a deployment row
type Deployment struct {
	ID                   string `json:"id,omitempty"`
	StackVersionID       string `json:"stack_version_id,omitempty"`
	ProjectID            string `json:"project_id,omitempty"`
	WorkloadID           string `json:"workload_id,omitempty"`
	Name                 string `json:"name,omitempty"`
	ReleaseVersion       int    `json:"release_version,omitempty"`
	LeaseType            string `json:"lease_type,omitempty"`
	LeasePeriodDays      int    `json:"lease_period_days,omitempty"`
	LeaseExpirationEpoch int64  `json:"lease_expiration_epoch,omitempty"`
	WorkloadName         string `json:"workload_name,omitempty"`
	LeaseExpiration      string `json:"lease_expiration,omitempty"`
	QosLevel             string `json:"qos_level,omitempty"`
	StackVersion         interface{} `json:"stack_version,omitempty"`
	Location             interface{} `json:"location,omitempty"`
	Status               interface{} `json:"status,omitempty"`
	Stack                interface{} `json:"stack,omitempty"`
	Release              interface{}  `json:"release,omitempty"`
	CreatedTime          string `json:"created_time,omitempty"`
	LastModified         string `json:"last_modified,omitempty"`
	CreatedByUserID      string `json:"created_by_user_id,omitempty"`
	CreatedByUser 		 interface{} `json:"created_by_user,omitempty"`
}

// DeploymentInput is used for the update/create of deployments
type DeploymentCreateInput struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	ReleaseID       string `json:"release_id,omitempty"`
	LocDeployID     string `json:"loc_deploy_id,omitempty`
	LeaseType       string `json:"lease_type,omitempty"`
	LeasePeriodDays int    `json:"lease_period_days,omitempty"`
	QosLevel        string `json:"qos_level,omitempty"`
	Components      interface{}  `json:"components,omitempty"`
}

type DeploymentUpdateInput struct {
	Name                 string `json:"name,omitempty"`
	LeaseType            string `json:"lease_type,omitempty"`
	LeasePeriodDays      int    `json:"lease_period_days,omitempty"`
	LeaseExpirationEpoch int64  `json:"lease_expiration_epoch,omitempty"`
	WorkloadName         string `json:"workload_name,omitempty"`
	LeaseExpiration      string `json:"lease_expiration,omitempty"`
	QosLevel             string `json:"qos_level,omitempty"`
	Command              string `json:"command,omitempty"`
}

// DeploymentParams filter parameters
type DeploymentParams struct {
	Name                 string `url:"name,omitempty"`
	StackVersionID       string `url:"stack_version_id"`
	ProjectID            string `url:"project_id"`
	WorkloadID           string `url:"workload_id"`
	ReleaseVersion       int    `url:"release_version"`
	LeaseType            string `url:"lease_type"`
	LeasePeriodDays      int    `url:"lease_period_days"`
	LeaseExpirationEpoch int64  `url:"lease_expiration_epoch"`
	WorkloadName         string `url:"workload_name"`
	LeaseExpiration      string `url:"lease_expiration"`
	QosLevel             string `url:"qos_level"`
}

// List gets a list of deployments with optional filter params
func (c *DeploymentService) List(params *DeploymentParams) ([]Deployment, *http.Response, error) {
	output := &struct{ Data []Deployment `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a deployment for the id specified
func (c *DeploymentService) Get(id string) (Deployment, *http.Response, error) {
	output := &struct{ Data Deployment `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a deployment
func (c *DeploymentService) Create(input *DeploymentCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a deployment for the id specified
func (c *DeploymentService) Update(id string, input *DeploymentUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the deployment for the id specified
func (c *DeploymentService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}