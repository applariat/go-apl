package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var projectRoleParams apl.ProjectRoleParams

func NewProjectRolesCommand() *cobra.Command {

	cmd := createListCommand(cmdListProjectRoles, "project-roles", "")
	getCmd := createGetCommand(cmdGetProjectRoles, "project-role", "")

	// command flags
	cmd.Flags().StringVar(&projectRoleParams.UserID, "user-id", "", "Filter project-roles by user_id")
	cmd.Flags().StringVar(&projectRoleParams.ProjectID, "project-id", "", "Filter project-roles by project_id")

	// add sub commands
	cmd.AddCommand(getCmd)

	return cmd
}

// cmdListProjectRoles returns a list of projectRoles
func cmdListProjectRoles(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&projectRoleParams, aplSvc.ProjectRoles.List)

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
