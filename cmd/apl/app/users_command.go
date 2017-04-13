package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var userParams apl.UserParams

func NewUsersCommand() *cobra.Command {
	cmd := createListCommand(cmdListUsers, "users", "")
	getCmd := createGetCommand(cmdGetUsers, "user", "")
	createCmd := createCreateCommand(cmdCreateUsers, "user", "")
	updateCmd := createUpdateCommand(cmdUpdateUsers, "user", "")
	deleteCmd := createDeleteCommand(cmdDeleteUsers, "user", "")

	// command flags
	cmd.Flags().StringVar(&userParams.FirstName, "first-name", "", "Filter users by first_name")
	cmd.Flags().StringVar(&userParams.LastName, "last-name", "", "Filter users by last_name")
	cmd.Flags().StringVar(&userParams.UserType, "user-type", "", "Filter users by user_type")
	cmd.Flags().StringVar(&userParams.WorkRole, "work-role", "", "Filter users by work_role")
	cmd.Flags().StringVar(&userParams.Email, "email", "", "Filter users by email")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListUsers returns a list of users
func cmdListUsers(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(&userParams, aplSvc.Users.List)
	if output != nil {
		fields := []string{"ID", "Email", "FirstName", "LastName", "UserType", "WorkRole", "CreatedTime"}
		printTableResultsCustom(output.([]apl.User), fields)
	}
}

// cmdGetUsers gets a specified user by user-id
func cmdGetUsers(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runGetCommand(args, aplSvc.Users.Get)
	if output != nil {
		fields := []string{"ID", "Email", "FirstName", "LastName", "UserType", "WorkRole", "CreatedTime"}
		printTableResultsCustom(output.(apl.User), fields)
	}
}

func cmdCreateUsers(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.UserCreateInput{}
	runCreateCommand(in, aplSvs.Users.Create)
}

func cmdUpdateUsers(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.UserUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Users.Update)
}

func cmdDeleteUsers(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Users.Delete)
}
