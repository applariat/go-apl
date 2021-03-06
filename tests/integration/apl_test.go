// Package tests contains integration tests.
//
// These tests call the live appLariat API and should be run by hand with care.
// They require env vars or config file
// +build integration

package tests

import (
	"github.com/applariat/go-apl/pkg/apl"
)

var (
	aplClient *apl.Client
)

func init() {
	aplClient = apl.NewClient()
}
