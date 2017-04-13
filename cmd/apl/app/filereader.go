package app

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
)

var inputFile string

// inputFileToStruct will read in a json or yaml file and Unmarshal it.
// this function requires the file extension to be .json or .yaml
func inputFileToStruct(v interface{}) error {

	if inputFile != "" {

		raw, err := ioutil.ReadFile(inputFile)

		if err != nil {
			return err
		}

		ext := path.Ext(inputFile)
		if ext == ".yaml" {
			err = yaml.Unmarshal(raw, &v)
			if err != nil {
				return err
			}
			return nil
		}

		err = json.Unmarshal(raw, &v)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Problem reading inputfile: %s", inputFile)
}

func addInputFileFlag(ccmd *cobra.Command) {
	ccmd.PersistentFlags().StringVarP(&inputFile, "file", "f", "", "Input file for create/update: json|yaml")
}

// check to see if string inputFile is defined
func checkInputFileDefined() error {
	if inputFile == "" {
		return fmt.Errorf("Missing input file --file yaml|json")
	}
	return nil
}

func checkInputFileExists() error {
	if inputFile != "" {
		ext := path.Ext(inputFile)
		if ext != ".json" && ext != ".yaml" {
			return fmt.Errorf("Input file must be .json or .yaml")
		}
		_, err := os.Stat(inputFile)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("Input file does not exist: %s", inputFile)
			}
		}
	}
	return nil
}
