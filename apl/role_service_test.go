package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

func TestRoleService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Roles.List(nil)

	if err != nil {
		t.Fatal(err)
	}
	for _, item := range out {
		//fmt.Println("Name:", item.Name)
		fmt.Println(item)
	}

}

func TestRoleService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.RoleParams{
		Role: "qa",
	}
	out, _, err := aplSvc.Roles.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}

func TestRoleService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.RoleUpdateInput{
		Workloads: []string{
			"UPDATED!",
			"wl-level2-apl",
		},
	}
	out, _, err := aplSvc.Roles.Update("43adc951-edd6-4dc7-b3d5-0eb30f91b6d2", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated:", out)

	in = &apl.RoleUpdateInput{
		Workloads: []string{
			"wl-level2-apl",
		},
	}
	out, _, err = aplSvc.Roles.Update("43adc951-edd6-4dc7-b3d5-0eb30f91b6d2", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Restored:", out)

}

func TestRoleService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Roles.Get("43adc951-edd6-4dc7-b3d5-0eb30f91b6d2")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

