package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// LocDeployService is the service object for loc_deploy operations
type LocDeployService struct {
	sling    *sling.Sling
	endpoint string
}

// NewLocDeploysService return a new LocDeployService
func NewLocDeploysService(sling *sling.Sling) *LocDeployService {
	return &LocDeployService{
		sling:    sling,
		endpoint: "loc_deploys",
	}
}

// LocDeploy represents a loc_deploy row
type LocDeploy struct {
	ID             string            `json:"id,omitempty"`
	Name           string            `json:"name"`
	OrgID          string            `json:"org_id,omitempty"`
	IsDeleted      bool              `json:"is_deleted"`
	AplManaged     bool              `json:"apl_managed"`
	AplOnApl       bool              `json:"apL_on_apl"`
	LocalScript    bool              `json:"local_script"`
	LocDeploysType string            `json:"loc_deploys_type"`
	CredentialID   string            `json:"credential_id"`
	CredentialType string            `json:"credential_type,omitempty"`
	LastModified   string            `json:"last_modified"`
	CreatedTime    string            `json:"created_time"`
	Description    string            `json:"description"`
	Workloads      []string          `json:"workloads"`
	Status         LocDeployStatus   `json:"status"`
	DNS            interface{} `json:"dns,omitempty"`
	Cluster        interface{}       `json:"cluster"`
	//CommandPayload interface{}       `json:"command_payload,omitempty"`
	//Policies      interface{}       `json:"policies,omitempty"`
	Config        interface{}       `json:"config"`
	Metadata      map[string]string `json:"meta_data,omitempty"`
	CreatedByUser interface{}          `json:"created_by_user"`
}

// LocDeployCreateInput is used for the create of loc_deploys
type LocDeployCreateInput struct {
	ID             string          `json:"id,omitempty"`
	Name           string          `json:"name"`
	ProjectID      string          `json:"project_id"`
	LocDeploysType string          `json:"loc_deploys_type"`
	CredentialID   string          `json:"credential_id"`
	CredentialType string          `json:"credential_type,omitempty"`
	Description    string          `json:"description,omitempty"`
	Status         LocDeployStatus `json:"status,omitempty"`
	Workloads      interface{}     `json:"workloads"`
	Config         interface{}     `json:"config"`
	//Policies       interface{}     `json:"policies"`
}

//// LocDeployUpdateInput is used for the update of loc_deploys
//type LocDeployUpdateInput struct {
//	Command string `json:"command"`
//}

type LocDeployStatus struct {
	RunningNodes     int    `json:"runningNodes,omitempty"`
	IdleNodes        int    `json:"idleNodes,omitempty"`
	Description      string `json:"description,omitempty"`
	State            string `json:"state,omitempty"`
	AvalableCapacity string `json:"availableCapacity,omitempty"`
	StoppedNodes     int    `json:"stoppedNodes,omitempty"`
}

// LocDeployParams filter parameters used in list operations
type LocDeployParams struct {
	Name           string `url:"name,omitempty"`
	LocDeploysType string `url:"loc_deploys_type,omitempty"`
	CredentialID   string `url:"credential_id,omitempty"`
	CredentialType string `url:"credential_type,omitempty"`
}

// List gets a list of loc_deploys with optional filter params
func (c *LocDeployService) List(params *LocDeployParams) ([]LocDeploy, *http.Response, error) {
	output := &struct {
		Data []LocDeploy `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a loc_deploy for the id specified
func (c *LocDeployService) Get(id string) (LocDeploy, *http.Response, error) {
	output := &struct {
		Data LocDeploy `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a loc_deploy
func (c *LocDeployService) Create(input *LocDeployCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

//// Update will update a loc_deploy for the id specified
//func (c *LocDeployService) Update(id string, input *LocDeployUpdateInput) (ModifyResult, *http.Response, error) {
//	path := fmt.Sprintf("%s/%s", c.endpoint, id)
//	return doUpdate(c.sling, path, input)
//}

// Delete will delete the loc_deploy for the id specified
func (c *LocDeployService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
