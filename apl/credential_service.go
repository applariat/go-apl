package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// CredentialService is the service object for credential operations
type CredentialService struct {
	sling    *sling.Sling
	endpoint string
}

// NewCredentialsService return a new credentialService
func NewCredentialsService(sling *sling.Sling) *CredentialService {
	return &CredentialService{
		sling:    sling,
		endpoint: "credentials",
	}
}

// Credential represents a credential row
type Credential struct {
	ID              string      `json:"id,omitempty"`
	CredentialType  string      `json:"credential_type"`
	Name            string      `json:"name"`
	Credentials     interface{} `json:"credentials"`
	CreatedByUserID string      `json:"created_by_user_id"`
	LastModified    string      `json:"last_modified"`
	CreatedTime     string      `json:"created_time"`
}

// CredentialCreateInput is used for the create of credentials
type CredentialCreateInput struct {
	ID             string      `json:"id,omitempty"`
	CredentialType string 	   `json:"credential_type"`
	Name           string      `json:"name"`
	Credentials    interface{} `json:"credentials"`
}

// CredentialUpdateInput is used for the update of credentials
type CredentialUpdateInput struct {
	ID             string      `json:"id,omitempty"`
	Name           string      `json:"name"`
	Credentials    interface{} `json:"credentials"`
}

// CredentialParams filter parameters used in list operations
type CredentialParams struct {
	CredentialType string `url:"credential_type,omitempty"`
	Name           string `url:"name,omitempty"`
}


// List gets a list of credentials with optional filter params
func (c *CredentialService) List(params *CredentialParams) ([]Credential, *http.Response, error) {
	output := &struct{Data []Credential `json:"data"`}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a credential for the id specified
func (c *CredentialService) Get(id string) (Credential, *http.Response, error) {
	output := &struct{Data Credential `json:"data"`}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a credential
func (c *CredentialService) Create(input *CredentialCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a credential for the id specified
func (c *CredentialService) Update(id string, input *CredentialUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the credential for the id specified
func (c *CredentialService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
