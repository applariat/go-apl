package app

import (
	"fmt"
	"strconv"
	"strings"
)

// Implement the flag.Value interface
// example flag --component "StackComponentID=my-scID,ServiceName=my-service,StackArtifactID=1,StackArtifactID=2,Instances=1"
// You can pass in as many --component flags as you want.

const (
	mapStackComponentID = "StackComponentID"
	mapServiceName      = "ServiceName"
	mapStackArtifactID  = "StackArtifactID"
	mapInstances        = "Instances"
)

type ComponentArgs struct {
	OriginalValue    string
	StackComponentID string
	ServiceName      string
	StackArtifactIDs []string
	Instances        int
}

type ComponentStringMap struct {
	Values []ComponentArgs
}

// Set takes in the key=value pairs, validates them and then assigns to struct value
func (s *ComponentStringMap) Set(value string) error {

	dc := ComponentArgs{}
	dc.OriginalValue = value

	// Split on comma to get key=value pairs
	for _, arg := range strings.Split(value, ",") {

		// split to key=value
		x := strings.Split(arg, "=")
		if len(x) < 2 {
			return fmt.Errorf("Incorrect key=value format for %s", arg)
		}

		key := x[0]
		value := x[1]

		if value == "" {
			return fmt.Errorf("value empty for key %s", key)
		}

		switch key {
		case mapStackComponentID:
			dc.StackComponentID = value
		case mapServiceName:
			dc.ServiceName = value
		case mapStackArtifactID:
			dc.StackArtifactIDs = append(dc.StackArtifactIDs, value)
		case mapInstances:
			// Sensible default here
			dc.Instances = 1
			instances, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("%s must be a number", key)
			}
			dc.Instances = instances
		default:
			return fmt.Errorf("Unsupported key %s", key)
		}
	}
	s.Values = append(s.Values, dc)

	return nil
}

func (s *ComponentStringMap) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *ComponentStringMap) Type() string {
	return fmt.Sprintf("%T", s)
}

func (s *ComponentStringMap) Usage() string {
	return "\"StackComponentID=my-scID,ServiceName=my-service,StackArtifactID=1,StackArtifactID=2,Instances=1\""
}
