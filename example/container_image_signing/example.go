// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package example

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v55/artifacts"
	"github.com/oracle/oci-go-sdk/v55/common"
	"github.com/oracle/oci-go-sdk/v55/example/helpers"
)

func main() {
	configProvider := common.DefaultConfigProvider()
	ctx := context.Background()

	// Create artifact client
	artifactClient, artifactError := artifacts.NewArtifactsClientWithConfigurationProvider(configProvider)
	helpers.FatalIfError(artifactError)

	// Upload Image and Signature Flow
	kmsKeyId := "ocid1.key.oc1..exampleuniqueID"
	kmsKeyVersionId := "ocid1.keyversion.oc1..exampleuniqueID"
	signingAlgo := "SHA_512_RSA_PKCS_PSS"
	compartmentId := "ocid1.compartment.oc1..exampleuniqueID"
	imageId := "ocid1.containerimage.oc1..exampleuniqueID"
	description := "Image built by TC"
	metadata := "{\"buildNumber\":\"123\"}"

	signature, err := SignAndUploadContainerImageSignatureMetadata(ctx, artifactClient, configProvider, kmsKeyId, kmsKeyVersionId, signingAlgo, compartmentId, imageId, description, metadata)
	helpers.FatalIfError(err)
	common.Logf(fmt.Sprintf("A signature has been successfully uploaded: %s", *signature))

	// Pull Image and Verify Signature Flow
	repoName := "repo-name"
	imageDigest := "sha256:12345"
	trustedKeys := []string{"ocid1.key.oc1..keyId1", "ocid1.key.oc1..keyId2"}

	verified, err := GetAndVerifyImageSignatureMetadata(ctx, artifactClient, configProvider, compartmentId, false, repoName, imageDigest, trustedKeys)
	helpers.FatalIfError(err)
	if verified {
		common.Logf("At least one of the signatures is verified")
	} else {
		common.Logf("None of the signatures is verified")
	}
}
