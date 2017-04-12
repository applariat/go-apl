package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	roleFilterName        string
	roleFilterRole        string
	roleFilterAccessLevel int

	rolesCmd       = createListCommand(cmdListRoles, "roles", "")
	rolesGetCmd    = createGetCommand(cmdGetRoles, "role", "")
	rolesUpdateCmd = createUpdateCommand(cmdUpdateRoles, "role", "")
)

func init() {

	// command flags
	rolesCmd.Flags().StringVar(&roleFilterName, "name", "", "Filter roles by name")
	rolesCmd.Flags().StringVar(&roleFilterRole, "role", "", "Filter roles by role")
	rolesCmd.Flags().IntVar(&roleFilterAccessLevel, "access-level", -1, "Filter deployments by access_level")

	// add sub commands
	rolesCmd.AddCommand(rolesGetCmd)
	rolesCmd.AddCommand(rolesUpdateCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(rolesCmd)
}

// cmdListRoles returns a list of roles
func cmdListRoles(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.RoleParams{
		Name: roleFilterName,
		Role: roleFilterRole,
	}

	if roleFilterAccessLevel != -1 {
		params.AccessLevel = roleFilterAccessLevel
	}

	output := runListCommand(params, aplSvc.Roles.List)

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
