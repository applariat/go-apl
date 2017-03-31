// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testRoleId     string
	testRoleFilter = "qa"
)

func TestRoleService_List(t *testing.T) {

	out, _, err := aplClient.Roles.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		testRoleId = out[0].ID
	} else {
		t.Fatal("No audit rows found")
	}

	fmt.Printf("Role found %d rows\n", rowCount)

}

func TestRoleService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.RoleParams{
		Role: testRoleFilter,
	}
	out, _, err := aplSvc.Roles.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Role rows found for filter", testRoleFilter)
	}

	fmt.Printf("Role filtered found %d rows for filter \"%s\"\n", rowCount, testRoleFilter)

}

//func TestRoleService_Update(t *testing.T) {
//	aplSvc := apl.NewClient()
//
//	in := &apl.RoleUpdateInput{
//		Workloads: []string{
//			"UPDATED!",
//			"wl-level2-apl",
//		},
//	}
//	out, _, err := aplSvc.Roles.Update(testRoleId, in)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("Updated:", out)
//
//	in = &apl.RoleUpdateInput{
//		Workloads: []string{
//			"wl-level2-apl",
//		},
//	}
//	out, _, err = aplSvc.Roles.Update(testRoleId, in)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("Restored:", out)
//
//}

func TestRoleService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Roles.Get(testRoleId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Role{}) {
		t.Fatal("No Role found for ID", testRoleId)
	}

	fmt.Println("Role found for ID", testRoleId)

}
