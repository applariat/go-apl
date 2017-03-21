package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// StackVersionService is the service object for stackVersion operations
type StackVersionService struct {
	sling    *sling.Sling
	endpoint string
}

// NewStackVersionsService return a new stackVersionService
func NewStackVersionsService(sling *sling.Sling) *StackVersionService {
	return &StackVersionService{
		sling:    sling,
		endpoint: "stackversions",
	}
}


// StackVersion represents a stackVersion row
type StackVersion struct {
	StackID string `json:"stack_id"`
	StackVersions []struct {
		ID string `json:"id"`
		ProjectID string `json:"project_id"`
		Version int `json:"version"`
		Releases []struct {
			MetaData `json:"meta_data"`
			ID string `json:"id"`
			Version int `json:"version"`
		} `json:"releases"`
		VersionSub int `json:"version_sub"`
		CreatedByUserID string `json:"created_by_user_id"`
		LastModified string `json:"last_modified"`
		CreatedTime string `json:"created_time"`
		CreatedByUser `json:"created_by_user"`
	} `json:"stack_versions"`
}

// StackVersionInput is used for the update/create of stackVersions
type StackVersionInput struct {
	ID             string `json:"id,omitempty"`
	Name           string      `json:"name"`
	//StackVersions    interface{} `json:"stackVersions"`
}

// StackVersionParams filter parameters
type StackVersionParams struct {
	StackID string `url:"stack_id,omitempty"`
}

// listStackVersionsOutput used to wrap the data for API result
type listStackVersionsOutput struct {
	Data []StackVersion `json:"data"`
}

// getStackVersionOutput used to wrap the data for API result
type getStackVersionOutput struct {
	Data StackVersion `json:"data"`
}

// List gets a list of stackVersions with optional filter params
func (c *StackVersionService) List(params *StackVersionParams) (*[]StackVersion, *http.Response, error) {
	output := new(listStackVersionsOutput)
	apiError := new(APIError)

	resp, err := c.sling.New().Get(c.endpoint).QueryStruct(params).Receive(output, apiError)

	return &output.Data, resp, relevantError(err, apiError)
}

// Get get a stackVersion for the id specified
func (c *StackVersionService) Get(id string) (*StackVersion, *http.Response, error) {
	output := new(getStackVersionOutput)
	apiError := new(APIError)

	path := fmt.Sprintf("%s/%s", c.endpoint, id)

	fmt.Println("path:", output)
	resp, err := c.sling.New().Get(path).Receive(output, apiError)

	return &output.Data, resp, relevantError(err, apiError)
}

// Create will create a stackVersion
func (c *StackVersionService) Create(input *StackVersionInput) (*CreateOutput, *http.Response, error) {
	output := new(CreateOutput)
	apiError := new(APIError)

	body := &CreateInput{Data: input}
	resp, err := c.sling.New().Post(c.endpoint).BodyJSON(body).Receive(output, apiError)

	return output, resp, relevantError(err, apiError)

}

// Update will update a stackVersion for the id specified
func (c *StackVersionService) Update(id string, input *StackVersionInput) (*ModifyOutput, *http.Response, error) {
	output := new(ModifyOutput)
	apiError := new(APIError)

	body := &CreateInput{Data: input}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := c.sling.New().Put(path).BodyJSON(body).Receive(output, apiError)

	return output, resp, relevantError(err, apiError)
}

// Delete will delete the stackVersion for the id specified
func (c *StackVersionService) Delete(id string) (*ModifyOutput, *http.Response, error) {
	output := new(ModifyOutput)
	apiError := new(APIError)

	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := c.sling.New().Delete(path).Receive(output, apiError)

	return output, resp, relevantError(err, apiError)

}
