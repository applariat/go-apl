package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var projectParams apl.ProjectParams

// NewProjectsCommand Creates a cobra command for Projects
func NewProjectsCommand() *cobra.Command {

	cmd := createListCommand(cmdListProjects, "projects", "")
	getCmd := createGetCommand(cmdGetProjects, "project", "")
	createCmd := createCreateCommand(cmdCreateProjects, "project", "")
	updateCmd := createUpdateCommand(cmdUpdateProjects, "project", "")
	deleteCmd := createDeleteCommand(cmdDeleteProjects, "project", "")

	// command flags
	cmd.Flags().StringVar(&projectParams.Name, "name", "", "Filter projects by name")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListProjects returns a list of projects
func cmdListProjects(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&projectParams, aplSvc.Projects.List)

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
