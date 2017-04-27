package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/applariat/go-apl/pkg/apl"
	"github.com/fatih/camelcase"
	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"reflect"
	"strings"
)

// Determine if the code should print out a table for itself
// or pass the output to the results printer for generic our json|yaml
func shouldPrintTable() bool {
	return printerType == "table"
}

// checkPrinterType makes sure the printer type is correct
func checkPrinterType() error {
	if printerType != "" {
		if printerType != "yaml" && printerType != "json" && printerType != "table" {
			return fmt.Errorf("Invalid output type: %s", printerType)
		}
	}
	return nil
}

// printResults formats and prints the results based on the output flag
func printResults(data interface{}) {

	switch printerType {

	case "json":
		if data == nil {
			fmt.Println("{}")
			return
		}
		j, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(j))

	case "yaml":
		if data == nil {
			fmt.Println("---")
			return
		}
		y, err := yaml.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(y))

	default:
		if data == nil {
			fmt.Println("No data found")
			return
		}

		// Print app.CLIError
		if cliError, ok := data.(CLIError); ok {
			data := [][]string{
				[]string{cliError.Message},
			}
			header := []string{"CLI Error"}
			printTableResults(data, header)
			return
		}

		// Print apl.APLError
		if aplError, ok := data.(apl.APIError); ok {
			data := [][]string{
				[]string{strconv.Itoa(aplError.StatusCode), aplError.Message},
			}
			header := []string{"Status Code", "Message"}
			printTableResults(data, header)
			return
		}

		// Print apl.CreateResult
		if createResult, ok := data.(apl.CreateResult); ok {
			result := createResult.Data.(string)
			data := [][]string{[]string{result}}
			header := []string{"New ID"}
			printTableResults(data, header)
			return
		}

		// Print apl.ModifyOutput
		if modifyResult, ok := data.(apl.ModifyResult); ok {
			data := [][]string{
				[]string{
					strconv.Itoa(modifyResult.Skipped),
					strconv.Itoa(modifyResult.Deleted),
					strconv.Itoa(modifyResult.Unchanged),
					strconv.Itoa(modifyResult.Errors),
					strconv.Itoa(modifyResult.Replaced),
					strconv.Itoa(modifyResult.Inserted),
				},
			}
			header := []string{"Skipped", "Deleted", "Unchanged", "Errors", "Replaced", "Inserted"}
			printTableResults(data, header)
			return
		}

	}
}

func printTableResults(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}

	table.Render() // Send output
}

// prints out a table
func printTableResultsCustom(results interface{}, fields []string) {

	if results != nil {
		header := make([]string, len(fields))
		output := reflect.ValueOf(results)

		switch reflect.TypeOf(results).Kind() {

		case reflect.Struct:

			dataRow := make([]string, len(fields))
			for i := 0; i < len(fields); i++ {
				key, val := getSubField(fields[i], output)
				header[i] = camelCaseToSpaces(key)
				dataRow[i] = fmt.Sprint(val)
			}
			data := [][]string{dataRow}

			printTableResults(data, header)

		case reflect.Slice:

			data := make([][]string, output.Len())

			for i := 0; i < output.Len(); i++ {
				outputRow := output.Index(i)
				dataRow := make([]string, len(fields))

				for j := 0; j < len(fields); j++ {
					key, val := getSubField(fields[j], outputRow)
					header[j] = camelCaseToSpaces(key)
					dataRow[j] = fmt.Sprint(val)
				}
				data = append(data, dataRow)
			}
			printTableResults(data, header)
		}

	}

}

// split the name up by dot notation and recurse to find the field in question
// Warning, this is recursive
func getSubField(key string, item reflect.Value) (string, reflect.Value) {

	sub := strings.SplitN(key, ".", 2)

	if len(sub) > 1 {
		nextKey := sub[1]
		return getSubField(nextKey, item)
	}

	return key, item.FieldByName(key)
}

// camelCaseToSpaces takes a slice of CamelCaseNames and puts spaces in them Camel Case Names
func camelCaseToSpaces(input string) string {
	return strings.Join(camelcase.Split(input), " ")
}

//func camelCaseSliceToSpaces(input []string) []string {
//	result := make([]string, len(input))
//	for idx, val := range input {
//		result[idx] = camelCaseToSpaces(val)
//	}
//	return result
//}
