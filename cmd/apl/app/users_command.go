package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	userFilterFirstName string
	userFilterLastName  string
	userFilterUserType  string
	userFilterWorkRole  string
	userFilterEmail     string

	usersCmd       = createListCommand(cmdListUsers, "users", "")
	usersGetCmd    = createGetCommand(cmdGetUsers, "user", "")
	usersCreateCmd = createCreateCommand(cmdCreateUsers, "user", "")
	usersUpdateCmd = createUpdateCommand(cmdUpdateUsers, "user", "")
	usersDeleteCmd = createDeleteCommand(cmdDeleteUsers, "user", "")
)

func init() {

	// command flags
	usersCmd.Flags().StringVar(&userFilterFirstName, "first-name", "", "Filter users by first_name")
	usersCmd.Flags().StringVar(&userFilterLastName, "last-name", "", "Filter users by last_name")
	usersCmd.Flags().StringVar(&userFilterUserType, "user-type", "", "Filter users by user_type")
	usersCmd.Flags().StringVar(&userFilterWorkRole, "work-role", "", "Filter users by work_role")
	usersCmd.Flags().StringVar(&userFilterEmail, "email", "", "Filter users by email")

	// add sub commands
	usersCmd.AddCommand(usersGetCmd)
	usersCmd.AddCommand(usersCreateCmd)
	usersCmd.AddCommand(usersUpdateCmd)
	usersCmd.AddCommand(usersDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(usersCmd)
}

// cmdListUsers returns a list of users
func cmdListUsers(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.UserParams{
		FirstName: userFilterFirstName,
		LastName:  userFilterLastName,
		UserType:  userFilterUserType,
		WorkRole:  userFilterWorkRole,
		Email:     userFilterEmail,
	}

	output := runListCommand(params, aplSvc.Users.List)

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
