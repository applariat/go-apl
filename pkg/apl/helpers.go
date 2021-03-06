package apl

import (
	"encoding/json"
	"github.com/dghubble/sling"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// Helper function for list
func doList(sling *sling.Sling, path string, params interface{}, output interface{}) (*http.Response, error) {
	apiError := new(WrappedAPIError)
	resp, err := sling.New().Get(path).QueryStruct(params).Receive(output, apiError)

	return resp, relevantError(err, apiError)
}

// Helper function for get
func doGet(sling *sling.Sling, path string, output interface{}) (*http.Response, error) {
	apiError := new(WrappedAPIError)
	if strings.HasSuffix(path, "/") {
		apiError.Message = "ID not provided or is empty"
		apiError.StatusCode = 400
		return nil, apiError
	}
	resp, err := sling.New().Get(path).Receive(output, apiError)
	return resp, relevantError(err, apiError)
}

// Helper function for create
func doCreate(sling *sling.Sling, path string, input interface{}) (CreateResult, *http.Response, error) {
	output := CreateResult{}
	apiError := new(WrappedAPIError)

	body := &CreateInput{Data: input}
	resp, err := sling.New().Post(path).BodyJSON(body).Receive(&output, apiError)
	return output, resp, relevantError(err, apiError)
}

// Helper function for update
func doUpdate(sling *sling.Sling, path string, input interface{}) (ModifyResult, *http.Response, error) {
	output := ModifyOutput{}
	apiError := new(WrappedAPIError)

	body := &CreateInput{Data: input}
	resp, err := sling.New().Put(path).BodyJSON(body).Receive(&output, apiError)
	return output.ModifyResult, resp, relevantError(err, apiError)
}

// Helper function for update
func doDelete(sling *sling.Sling, path string) (ModifyResult, *http.Response, error) {
	output := ModifyOutput{}
	apiError := new(WrappedAPIError)

	if strings.HasSuffix(path, "/") {
		apiError.Message = "ID not provided or is empty"
		apiError.StatusCode = 400
		return ModifyResult{Errors: 1}, nil, apiError
	}

	resp, err := sling.New().Delete(path).Receive(&output, apiError)
	return output.ModifyResult, resp, relevantError(err, apiError)
}

// Helper function for update
func doEdit(sling *sling.Sling, path string) (ModifyResult, *http.Response, error) {
	output := ModifyOutput{}
	apiError := new(WrappedAPIError)

	if strings.HasSuffix(path, "/") {
		apiError.Message = "ID not provided or is empty"
		apiError.StatusCode = 400
		return ModifyResult{Errors: 1}, nil, apiError
	}

	resp, err := sling.New().Delete(path).Receive(&output, apiError)
	return output.ModifyResult, resp, relevantError(err, apiError)
}

// editFile creates a file
func editFile(fileData interface{}) ([]byte, error) {

	rawData, err := json.MarshalIndent(fileData, "", "  ")
	if err != nil {
		return rawData, err
	}

	fileName := "/tmp/testedit.yaml"
	ioutil.WriteFile(fileName, rawData, 0644)
	cmd := exec.Command("vi", fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return rawData, err
	}
	return rawData, nil
}
