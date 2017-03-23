package apl

import (
	"github.com/dghubble/sling"
	"fmt"
	"net/http"
)

// StackArtifactService is the service object for stack_artifact operations
type StackArtifactService struct {
	sling    *sling.Sling
	endpoint string
}

// NewStackArtifactsService return a new StackArtifactService
func NewStackArtifactsService(sling *sling.Sling) *StackArtifactService {
	return &StackArtifactService{
		sling:    sling,
		endpoint: "stack_artifacts",
	}
}

// StackArtifact represents a stack_artifact row
type StackArtifact struct {
	ID              string `json:"id,omitempty"`
	LocArtifactID   string `json:"loc_artifact_id,omitempty"`
	ProjectID       string `json:"project_id,omitempty"`
	StackID         string `json:"stack_id,omitempty"`
	Name            string `json:"name,omitempty"`
	Type            string `json:"type,omitempty"`
	Version         string `json:"version,omitempty"`
	// TODO: Remove FileName, it is not supported later on. Replaced with ArtifactName
	FileName        string `json:"file_name,omitempty"`
	ArtifactName    string `json:"artifact_name,omitempty"`
	Package         string `json:"package,omitempty"`
	MetaData `json:"meta_data,omitempty"`
	CreatedByUserID string `json:"created_by_user_id,omitempty"`
	CreatedTime     string `json:"created_time,omitempty"`
	LastModified    string `json:"last_modified,omitempty"`
}

// StackArtifactInput is used for the update/create of stack_artifacts
type StackArtifactCreateInput struct {
	ID            string `json:"id,omitempty"`
	LocArtifactID string `json:"loc_artifact_id,omitempty"`
	ProjectID     string `json:"project_id,omitempty"`
	StackID       string `json:"stack_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty,omitempty"`
	Version       string `json:"version,omitempty"`
	ArtifactName  string `json:"artifact_name,omitempty"`
	Package       string `json:"package,omitempty"`
	MetaData `json:"meta_data,omitempty"`
}

type StackArtifactUpdateInput struct {
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
	Version       string `json:"version,omitempty"`
	ArtifactName  string `json:"artifact_name,omitempty"`
	Package       string `json:"package,omitempty"`
	MetaData `json:"meta_data,omitempty"`
}

// StackArtifactParams filter parameters
type StackArtifactParams struct {
	LocArtifactID string `url:"loc_artifact_id,omitempty"`
	ProjectID     string `url:"project_id,omitempty"`
	StackID       string `url:"stack_id,omitempty"`
	Name          string `url:"name,omitempty"`
	Type          string `url:"type,omitempty"`
	Version       string `url:"version,omitempty"`
	ArtifactName  string `url:"artifact_name,omitempty"`
	Package       string `url:"package,omitempty"`
}

// List gets a list of stack_artifacts with optional filter params
func (c *StackArtifactService) List(params *StackArtifactParams) ([]StackArtifact, *http.Response, error) {
	output := &struct{Data []StackArtifact `json:"data"`}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a stack_artifact for the id specified
func (c *StackArtifactService) Get(id string) (StackArtifact, *http.Response, error) {
	output := &struct{Data StackArtifact `json:"data"`}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a stack_artifact
func (c *StackArtifactService) Create(input *StackArtifactCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a stack_artifact for the id specified
func (c *StackArtifactService) Update(id string, input *StackArtifactUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the stack_artifact for the id specified
func (c *StackArtifactService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}