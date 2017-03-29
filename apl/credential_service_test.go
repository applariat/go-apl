package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

var (
	testCredentialId string
	testCredentialFilter string
)

func TestCredentialService_Create(t *testing.T) {

	testCredentialId = "credential-test-id"
	testCredentialFilter = "docker"

	aplSvs := apl.NewClient()

	in := &apl.CredentialCreateInput{
		ID:             testCredentialId,
		Name:           "creds for credential test",
		CredentialType: testCredentialFilter,
		Credentials: map[string]string{
			"registry_email":    "floopy@glorp.com",
			"registry_password": "urpasswd",
			"registry_user":     "apl_registry",
		},
	}

	out, _, err := aplSvs.Credentials.Create(in)

	if err != nil {
		t.Fatal(err)
	}


	fmt.Println("New Credential ID:", out.Data)

}

func TestCredentialService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Credentials.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("Credential found %d rows\n", rowCount)
	} else {
		t.Fatal("No Credential rows found")
	}

}

func TestCredentialService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.CredentialParams{
		CredentialType: testCredentialFilter,
	}
	out, _, err := aplSvc.Credentials.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Credential rows found for filter", testCredentialFilter)
	}

	fmt.Printf("Credential filtered found %d rows for filter \"%s\"\n", rowCount, testCredentialFilter)

}

func TestCredentialService_Update(t *testing.T) {
	aplSvc := apl.NewClient()


	in := &apl.CredentialUpdateInput{Name: "UPDATED!"}
	out, _, err := aplSvc.Credentials.Update(testCredentialId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}

func TestCredentialService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Credentials.Get(testCredentialId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Credential{}) {
		t.Fatal("No Credential found for ID", testCredentialId)
	}

	fmt.Println("Credential found for ID", testCredentialId)

}

func TestCredentialService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Credentials.Delete(testCredentialId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
