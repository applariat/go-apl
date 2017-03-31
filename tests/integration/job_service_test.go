// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testJobId     string
	testJobFilter string
)

func TestJobService_List(t *testing.T) {

	out, _, err := aplClient.Jobs.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		testJobId = out[0].ID
		testJobFilter = out[0].ResourceID
	} else {
		t.Fatal("No Job rows found")
	}

	fmt.Printf("Job found %d rows\n", rowCount)

}

func TestJobService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.JobParams{
		ResourceID: testJobFilter,
	}
	out, _, err := aplSvc.Jobs.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Job rows found for filter", testJobFilter)
	}

	fmt.Printf("Job filtered found %d rows for filter \"%s\"\n", rowCount, testJobFilter)

}

func TestJobService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Jobs.Get(testJobId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Job{}) {
		t.Fatal("No Job found for ID", testJobId)
	}

	fmt.Println("Job found for ID", testJobId)
}
