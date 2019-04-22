package example

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/resourcemanager"
)

// ExampleResourceManager for how to do CRUD for Resource Manager Stack
// The comparement id is read from the environment variable OCI_COMPARTMENT_ID
func ExampleResourceManager() {
	provider := common.DefaultConfigProvider()
	client, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(provider)
	helpers.FatalIfError(err)

	ctx := context.Background()

	stackID := createStack(ctx, provider, client)
	defer deleteStack(ctx, stackID, client)
	listStacks(ctx, client)
	updateStack(ctx, stackID, client)
	getStack(ctx, stackID, client)

	// Output:
	// create stack completed
	// list stacks completed
	// update stack completed
	// get stack completed
	// delete stack completed
}

var zipFile string = "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"

func createStack(ctx context.Context, provider common.ConfigurationProvider, client resourcemanager.ResourceManagerClient) string {
	stackName := fmt.Sprintf("test-%s", helpers.GetRandomString(8))
	region, _ := provider.Region()
	tenancy_ocid, _ := provider.TenancyOCID()

	// create resource manager stack with type ZIP_UPLOAD by passing a base64 encoded Terraform zip string
	// uer has multiple ways to create stack, details check https://docs.cloud.oracle.com/iaas/api/#/en/resourcemanager/20180917/datatypes/CreateConfigSourceDetails
	req := resourcemanager.CreateStackRequest{
		CreateStackDetails: resourcemanager.CreateStackDetails{
			CompartmentId: helpers.CompartmentID(),
			ConfigSource: resourcemanager.CreateZipUploadConfigSourceDetails{
				WorkingDirectory:     common.String("vcn"),
				ZipFileBase64Encoded: common.String(zipFile),
			},
			DisplayName: common.String(stackName),
			Description: common.String(fmt.Sprintf("%s-description", stackName)),
			Variables: map[string]string{
				"compartment_ocid": *helpers.CompartmentID(),
				"region":           region,
				"tenancy_ocid":     tenancy_ocid,
			},
		},
	}

	stackResp, err := client.CreateStack(ctx, req)
	helpers.FatalIfError(err)

	fmt.Println("create stack completed")
	return *stackResp.Stack.Id
}

func updateStack(ctx context.Context, StackId string, client resourcemanager.ResourceManagerClient) {
	stackName := fmt.Sprintf("test-v1-%s", helpers.GetRandomString(8))

	// update displayName and description of resource manager stack
	req := resourcemanager.UpdateStackRequest{
		StackId: common.String(StackId),
		UpdateStackDetails: resourcemanager.UpdateStackDetails{
			DisplayName: common.String(stackName),
			Description: common.String(fmt.Sprintf("%s-description", stackName)),
		},
	}

	_, err := client.UpdateStack(ctx, req)
	helpers.FatalIfError(err)

	fmt.Println("update stack completed")
}

func listStacks(ctx context.Context, client resourcemanager.ResourceManagerClient) {
	req := resourcemanager.ListStacksRequest{
		CompartmentId: helpers.CompartmentID(),
	}

	// list resource manager stack
	_, err := client.ListStacks(ctx, req)
	helpers.FatalIfError(err)

	fmt.Println("list stacks completed")
}

func getStack(ctx context.Context, StackId string, client resourcemanager.ResourceManagerClient) {
	req := resourcemanager.GetStackRequest{
		StackId: common.String(StackId),
	}

	// get details a particular resource manager stack
	_, err := client.GetStack(ctx, req)
	helpers.FatalIfError(err)

	fmt.Println("get stack completed")
}

func deleteStack(ctx context.Context, StackId string, client resourcemanager.ResourceManagerClient) {
	req := resourcemanager.DeleteStackRequest{
		StackId: common.String(StackId),
	}

	// delete a resource manager stack
	_, err := client.DeleteStack(ctx, req)
	helpers.FatalIfError(err)

	fmt.Println("delete stack completed")
}
