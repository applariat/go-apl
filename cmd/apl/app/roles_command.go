package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var roleParams apl.RoleParams

func NewRolesCommand() *cobra.Command {

	cmd := createListCommand(cmdListRoles, "roles", "")
	getCmd := createGetCommand(cmdGetRoles, "role", "")
	updateCmd := createUpdateCommand(cmdUpdateRoles, "role", "")

	// command flags
	cmd.Flags().StringVar(&roleParams.Name, "name", "", "Filter roles by name")
	cmd.Flags().StringVar(&roleParams.Role, "role", "", "Filter roles by role")
	//cmd.Flags().IntVar(&roleParams.AccessLevel, "access-level", -1, "Filter deployments by access_level")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(updateCmd)

	return cmd
}

// cmdListRoles returns a list of roles
func cmdListRoles(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&roleParams, aplSvc.Roles.List)

	if output != nil {
		fields := []string{"ID", "Name", "Role", "AccessLevel", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Role), fields)
	}
}

// cmdGetRoles gets a specified role by role-id
func cmdGetRoles(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Roles.Get)

	if output != nil {
		fields := []string{"ID", "Name", "Role", "AccessLevel", "CreatedTime"}
		printTableResultsCustom(output.(apl.Role), fields)
	}
}

func cmdUpdateRoles(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.RoleUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Roles.Update)
}
