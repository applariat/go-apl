package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	projectRoleFilterUserID    string
	projectRoleFilterProjectID string

	projectRolesCmd    = createListCommand(cmdListProjectRoles, "project-roles", "")
	projectRolesGetCmd = createGetCommand(cmdGetProjectRoles, "project-role", "")
)

func init() {

	// command flags
	projectRolesCmd.Flags().StringVar(&projectRoleFilterUserID, "user-id", "", "Filter project-roles by user_id")
	projectRolesCmd.Flags().StringVar(&projectRoleFilterProjectID, "project-id", "", "Filter project-roles by project_id")

	// add sub commands
	projectRolesCmd.AddCommand(projectRolesGetCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(projectRolesCmd)
}

// cmdListProjectRoles returns a list of projectRoles
func cmdListProjectRoles(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.ProjectRoleParams{
		UserID:    projectRoleFilterUserID,
		ProjectID: projectRoleFilterProjectID,
	}

	output := runListCommand(params, aplSvc.ProjectRoles.List)

	if output != nil {
		fields := []string{"ID", "UserID", "ProjectID", "CreatedTime"}
		printTableResultsCustom(output.([]apl.ProjectRole), fields)
	}
}

// cmdGetProjectRoles gets a specified projectRole by projectRole-id
func cmdGetProjectRoles(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.ProjectRoles.Get)

	if output != nil {
		fields := []string{"ID", "UserID", "ProjectID", "CreatedTime"}
		printTableResultsCustom(output.(apl.ProjectRole), fields)
	}
}
