// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testProjectRoleId     string
	testProjectRoleFilter string
)

func TestProjectRoleService_List(t *testing.T) {

	out, _, err := aplClient.ProjectRoles.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		testProjectRoleId = out[0].ID
		testProjectRoleFilter = out[0].ProjectID
	} else {
		t.Fatal("No audit rows found")
	}

	fmt.Printf("ProjectRole found %d rows\n", rowCount)

}

func TestProjectRoleService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.ProjectRoleParams{
		ProjectID: testProjectRoleFilter,
	}
	out, _, err := aplSvc.ProjectRoles.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No ProjectRole rows found for filter", testProjectRoleFilter)
	}

	fmt.Printf("ProjectRole filtered found %d rows for filter \"%s\"\n", rowCount, testProjectRoleFilter)

}

func TestProjectRoleService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.ProjectRoles.Get(testProjectRoleId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.ProjectRole{}) {
		t.Fatal("No ProjectRole found for ID", testProjectRoleId)
	}

	fmt.Println("ProjectRole found for ID", testProjectRoleId)

}
