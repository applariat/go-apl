package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

func TestDeploymentService_Create(t *testing.T) {
	aplSvs := apl.NewClient()

	in := &apl.DeploymentCreateInput{
		ID:             "chris-test-id",
		Name: "Chris Test",
	}

	out, _, err := aplSvs.Deployments.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.PrimaryKey)

}

func TestDeploymentService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Deployments.List(nil)

	if err != nil {
		t.Fatal(err)
	}
	for _, item := range out {
		//fmt.Println("Name:", item.Name)
		fmt.Println(item.Name)
	}

}

func TestDeploymentService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.DeploymentParams{
		//LeaseType: "temporary",
		//ProjectID: "p-mobile-apps-apl",
		Name: "mobile-app-r3-to-gke",
	}
	out, _, err := aplSvc.Deployments.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}

func TestDeploymentService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.DeploymentUpdateInput{Name: "stack artifact UPDATED!"}
	out, _, err := aplSvc.Deployments.Update("chris-test-id", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated:", out)

}

func TestDeploymentService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Deployments.Get("3d9c7ac3-7b87-4b4c-ba5d-750a91629d05")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Deployment ID:", out.ID)
}

func TestDeploymentService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Deployments.Delete("3d9c7ac3-7b87-4b4c-ba5d-750a91629d05")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Deleted:", out.Deleted)
}
