package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// LocArtifactService is the service object for loc_artifact operations
type LocArtifactService struct {
	sling    *sling.Sling
	endpoint string
}

// NewLocArtifactsService return a new LocArtifactService
func NewLocArtifactsService(sling *sling.Sling) *LocArtifactService {
	return &LocArtifactService{
		sling:    sling,
		endpoint: "loc_artifacts",
	}
}

// LocArtifact represents a loc_artifact row
type LocArtifact struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name"`
	LocArtifactsType   string `json:"loc_artifacts_type"`
	Bucket             string `json:"bucket,omitempty"`
	CredentialID       string `json:"credential_id"`
	CredentialType     string `json:"credential_type,omitempty"`
	SecretCredentialID string `json:"secret_credential_id,omitempty"`
	RegistryURI        string `json:"registry_uri,omitempty"`
	ProjectBlacklist   interface{} `json:"project_blacklist,omitempty"`
	SupportedTypes     interface{} `json:"supported_types,omitempty"`
	Metadata           interface{} `json:"metadata,omitempty"`
	URL                string `json:"url,omitempty"`
	LastModified       string `json:"last_modified"`
	CreatedTime        string `json:"created_time"`
	CreatedByUser `json:"created_by_user"`
}

// LocArtifactCreateInput is used for the create of loc_artifacts
type LocArtifactCreateInput struct {
	ID                 string      `json:"id,omitempty"`
	Name               string `json:"name"`
	LocArtifactsType   string `json:"loc_artifacts_type"`
	Bucket             string `json:"bucket,omitempty"`
	CredentialID       string `json:"credential_id"`
	CredentialType     string `json:"credential_type,omitempty"`
	SecretCredentialID string `json:"secret_credential_id,omitempty"`
	RegistryURI        string `json:"registry_uri,omitempty"`

	ProjectBlacklist   interface{} `json:"project_blacklist,omitempty"`
	SupportedTypes     interface{} `json:"supported_types,omitempty"`
	URL                string `json:"url,omitempty"`
}

// LocArtifactUpdateInput is used for the update of loc_artifacts
type LocArtifactUpdateInput struct {
	Name string `json:"name,omitempty"`
}

// LocArtifactParams filter parameters used in list operations
type LocArtifactParams struct {
	Name               string `url:"name,omitempty"`
	LocArtifactsType   string `url:"loc_artifacts_type,omitempty"`
	Bucket             string `url:"bucket,omitempty"`
	CredentialID       string `url:"credential_id,omitempty"`
	CredentialType     string `url:"credential_type,omitempty"`
	SecretCredentialID string `url:"secret_credential_id,omitempty"`
	RegistryURI        string `url:"registry_uri,omitempty"`
	URL                string `url:"url,omitempty"`
}


// List gets a list of loc_artifacts with optional filter params
func (c *LocArtifactService) List(params *LocArtifactParams) ([]LocArtifact, *http.Response, error) {
	output := &struct{ Data []LocArtifact `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a loc_artifact for the id specified
func (c *LocArtifactService) Get(id string) (LocArtifact, *http.Response, error) {
	output := &struct{ Data LocArtifact `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a loc_artifact
func (c *LocArtifactService) Create(input *LocArtifactCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a loc_artifact for the id specified
func (c *LocArtifactService) Update(id string, input *LocArtifactUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the loc_artifact for the id specified
func (c *LocArtifactService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
