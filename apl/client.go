package apl

import (
	"github.com/dghubble/sling"
)

// Client is a client to access the appLariat API
type Client struct {
	sling *sling.Sling

	// Different appLariat API Services
	Audits 			*AuditService
	Credentials 	*CredentialService
	Components  	*ComponentService
	Deployments 	*DeploymentService
	Events 			*EventService
	Forms 			*FormService
	Jobs 			*JobService
	LocArtifacts 	*LocArtifactService
	ProjectRoles 	*ProjectRoleService
	Roles 			*RoleService
	StackArtifacts  *StackArtifactService
	Types 			*TypeService
	Users			*UserService
	Workloads 		*WorkloadService

	//StackVersions *StackVersionService

}

// NewClient returns the client object to access the applariat API
func NewClient() *Client {

	base := sling.New().Client(getOauth2HTTPClient()).Base(APLConfig.API)

	return &Client{
		Audits: NewAuditsService(base.New()),
		Credentials: NewCredentialsService(base.New()),
		Components:  NewComponentsService(base.New()),
		Deployments: NewDeploymentsService(base.New()),
		Events: NewEventsService(base.New()),
		Forms: NewFormsService(base.New()),
		Jobs: NewJobsService(base.New()),
		LocArtifacts: NewLocArtifactsService(base.New()),
		ProjectRoles: NewProjectRolesService(base.New()),
		Roles: NewRolesService(base.New()),
		StackArtifacts: NewStackArtifactsService(base.New()),
		Types: NewTypesService(base.New()),
		Users: NewUsersService(base.New()),
		Workloads: NewWorkloadsService(base.New()),

		//StackVersions: NewStackVersionsService(base.New()),
	}

}
