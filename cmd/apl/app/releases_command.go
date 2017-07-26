package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var releaseParams apl.ReleaseParams

// NewReleasesCommand Creates a cobra command for Releases
func NewReleasesCommand() *cobra.Command {

	cmd := createListCommand(cmdListReleases, "releases", "")
	cmd.Flags().StringVar(&releaseParams.Name, "name", "", "Filter releases by name")
	cmd.Flags().StringVar(&releaseParams.Version, "version", "", "Filter releases by version")
	cmd.Flags().StringVar(&releaseParams.StackID, "stack-id", "", "Filter releases by stack_id")
	cmd.Flags().StringVar(&releaseParams.StackVersionID, "stack-version-id", "", "Filter releases by stack_version_id")
	cmd.Flags().StringVar(&releaseParams.ProjectID, "project-id", "", "Filter releases by project_id")
	cmd.Flags().StringVar(&releaseParams.LocImageID, "loc-image-id", "", "Filter releases by loc_image_id")
	cmd.Flags().StringVar(&releaseParams.BuildStatus, "build-status", "", "Filter releases by build_status")

	// Get
	getCmd := createGetCommand(cmdGetReleases, "release", "")
	cmd.AddCommand(getCmd)

	// Create
	createCmd := NewReleasesCreateCommand()
	cmd.AddCommand(createCmd)

	// Delete
	deleteCmd := createDeleteCommand(cmdDeleteReleases, "release", "")
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListReleases returns a list of releases
func cmdListReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(&releaseParams, aplSvc.Releases.List)
	if output != nil {
		fields := []string{"ID", "StackID", "Version", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Release), fields)
	}
}

// cmdGetReleases gets a specified release by release-id
func cmdGetReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runGetCommand(args, aplSvc.Releases.Get)
	if output != nil {
		fields := []string{"ID", "StackID", "Version", "CreatedTime"}
		printTableResultsCustom(output.(apl.Release), fields)
	}
}

func cmdDeleteReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Releases.Delete)
}
