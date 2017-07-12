// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"testing"
)

var (
	testStackArtifactId     = "test-stack-artifact-id"
	testStackArtifactFilter = "archive"
)

func TestStackArtifactService_Create(t *testing.T) {

	in := &apl.StackArtifactCreateInput{
		ID:                testStackArtifactId,
		Name:              "Chris Test Zip",
		ArtifactName:      "Chris/chris.zip",
		LocArtifactID:     "la-gs-apl",
		StackID:           "aa409e87-70ef-4977-8588-10a618a1612f",
		Version:           "1.1.1",
		StackArtifactType: "code",
	}

	out, _, err := aplClient.StackArtifacts.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.Data)

}

func TestStackArtifactService_List(t *testing.T) {

	out, _, err := aplClient.StackArtifacts.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No StackArtifact rows found")
	}

	fmt.Printf("StackArtifact found %d rows\n", rowCount)

}

func TestStackArtifactService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.StackArtifactParams{
		Package: testStackArtifactFilter,
	}
	out, _, err := aplSvc.StackArtifacts.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No StackArtifact rows found for filter", testStackArtifactFilter)
	}

	fmt.Printf("StackArtifact filtered found %d rows for filter \"%s\"\n", rowCount, testStackArtifactFilter)

}

func TestStackArtifactService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.StackArtifactUpdateInput{Name: "stack artifact UPDATED!"}
	out, _, err := aplSvc.StackArtifacts.Update(testStackArtifactId, in)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)

}

func TestStackArtifactService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.StackArtifacts.Get(testStackArtifactId)

	if err != nil {
		t.Fatal(err)
	}
	if out == (apl.StackArtifact{}) {
		t.Fatal("No StackArtifact found for ID", testStackArtifactId)
	}

	fmt.Println("StackArtifact found for ID", testStackArtifactId)
}

func TestStackArtifactService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	out, _, err := aplSvc.StackArtifacts.Delete(testStackArtifactId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
