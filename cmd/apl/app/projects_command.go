package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	projectFilterName string

	projectsCmd       = createListCommand(cmdListProjects, "projects", "")
	projectsGetCmd    = createGetCommand(cmdGetProjects, "project", "")
	projectsCreateCmd = createCreateCommand(cmdCreateProjects, "project", "")
	projectsUpdateCmd = createUpdateCommand(cmdUpdateProjects, "project", "")
	projectsDeleteCmd = createDeleteCommand(cmdDeleteProjects, "project", "")
)

func init() {

	// command flags
	projectsCmd.Flags().StringVar(&projectFilterName, "name", "", "Filter projects by name")

	// add sub commands
	projectsCmd.AddCommand(projectsGetCmd)
	projectsCmd.AddCommand(projectsCreateCmd)
	projectsCmd.AddCommand(projectsUpdateCmd)
	projectsCmd.AddCommand(projectsDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(projectsCmd)
}

// cmdListProjects returns a list of projects
func cmdListProjects(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.ProjectParams{
		Name: projectFilterName,
	}

	output := runListCommand(params, aplSvc.Projects.List)

	if output != nil {
		fields := []string{"ID", "Name", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Project), fields)
	}
}

// cmdGetProjects gets a specified project by project-id
func cmdGetProjects(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Projects.Get)

	if output != nil {
		fields := []string{"ID", "Name", "CreatedTime"}
		printTableResultsCustom(output.(apl.Project), fields)
	}
}

func cmdCreateProjects(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.ProjectCreateInput{}
	runCreateCommand(in, aplSvs.Projects.Create)
}

func cmdUpdateProjects(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.ProjectUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Projects.Update)
}

func cmdDeleteProjects(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Projects.Delete)
}
