// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"testing"
)

var (
	testComponentId     string
	testComponentFilter string
)

func TestComponentService_List(t *testing.T) {

	out, _, err := aplClient.Components.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		testComponentId = out[0].ID
		testComponentFilter = out[0].Category
	} else {
		t.Fatal("No Component rows found")
	}

	fmt.Printf("Component found %d rows\n", rowCount)

}

func TestComponentService_ListByType(t *testing.T) {

	params := &apl.ComponentParams{
		Category: testComponentFilter,
	}
	out, _, err := aplClient.Components.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Component rows found for filter", testComponentFilter)
	}

	fmt.Printf("Component filtered found %d rows for filter \"%s\"\n", rowCount, testComponentFilter)

}

func TestComponentService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Components.Get(testComponentId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Component{}) {
		t.Fatal("No Component found for ID", testComponentId)
	}

	fmt.Println("Component found for ID", testComponentId)

}
