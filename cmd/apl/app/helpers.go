package app

import (
	"fmt"
	"github.com/applariat/roper"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"reflect"
	"time"
)

// All of the commands are created and ran the same way.
// The helpers do the work in a consistent fashion.

// createListCommand returns a standard list command
func createListCommand(cb func(ccmd *cobra.Command, args []string), label string, desc string) *cobra.Command {
	ccmd := &cobra.Command{
		Use:   label,
		Short: fmt.Sprintf("manage %s", label),
		Long:  desc,
		Run:   cb,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkListCommandHasNoArgs(args)
		},
	}
	return ccmd
}

// createGetCommand returns a standard get command
func createGetCommand(cb func(ccmd *cobra.Command, args []string), label string, desc string) *cobra.Command {
	return &cobra.Command{
		Use:   "get [ID]",
		Short: fmt.Sprintf("get a %s", label),
		Long:  desc,
		Run:   cb,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkCommandHasIDInArgs(args, label)
		},
	}
}

// createCreateCommand returns a standard post/create command
func createCreateCommand(cb func(ccmd *cobra.Command, args []string), label string, desc string) *cobra.Command {
	ccmd := &cobra.Command{
		Use:   "create [-f inputFile.yaml]",
		Short: fmt.Sprintf("create a %s", label),
		Long:  desc,
		Run:   cb,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkInputFileDefined()
		},
	}
	addInputFileFlag(ccmd)
	return ccmd
}

// createUpdateCommand returns a standard put/update command
func createUpdateCommand(cb func(ccmd *cobra.Command, args []string), label string, desc string) *cobra.Command {
	ccmd := &cobra.Command{
		Use:   "update [-f inputFile.yaml]",
		Short: fmt.Sprintf("update a %s", label),
		Long:  desc,
		Run:   cb,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := checkCommandHasIDInArgs(args, "credential")
			if err != nil {
				return err
			}

			err = checkInputFileDefined()
			if err != nil {
				return err
			}

			return nil
		},
	}
	addInputFileFlag(ccmd)
	return ccmd

}

// createDeleteCommand returns a standard delete command
func createDeleteCommand(cb func(ccmd *cobra.Command, args []string), label string, desc string) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [ID]",
		Short: fmt.Sprintf("delete a %s", label),
		Long:  desc,
		Run:   cb,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkCommandHasIDInArgs(args, "credential")
		},
	}
}

// Helper to run the standard list command
func runListCommand(params interface{}, callMe interface{}) interface{} {

	in := []reflect.Value{}
	if params != nil {
		in = []reflect.Value{reflect.ValueOf(params)}
	}

	response := reflect.ValueOf(callMe).Call(in)
	output, _, err := getReflectOutput(response)
	if err != nil {
		printResults(err)
		return nil
	}

	// check for empty results
	if reflect.ValueOf(output).Kind() == reflect.Slice && reflect.ValueOf(output).Len() == 0 {
		printResults(nil)
		return nil
	}

	// We return output if it needs to be printed as a table
	if shouldPrintTable() {
		return output
	}

	printResults(output)

	return nil
}

// Helper to run the standard list command
func runGetCommand(args []string, callMe interface{}) interface{} {

	var id string
	if len(args) != 0 {
		id = args[0]
	}

	in := []reflect.Value{reflect.ValueOf(id)}

	response := reflect.ValueOf(callMe).Call(in)
	output, _, err := getReflectOutput(response)
	if err != nil {
		printResults(err)
		return nil
	}

	// check for empty results
	if reflect.ValueOf(output).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(output)).Elem().Interface()
		if reflect.DeepEqual(output, empty) {
			printResults(nil)
			return nil
		}
	}

	// We return output if it needs to be printed as a table
	if shouldPrintTable() {
		return output
	}

	printResults(output)

	return nil
}

// Helper to run the standard create command
func runCreateCommand(input interface{}, callMe interface{}) {

	err := processInputFile(input)
	if err != nil {
		printResults(NewCLIError(err.Error()))
		return
	}

	in := []reflect.Value{}
	if input != nil {
		in = []reflect.Value{reflect.ValueOf(input)}
	}

	response := reflect.ValueOf(callMe).Call(in)
	output, _, err := getReflectOutput(response)
	if err != nil {
		printResults(err)
		os.Exit(1)
	}

	printResults(output)
}

// Helper to run the standard update command
func runUpdateCommand(args []string, input interface{}, callMe interface{}) {

	err := roper.Unmarshal(inputFile, input)
	if err != nil {
		printResults(NewCLIError(err.Error()))
		return
	}

	var id string
	if len(args) != 0 {
		id = args[0]
	}

	in := []reflect.Value{}
	if input != nil {
		in = []reflect.Value{
			reflect.ValueOf(id),
			reflect.ValueOf(input),
		}
	}

	response := reflect.ValueOf(callMe).Call(in)
	output, _, err := getReflectOutput(response)
	if err != nil {
		printResults(err)
		os.Exit(1)
	}

	printResults(output)
}

// Helper to run the standard delete command
func runDeleteCommand(args []string, callMe interface{}) {

	var id string
	if len(args) != 0 {
		id = args[0]
	}

	in := []reflect.Value{reflect.ValueOf(id)}

	response := reflect.ValueOf(callMe).Call(in)
	output, _, err := getReflectOutput(response)
	if err != nil {
		printResults(err)
		os.Exit(1)
	}

	printResults(output)

}

func isInputFileDefined() bool {
	return inputFile != ""
}

func checkInputFileDefined() error {
	if !isInputFileDefined() {
		return fmt.Errorf("Input for create/update: json|yaml|-")
	}
	return nil
}

func addInputFileFlag(ccmd *cobra.Command) {
	ccmd.PersistentFlags().StringVarP(&inputFile, "file", "f", "", "Input for create/update: json|yaml|-")
}

func processInputFile(input interface{}) error {
	// We check to see if inputFile flag is populated
	if isInputFileDefined() {
		return roper.Unmarshal(inputFile, input)
	}
	return nil
}

// Helper to cast response
func getReflectOutput(response []reflect.Value) (interface{}, *http.Response, error) {

	output := response[0].Interface()
	resp := response[1].Interface().(*http.Response)

	var err error
	if !response[2].IsNil() {
		err = response[2].Interface().(error)
	}

	return output, resp, err
}

// check arguments for only one string. Used for get/update/delete sub commands
func checkCommandHasIDInArgs(args []string, label string) error {
	l := len(args)
	if l == 0 {
		return fmt.Errorf("Missing %s id.", label)
	}
	if l > 1 {
		return fmt.Errorf("Too many arguments provided.")
	}
	if args[0] == "" {
		return fmt.Errorf("Missing %s id.", label)
	}
	return nil
}

// checkListCommandHasNoArgs makes sure there are no args for lists!
func checkListCommandHasNoArgs(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("Non flag arguments are not permitted in list mode.")
	}
	return nil
}

func timeDuration(d time.Duration) string {

	s := int(d.Seconds())

	if s < -1 {
		return fmt.Sprintf("??")
	} else if s < 0 {
		return fmt.Sprintf("0s")
	} else if s < 60 {
		return fmt.Sprintf("%ds", s)
	}

	m := int(d.Minutes())
	if m < 60 {
		return fmt.Sprintf("%dm", m)
	}

	h := int(d.Hours())
	if h < 24 {
		return fmt.Sprintf("%dh", h)
	} else if h < 24*365 {
		return fmt.Sprintf("%dd", h/24)
	}

	return fmt.Sprintf("%dy", int(d.Hours()/24/365))
}
