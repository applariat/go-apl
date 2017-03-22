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
	CreatedByUserID string      `json:"created_by_user_id,omitempty"`
	LastModified    string      `json:"last_modified,omitempty"`
	CreatedTime     string      `json:"created_time,omitempty"`
}

// CredentialInput is used for the update/create of credentials
type CredentialInput struct {
	ID             string      `json:"id,omitempty"`
	CredentialType string 	   `json:"credential_type"`
	Name           string      `json:"name"`
	Credentials    interface{} `json:"credentials"`
}

// CredentialParams filter parameters
type CredentialParams struct {
	CredentialType string `url:"credential_type,omitempty"`
	Name           string `url:"name,omitempty"`
}

// listCredentialsOutput used to wrap the data for API result
type listCredentialsOutput struct {
	Data []Credential `json:"data"`
}

// getCredentialOutput used to wrap the data for API result
type getCredentialOutput struct {
	Data Credential `json:"data"`
}

// List gets a list of credentials with optional filter params
func (c *CredentialService) List(params *CredentialParams) ([]Credential, *http.Response, error) {
	output := new(listCredentialsOutput)
	apiError := new(APIError)

	resp, err := c.sling.New().Get(c.endpoint).QueryStruct(params).Receive(output, apiError)
	return output.Data, resp, relevantError(err, apiError)
}

// Get get a credential for the id specified
func (c *CredentialService) Get(id string) (Credential, *http.Response, error) {
	output := new(getCredentialOutput)
	apiError := new(APIError)

	path := fmt.Sprintf("%s/%s", c.endpoint, id)

	resp, err := c.sling.New().Get(path).Receive(output, apiError)

	return output.Data, resp, relevantError(err, apiError)
}

// Create will create a credential
func (c *CredentialService) Create(input *CredentialInput) (CreateResult, *http.Response, error) {
	output := CreateResult{}
	apiError := new(APIError)

	body := &CreateInput{Data: input}
	resp, err := c.sling.New().Post(c.endpoint).BodyJSON(body).Receive(&output, apiError)

	return output, resp, relevantError(err, apiError)
}

// Update will update a credential for the id specified
func (c *CredentialService) Update(id string, input *CredentialInput) (ModifyResult, *http.Response, error) {
	output := ModifyOutput{}
	apiError := new(APIError)

	body := &CreateInput{Data: input}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := c.sling.New().Put(path).BodyJSON(body).Receive(&output, apiError)

	return output.ModifyResult, resp, relevantError(err, apiError)
}

// Delete will delete the credential for the id specified
func (c *CredentialService) Delete(id string) (ModifyResult, *http.Response, error) {
	output := ModifyOutput{}
	apiError := new(APIError)

	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := c.sling.New().Delete(path).Receive(&output, apiError)

	return output.ModifyResult, resp, relevantError(err, apiError)
}
