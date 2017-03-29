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
	ID string      `json:"id,omitempty"`

	CreatedByUserID string `json:"created_by_user_id"`
	Status interface{} `json:"status"`
	Workloads []string `json:"workloads"`
	IsDeleted bool `json:"is_deleted"`
	Name string `json:"name"`
	LocDeploysType string `json:"loc_deploys_type"`
	CommandPayload interface{} `json:"command_payload,omitempty"`
	AplManaged bool `json:"apl_managed"`

	CredentialID string `json:"credential_id"`
	Cluster interface{} `json:"cluster"`
	LastModified string `json:"last_modified"`
	Command string `json:"command,omitempty"`
	Policies interface{} `json:"policies,omitempty"`
	CreatedTime string `json:"created_time"`
	CredentialType string `json:"credential_type,omitempty"`
	Config interface{} `json:"config"`
	CreatedByUser `json:"created_by_user"`
}

// LocDeployCreateInput is used for the create of loc_deploys
type LocDeployCreateInput struct {
	ID string      `json:"id,omitempty"`
}

// LocDeployUpdateInput is used for the update of loc_deploys
type LocDeployUpdateInput struct {
	ID string      `json:"id,omitempty"`
}

// LocDeployParams filter parameters used in list operations
type LocDeployParams struct {
	Name string `url:"name,omitempty"`
}


// List gets a list of loc_deploys with optional filter params
func (c *LocDeployService) List(params *LocDeployParams) ([]LocDeploy, *http.Response, error) {
	output := &struct{ Data []LocDeploy `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a loc_deploy for the id specified
func (c *LocDeployService) Get(id string) (LocDeploy, *http.Response, error) {
	output := &struct{ Data LocDeploy `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a loc_deploy
func (c *LocDeployService) Create(input *LocDeployCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a loc_deploy for the id specified
func (c *LocDeployService) Update(id string, input *LocDeployUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the loc_deploy for the id specified
func (c *LocDeployService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
