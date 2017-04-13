package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var eventParams apl.EventParams

func NewEventsCommand() *cobra.Command {

	cmd := createListCommand(cmdListEvents, "events", "")
	getCmd := createGetCommand(cmdGetEvents, "event", "")
	createCmd := createCreateCommand(cmdCreateEvents, "event", "")

	// command flags
	cmd.Flags().StringVar(&eventParams.ObjectType, "object-type", "", "Filter events by object_type")
	cmd.Flags().StringVar(&eventParams.Source, "source", "", "Filter events by source")
	cmd.Flags().StringVar(&eventParams.ObjectName, "object-name", "", "Filter events by object_name")
	//cmd.Flags().StringVar(&eventParams.Active, "active", "", "Filter events by active")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)

	return cmd
}

// cmdListEvents returns a list of events
func cmdListEvents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&eventParams, aplSvc.Events.List)

	if output != nil {
		fields := []string{"ID", "ObjectName", "ObjectType", "Active"}
		printTableResultsCustom(output.([]apl.Event), fields)
	}
}

// cmdGetEvents gets a specified event by event-id
func cmdGetEvents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Events.Get)

	if output != nil {
		fields := []string{"ID", "ObjectName", "ObjectType", "Active"}
		printTableResultsCustom(output.(apl.Event), fields)
	}
}

func cmdCreateEvents(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.EventCreateInput{}
	runCreateCommand(in, aplSvs.Events.Create)
}
