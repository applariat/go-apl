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
	LocDeploys  	*LocDeployService
	Orgs 			*OrgService
	PolicyResults 	*PolicyResultService
	PolicySchedules *PolicyScheduleService
	Policies 		*PolicyService
	ProjectRoles 	*ProjectRoleService
	Projects 		*ProjectService
	Releases 		*ReleaseService
	Roles 			*RoleService
	StackArtifacts  *StackArtifactService
	StackComponents *StackComponentService
	Stacks 			*StackService
	StackVersions 	*StackVersionService
	Types 			*TypeService
	Users			*UserService
	Workloads 		*WorkloadService
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
		LocDeploys: NewLocDeploysService(base.New()),
		Orgs: NewOrgsService(base.New()),
		PolicyResults: NewPolicyResultsService(base.New()),
		PolicySchedules: NewPolicySchedulesService(base.New()),
		Policies: NewPolicyService(base.New()),
		ProjectRoles: NewProjectRolesService(base.New()),
		Projects: NewProjectsService(base.New()),
		Releases: NewReleasesService(base.New()),
		Roles: NewRolesService(base.New()),
		StackArtifacts: NewStackArtifactsService(base.New()),
		StackComponents: NewStackComponentsService(base.New()),
		Stacks: NewStacksService(base.New()),
		StackVersions: NewStackVersionsService(base.New()),
		Types: NewTypesService(base.New()),
		Users: NewUsersService(base.New()),
		Workloads: NewWorkloadsService(base.New()),
	}

}
