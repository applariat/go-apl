// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"testing"
)

var (
	testProjectId     = "project-test-id"
	testProjectFilter = "project-test-name"
)

func TestProjectService_Create(t *testing.T) {

	in := &apl.ProjectCreateInput{
		ID:       testProjectId,
		Name:     testProjectFilter,
		Settings: map[string]int{"settings": 1},
		//Users: "[]",
	}

	out, _, err := aplClient.Projects.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	data, _ := out.Data.(map[string]string)
	fmt.Println("New Project ID:", data["project_id"])

}

func TestProjectService_List(t *testing.T) {

	out, _, err := aplClient.Projects.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("Project found %d rows\n", rowCount)
	} else {
		t.Fatal("No Project rows found")
	}

}

func TestProjectService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.ProjectParams{
		Name: testProjectFilter,
	}
	out, _, err := aplSvc.Projects.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Project rows found for filter", testProjectFilter)
	}

	fmt.Printf("Project filtered found %d rows for filter \"%s\"\n", rowCount, testProjectFilter)

}

func TestProjectService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.ProjectUpdateInput{Name: "name UPDATED!"}
	out, _, err := aplSvc.Projects.Update(testProjectId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}

func TestProjectService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Projects.Get(testProjectId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Project{}) {
		t.Fatal("No Project found for ID", testProjectId)
	}

	fmt.Println("Project found for ID", testProjectId)

}

func TestProjectService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	//testProjectId = "project-test-id8"
	out, _, err := aplSvc.Projects.Delete(testProjectId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
