package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)


// ReleaseService is the service object for release operations
type ReleaseService struct {
	sling    *sling.Sling
	endpoint string
}

// NewReleasesService return a new ReleaseService
func NewReleasesService(sling *sling.Sling) *ReleaseService {
	return &ReleaseService{
		sling:    sling,
		endpoint: "releases",
	}
}

// Release represents a release row
type Release struct {
	ID              string      `json:"id,omitempty"`
	Version         int `json:"version"`
	StackVersionID  string `json:"stack_version_id"`
	StackID         string `json:"stack_id"`
	LastModified    string `json:"last_modified"`
	Components      interface{} `json:"components"`
	CreatedTime     string `json:"created_time"`
	ProjectID       string `json:"project_id"`
	CreatedByUser `json:"created_by_user"`
	BuildStatus     string `json:"build_status,omitempty"`
	LocImageID      string `json:"loc_image_id,omitempty"`
}

// ReleaseCreateInput is used for the create of releases
type ReleaseCreateInput struct {
	ID string      `json:"id,omitempty"`
}

// ReleaseParams filter parameters used in list operations
type ReleaseParams struct {
	Name string `url:"name,omitempty"`
}

// List gets a list of releases with optional filter params
func (c *ReleaseService) List(params *ReleaseParams) ([]Release, *http.Response, error) {
	output := &struct{ Data []Release `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a release for the id specified
func (c *ReleaseService) Get(id string) (Release, *http.Response, error) {
	output := &struct{ Data Release `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a release
func (c *ReleaseService) Create(input *ReleaseCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Delete will delete the release for the id specified
func (c *ReleaseService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
