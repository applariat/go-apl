package apl

import (
	"github.com/dghubble/sling"
)

// Client is a client to access the appLariat API
type Client struct {
	sling *sling.Sling

	// Different appLariat API Services
	Credentials *CredentialService
	//Types       *TypeService
	Components  *ComponentService
	//StackVersions *StackVersionService
	StackArtifacts  *StackArtifactService
	Deployments 	*DeploymentService
}

// NewClient returns the client object to access the applariat API
func NewClient() *Client {

	base := sling.New().Client(getOauth2HTTPClient()).Base(APLConfig.API)

	return &Client{
		Credentials: NewCredentialsService(base.New()),
		//Types:       NewTypesService(base.New()),
		Components:  NewComponentsService(base.New()),
		//StackVersions: NewStackVersionsService(base.New()),
		StackArtifacts: NewStackArtifactsService(base.New()),
		Deployments: NewDeploymentsService(base.New()),
	}

}
