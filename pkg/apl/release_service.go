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
	ID             string      `json:"id,omitempty"`
	Version        int         `json:"version"`
	StackID        string      `json:"stack_id"`
	StackVersionID string      `json:"stack_version_id"`
	ProjectID      string      `json:"project_id"`
	LocImageID     string      `json:"loc_image_id,omitempty"`
	BuildStatus    string      `json:"build_status,omitempty"`
	Components     interface{} `json:"components"`
	MetaData       interface{} `json:"meta_data,omitempty"`

	LastModified  string `json:"last_modified"`
	CreatedTime   string `json:"created_time"`
	CreatedByUser `json:"created_by_user"`
}

// ReleaseCreateInput is used for the create of releases
type ReleaseCreateInput struct {
	ID             string      `json:"id,omitempty"`
	Version        int         `json:"version,omitempty"`
	StackID        string      `json:"stack_id"`
	StackVersionID string      `json:"stack_version_id"`
	ProjectID      string      `json:"project_id,omitempty"`
	LocImageID     string      `json:"loc_image_id,omitempty"`
	BuildStatus    string      `json:"build_status,omitempty"`
	Components     interface{} `json:"components"`
	MetaData       interface{} `json:"meta_data,omitempty"`
}

type ReleaseOverrideArtifactBase struct {
	StackArtifactID string `json:"stack_artifact_id,omitempty"`
}

type ReleaseOverrideArtifact struct {
	Builder *ReleaseOverrideArtifactBase `json:"builder,omitempty"`
	Code    *ReleaseOverrideArtifactBase `json:"code,omitempty"`
	Image   *ReleaseOverrideArtifactBase `json:"image,omitempty"`
} // `json:"artifacts"`

type ReleaseOverrideRelease struct {
	Artifacts ReleaseOverrideArtifact `json:"artifacts,omitempty"`
} // `json:"release"`

type ReleaseOverrideService struct {
	Name    string                 `json:"name"`
	Release ReleaseOverrideRelease `json:"release"`
} // `json:"services"`

type ReleaseOverrideComponent struct {
	Name             string                   `json:"name"`
	StackComponentID string                   `json:"stack_component_id"`
	Services         []ReleaseOverrideService `json:"services"`
} // `json:"components"`

// ReleaseParams filter parameters used in list operations
type ReleaseParams struct {
	Name           string `url:"name,omitempty"`
	Version        string `url:"version,omitempty"`
	StackID        string `url:"stack_id,omitempty"`
	StackVersionID string `url:"stack_version_id,omitempty"`
	ProjectID      string `url:"project_id,omitempty"`
	LocImageID     string `url:"loc_image_id,omitempty"`
	BuildStatus    string `url:"build_status,omitempty"`
}

// List gets a list of releases with optional filter params
func (c *ReleaseService) List(params *ReleaseParams) ([]Release, *http.Response, error) {
	output := &struct {
		Data []Release `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a release for the id specified
func (c *ReleaseService) Get(id string) (Release, *http.Response, error) {
	output := &struct {
		Data Release `json:"data"`
	}{}
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
