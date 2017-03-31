// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testReleaseId     = "release-test-id"
	testReleaseFilter = 1
)

func TestReleaseService_Create(t *testing.T) {

	// TODO: Fix TestReleaseService_Create!
	t.SkipNow()

	in := &apl.ReleaseCreateInput{
		ID:             testReleaseId,
		Version:        testReleaseFilter,
		StackID:        "",
		StackVersionID: "",
		ProjectID:      "",
		LocImageID:     "",
		BuildStatus:    "",
		Components:     "[]",
	}

	out, _, err := aplClient.Releases.Create(in)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("New Release ID:", out.Data)

}

func TestReleaseService_List(t *testing.T) {

	out, _, err := aplClient.Releases.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("Release found %d rows\n", rowCount)
	} else {
		t.Fatal("No Release rows found")
	}

}

func TestReleaseService_ListByType(t *testing.T) {

	// TODO: Fix TestReleaseService_ListByType!
	t.SkipNow()

	aplSvc := apl.NewClient()

	params := &apl.ReleaseParams{
		Version: "1",
	}
	out, _, err := aplSvc.Releases.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Release rows found for filter", testReleaseFilter)
	}

	fmt.Printf("Release filtered found %d rows for filter \"%s\"\n", rowCount, testReleaseFilter)

}

func TestReleaseService_Get(t *testing.T) {

	// TODO: Fix TestReleaseService_Get!
	t.SkipNow()

	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Releases.Get(testReleaseId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Release{}) {
		t.Fatal("No Release found for ID", testReleaseId)
	}

	fmt.Println("Release found for ID", testReleaseId)

}

func TestReleaseService_Delete(t *testing.T) {

	// TODO: Fix TestReleaseService_Delete!
	t.SkipNow()

	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Releases.Delete(testReleaseId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
