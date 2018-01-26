package objectstorage

import (
	"fmt"
	"strings"
	"testing"

	"github.com/oracle/oci-go-sdk/common"
	identitySample "github.com/oracle/oci-go-sdk/example/identity"
)

func TestMain(t *testing.T) {
	// The OCID of the tenancy containing the compartment.
	tenancyID, _ := common.DefaultConfigProvider().TenancyOCID()
	rootCompartment, _ := identitySample.GetCompartment(tenancyID)
	fmt.Printf("find compartment: %s \n", *rootCompartment.Name)

	// TODO: why has to be lower case
	namespaceName := strings.ToLower(*rootCompartment.Name)

	// create a compartment
	compartmentId, err := identitySample.CreateCompartment()

	// create a bucket
	err = CreateBucket(namespaceName, *compartmentId, "testBucket")

	if err != nil {
		fmt.Printf("Error creating bucket: %s", err)
		t.Error(err)
		t.Fail()
	}

	UploadFileToBucket(namespaceName, "testBucket", "testObject", "./objectstorage_test.go")
}
