package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

var (
	testLocArtifactId string
	testLocArtifactFilter = "docker"
)

func TestLocArtifactService_Create(t *testing.T) {

	// TODO: Fix TestLocArtifactService_Create!
	t.SkipNow()

	testLocArtifactId = "loc-artifact-test-id"

	aplSvs := apl.NewClient()

	in := &apl.LocArtifactCreateInput{
		ID: testLocArtifactId,
		Name: "LocArtifact Test",
		LocArtifactsType: testLocArtifactFilter,
	}

	out, _, err := aplSvs.LocArtifacts.Create(in)

	if err != nil {
		t.Fatal(err)
	}


	fmt.Println("New LocArtifact ID:", out.Data)

}

func TestLocArtifactService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.LocArtifacts.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("LocArtifact found %d rows\n", rowCount)
	} else {
		t.Fatal("No LocArtifact rows found")
	}

}

func TestLocArtifactService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.LocArtifactParams{
		LocArtifactsType: testLocArtifactFilter,
	}
	out, _, err := aplSvc.LocArtifacts.List(params)

	if err != nil {
		t.Fatal(err)
	}
	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No LocArtifact rows found for filter", testLocArtifactFilter)
	}

	fmt.Printf("LocArtifact filtered found %d rows for filter \"%s\"\n", rowCount, testLocArtifactFilter)


}

func TestLocArtifactService_Update(t *testing.T) {

	// TODO: Fix TestLocArtifactService_Update!
	t.SkipNow()

	aplSvc := apl.NewClient()


	in := &apl.LocArtifactUpdateInput{Name: "name UPDATED!"}
	out, _, err := aplSvc.LocArtifacts.Update(testLocArtifactId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)


}

func TestLocArtifactService_Get(t *testing.T) {

	// TODO: Fix TestLocArtifactService_Get!
	t.SkipNow()

	aplSvc := apl.NewClient()

	out, _, err := aplSvc.LocArtifacts.Get(testLocArtifactId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.LocArtifact{}) {
		t.Fatal("No LocArtifact found for ID", testLocArtifactId)
	}

	fmt.Println("LocArtifact found for ID", testLocArtifactId)

}

func TestLocArtifactService_Delete(t *testing.T) {

	// TODO: Fix TestLocArtifactService_Delete!
	t.SkipNow()

	aplSvc := apl.NewClient()
	out, _, err := aplSvc.LocArtifacts.Delete(testLocArtifactId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
