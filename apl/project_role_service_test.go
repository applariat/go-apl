package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

func TestProjectRoleService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.ProjectRoles.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestProjectRoleService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.ProjectRoleParams{
		ProjectID: "e7cda28a-60ff-4cad-a74d-7f5306b3f22f",
	}
	out, _, err := aplSvc.ProjectRoles.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}


func TestProjectRoleService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.ProjectRoles.Get("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

