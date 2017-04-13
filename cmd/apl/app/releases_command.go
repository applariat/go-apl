package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var releaseParams apl.ReleaseParams

func NewReleasesCommand() *cobra.Command {

	cmd := createListCommand(cmdListReleases, "releases", "")
	getCmd := createGetCommand(cmdGetReleases, "release", "")
	createCmd := createCreateCommand(cmdCreateReleases, "release", "")
	deleteCmd := createDeleteCommand(cmdDeleteReleases, "release", "")

	// command flags
	cmd.Flags().StringVar(&releaseParams.Name, "name", "", "Filter releases by name")
	cmd.Flags().StringVar(&releaseParams.Version, "version", "", "Filter releases by version")
	cmd.Flags().StringVar(&releaseParams.StackID, "stack-id", "", "Filter releases by stack_id")
	cmd.Flags().StringVar(&releaseParams.StackVersionID, "stack-version-id", "", "Filter releases by stack_version_id")
	cmd.Flags().StringVar(&releaseParams.ProjectID, "project-id", "", "Filter releases by project_id")
	cmd.Flags().StringVar(&releaseParams.LocImageID, "loc-image-id", "", "Filter releases by loc_image_id")
	cmd.Flags().StringVar(&releaseParams.BuildStatus, "build-status", "", "Filter releases by build_status")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
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

func cmdCreateReleases(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.ReleaseCreateInput{}
	runCreateCommand(in, aplSvs.Releases.Create)
}

func cmdDeleteReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Releases.Delete)
}
