package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Global flags. For more flags see flags.go
	printerType string

	// APLCmd ...
	AppLariatCmd = &cobra.Command{
		Use:   "apl",
		Short: "apl",
		Long:  `The appLariat (apl) Command Line Interface is a unified tool to manage your appLariat service. You can control all appLariat services from the command line and automate them through scripts.`,
		PersistentPreRunE: func(ccmd *cobra.Command, args []string) error {

			err := checkPrinterType()
			if err != nil {
				return err
			}

			err = checkInputFileExists()
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {

	// persistent flags, globals
	AppLariatCmd.PersistentFlags().StringVarP(&printerType, "output", "o", "table", "Output format: json|yaml")

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
