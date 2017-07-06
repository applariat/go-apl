package main

import (
	"fmt"
	"github.com/applariat/roper"
	"github.com/spf13/pflag"
)

// TestStruct represents the data in the file
type TestStruct struct {
	One   int    `json:"one"`
	Two   string `json:"two"`
	Three []struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	} `json:"three"`
}

func main() {

	var inputFile string
	var out = new(TestStruct)

	pflag.StringVarP(&inputFile, "file", "f", "", "Input for create/update: json|yaml|-")
	pflag.Parse()

	err := roper.Unmarshal(inputFile, &out)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", out)
}
