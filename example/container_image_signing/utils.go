// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package example

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oracle/oci-go-sdk/v44/artifacts"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/keymanagement"
	"regexp"
	"strings"
)

/*
SignAndUploadContainerImageSignatureMetadata calls KMS to sign the message then calls OCIR to upload the returned signature
  Description: Sign a container image and upload the signature to the image
  Response: The signed container image signature metadata.
  Parameters:
   - kmsKeyId:
  	  description: The OCID of the kmsKeyId used to sign the container image. eg) ocid1.key.oc1..exampleuniqueID
  	  maxLength: 255
  	  minLength: 1
   - kmsKeyVersionId:
  	  description: The OCID of the kmsKeyVersionId used to sign the container image. eg) ocid1.keyversion.oc1..exampleuniqueID
  	  maxLength: 255
  	  minLength: 1
   - signingAlgorithm:
  	  - description: The algorithm to be used for signing. These are the only supported signing algorithms for container images.
  	  	 - SHA_224_RSA_PKCS_PSS
  	  	 - SHA_256_RSA_PKCS_PSS
  	  	 - SHA_384_RSA_PKCS_PSS
  	  	 - SHA_512_RSA_PKCS_PSS
   - compartmentId:
  	  description: The OCID of the compartment in which the container repository exists. eg) ocid1.compartment.oc1..exampleuniqueID
  	  maxLength: 100
  	  minLength: 1
   - imageId:
  	  description: The OCID of the container image. eg) ocid1.containerimage.oc1..exampleuniqueID
  	  maxLength: 255
  	  minLength: 1
   - repoPath:
  	  description The docker repository path. eg) odx-registry/busybox
   - imageDigest:
  	  description: The sha256 digest of the docker image. eg) sha256:12345
   - description:
  	  description: An user inputted message.
   - metadata:
  	  description: An user defined information about the container image in JSON format eg) {"buildNumber":"123"}
  	  restriction:
  	   - should only contains alphanumeric key strings.
  	   - should be alphabetically sorted.
  	   - should not have whitespaces or escape characters.
*/
func SignAndUploadContainerImageSignatureMetadata(ctx context.Context, artifactClient artifacts.ArtifactsClient, configProvider common.ConfigurationProvider, kmsKeyId string, kmsKeyVersionId string, signingAlgorithm string, compartmentId string, imageId string, description string, metadata string) (*artifacts.ContainerImageSignature, error) {
	signingAlgoKms := mappingSignDataDetailsSigningAlgorithm[signingAlgorithm]
	signingAlgoOcir := mappingCreateContainerImageSignatureDetailsSigningAlgorithm[signingAlgorithm]

	region, err := configProvider.Region()
	if err != nil {
		return nil, err
	}

	// Create KMS client
	vaultCryptoClient, cryptoClientError := buildVaultCryptoClient(configProvider, kmsKeyId, region)
	if cryptoClientError != nil {
		return nil, cryptoClientError
	}

	// Get container image metadata
	common.Debugf("Obtaining container image metadata by the image ID")
	containerImageMetadata, err := getContainerImageMetadata(ctx, artifactClient, imageId)
	if err != nil {
		return nil, err
	}
	common.Debugf(fmt.Sprintf("Container image metadata: %s", *containerImageMetadata))

	// Generate message
	message := Message{
		Description:      description,
		ImageDigest:      *containerImageMetadata.Digest,
		KmsKeyId:         kmsKeyId,
		KmsKeyVersionId:  kmsKeyVersionId,
		Metadata:         metadata,
		Region:           region,
		RepositoryName:   *containerImageMetadata.RepositoryName,
		SigningAlgorithm: signingAlgorithm,
	}

	jsonByte, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	encodedJson := base64.StdEncoding.EncodeToString(jsonByte)

	// Sign image digest
	common.Debugf("Generating signature")
	signedData, err := signContainerImage(ctx, vaultCryptoClient, encodedJson, kmsKeyId, kmsKeyVersionId, signingAlgoKms)
	if err != nil {
		return nil, err
	}
	common.Debugf(fmt.Sprintf("Signature: %s", *signedData.Signature))

	// Upload signature metadata
	common.Debugf("Uploading signature")
	containerImageSignatureUploaded, err := uploadSignatureMetadata(ctx, artifactClient, compartmentId, imageId, kmsKeyId, kmsKeyVersionId, signingAlgoOcir, encodedJson, *signedData.Signature)
	if err != nil {
		return nil, err
	}
	common.Debugf(fmt.Sprintf("Uploaded signature: %s\nMessage: %s\nID: %s\n", *containerImageSignatureUploaded.Signature, *containerImageSignatureUploaded.Message, *containerImageSignatureUploaded.Id))

	return containerImageSignatureUploaded, err
}

/*
GetAndVerifyImageSignatureMetadata calls OCIR to list all the signatures satisfying the user provided criterion then calls KMS to verify the returned signatures
  Description: Fetch a container image signature metadata and verity the signature
  Response: Boolean to indicate if any of the signatures of the container image is verified
  Parameters:
   - compartmentId:
	  description: The OCID of the compartment in which the container repository exists. eg) ocid1.compartment.oc1..exampleuniqueID
	  maxLength: 100
	  minLength: 1
   - compartmentIdInSubtree:
	  description: When set to true, the hierarchy of compartments is traversed
   - repositoryName:
	  description: The repository name in which the container image exists eg) busybox
   - imageDigest:
	  description: The sha256 digest of the docker image. eg) sha256:12345
   - trustedKeys:
	  description: List of OCIDs of the kmsKeyId used to sign the container image.
*/
func GetAndVerifyImageSignatureMetadata(ctx context.Context, artifactClient artifacts.ArtifactsClient, configProvider common.ConfigurationProvider, compartmentId string, compartmentIdInSubtree bool, repositoryName string, imageDigest string, trustedKeys []string) (bool, error) {
	return getAndVerifyImageSignatureMetadataHelper(ctx, artifactClient, configProvider, compartmentId, compartmentIdInSubtree, repositoryName, imageDigest, trustedKeys, "")
}

// getAndVerifyImageSignatureMetadataHelper is a helper function of GetAndVerifyImageSignatureMetadata
func getAndVerifyImageSignatureMetadataHelper(ctx context.Context, artifactClient artifacts.ArtifactsClient, configProvider common.ConfigurationProvider, compartmentId string, compartmentIdInSubtree bool, repositoryName string, imageDigest string, trustedKeys []string, page string) (bool, error) {
	// Fetch signature metadata
	common.Debugf("Fetching signatures")
	signatureCollection, nextPage, listErr := listContainerImageSignaturesWithRepoPath(ctx, artifactClient, compartmentId, compartmentIdInSubtree, repositoryName, imageDigest, page)
	if listErr != nil {
		return false, listErr
	}
	common.Debugf(fmt.Sprintf("Fetched signature: %d signatures in compartment %s with image URL: %s:%s. Remaining count %d", len(signatureCollection.Items), compartmentId, repositoryName, imageDigest, *signatureCollection.RemainingItemsCount))

	// filter out the keys
	ContainerImageSignatureSummaries := filterItemByTrustedKeys(signatureCollection.Items, trustedKeys)
	common.Debugf(fmt.Sprintf("Filtered out %d signatures by the trusted keys", len(signatureCollection.Items)-len(ContainerImageSignatureSummaries)))

	// Verify signature
	common.Debugf("Verifying signature")
	verified, verifyErr := verifySignatures(ctx, configProvider, ContainerImageSignatureSummaries)

	if verifyErr != nil {
		return false, verifyErr
	} else if !verified && nextPage != "" {
		return getAndVerifyImageSignatureMetadataHelper(ctx, artifactClient, configProvider, compartmentId, compartmentIdInSubtree, repositoryName, imageDigest, trustedKeys, nextPage)
	} else {
		return verified, nil
	}
}

// signContainerImage returns the signature of the provided message signed with the provided keyId, keyVersionId, and signingAlgorithm
func signContainerImage(ctx context.Context, client keymanagement.KmsCryptoClient, message string, keyId string, keyVersionId string, signingAlgorithm keymanagement.SignDataDetailsSigningAlgorithmEnum) (*keymanagement.SignedData, error) {
	signDataDetails := keymanagement.SignDataDetails{
		Message:          &message,
		KeyId:            &keyId,
		KeyVersionId:     &keyVersionId,
		SigningAlgorithm: signingAlgorithm,
		MessageType:      keymanagement.SignDataDetailsMessageTypeRaw,
	}

	request := keymanagement.SignRequest{
		SignDataDetails: signDataDetails,
	}

	response, err := client.Sign(ctx, request)
	if err != nil {
		return nil, err
	}

	return &response.SignedData, err
}

// getContainerImageMetadata returns ...
func getContainerImageMetadata(ctx context.Context, artifactClient artifacts.ArtifactsClient, imageId string) (*artifacts.ContainerImage, error) {
	request := artifacts.GetContainerImageRequest{
		ImageId: &imageId,
	}

	response, err := artifactClient.GetContainerImage(ctx, request)
	if err != nil {
		return nil, err
	}

	return &response.ContainerImage, err
}

// uploadSignatureMetadata uploads the signature and signed message to the container image with the imageId
func uploadSignatureMetadata(ctx context.Context, artifactClient artifacts.ArtifactsClient, compartmentId string, imageId string, keyId string, keyVersionId string, signingAlgorithm artifacts.CreateContainerImageSignatureDetailsSigningAlgorithmEnum, message string, signature string) (*artifacts.ContainerImageSignature, error) {
	signatureDetails := artifacts.CreateContainerImageSignatureDetails{
		CompartmentId:    &compartmentId,
		ImageId:          &imageId,
		KmsKeyId:         &keyId,
		KmsKeyVersionId:  &keyVersionId,
		SigningAlgorithm: signingAlgorithm,
		Message:          &message,
		Signature:        &signature,
	}

	request := artifacts.CreateContainerImageSignatureRequest{
		CreateContainerImageSignatureDetails: signatureDetails,
	}

	response, err := artifactClient.CreateContainerImageSignature(ctx, request)
	if err != nil {
		return nil, err
	}

	if response.RawResponse.StatusCode != 200 {
		err := fmt.Sprintf("failed to upload the signature to OCI Registry. Status Code: %d", response.RawResponse.StatusCode)
		return nil, errors.New(err)
	}

	return &response.ContainerImageSignature, err
}

// listContainerImageSignaturesWithRepoPath lists all the container image signatures uploaded to the image referenced with the repositoryName and imageDigest
func listContainerImageSignaturesWithRepoPath(ctx context.Context, artifactClient artifacts.ArtifactsClient, compartmentId string, compartmentIdInSubtree bool, repositoryName string, imageDigest string, page string) (*artifacts.ContainerImageSignatureCollection, string, error) {
	var request artifacts.ListContainerImageSignaturesRequest

	if page != "" {
		request = artifacts.ListContainerImageSignaturesRequest{
			CompartmentId:          &compartmentId,
			CompartmentIdInSubtree: &compartmentIdInSubtree,
			RepositoryName:         &repositoryName,
			ImageDigest:            &imageDigest,
			Page:                   &page,
		}
	} else {
		request = artifacts.ListContainerImageSignaturesRequest{
			CompartmentId:          &compartmentId,
			CompartmentIdInSubtree: &compartmentIdInSubtree,
			RepositoryName:         &repositoryName,
			ImageDigest:            &imageDigest,
		}
	}

	response, err := artifactClient.ListContainerImageSignatures(ctx, request)
	if err != nil {
		return nil, "", err
	}

	if response.RawResponse.StatusCode != 200 {
		err := fmt.Sprintf("failed to list the signatures of repositoryName %s, imageDigest %s. Status Code: %d", repositoryName, imageDigest, response.RawResponse.StatusCode)
		return nil, "", errors.New(err)
	}

	var nextPage string
	if response.OpcNextPage == nil {
		nextPage = ""
	} else {
		nextPage = *response.OpcNextPage
	}

	return &response.ContainerImageSignatureCollection, nextPage, err
}

// verifySignatures returns true if any of the provided container image signatures in the ContainerImageSignatureSummary items is valid
func verifySignatures(ctx context.Context, configProvider common.ConfigurationProvider, containerImageSignatureSummary []artifacts.ContainerImageSignatureSummary) (bool, error) {
	region, err := configProvider.Region()
	if err != nil {
		return false, err
	}

	for _, signatureSummary := range containerImageSignatureSummary {
		vaultCryptoClient, cryptoClientError := buildVaultCryptoClient(configProvider, *signatureSummary.KmsKeyId, region)
		if cryptoClientError != nil {
			return false, cryptoClientError
		}

		algo := mappingVerifyDataDetailsSigningAlgorithm[string(signatureSummary.SigningAlgorithm)]
		verifiedData, verifyErr := verifySignature(ctx, vaultCryptoClient, *signatureSummary.Message, *signatureSummary.Signature, *signatureSummary.KmsKeyId, *signatureSummary.KmsKeyVersionId, algo)
		if verifyErr != nil {
			return false, verifyErr
		} else if *verifiedData.IsSignatureValid {
			return true, nil
		}
	}
	return false, nil
}

// verifySignature verifies if the provided signature is valid given signed message, keyId, keyVersionId, and signingAlgorithm
func verifySignature(ctx context.Context, client keymanagement.KmsCryptoClient, message string, signature string, keyId string, keyVersionId string, signingAlgorithm keymanagement.VerifyDataDetailsSigningAlgorithmEnum) (*keymanagement.VerifiedData, error) {
	verifyDataDetails := keymanagement.VerifyDataDetails{
		KeyId:            &keyId,
		KeyVersionId:     &keyVersionId,
		SigningAlgorithm: signingAlgorithm,
		Message:          &message,
		Signature:        &signature,
	}

	request := keymanagement.VerifyRequest{
		VerifyDataDetails: verifyDataDetails,
	}

	response, err := client.Verify(ctx, request)
	if err != nil {
		return nil, err
	}

	return &response.VerifiedData, err
}

// buildVaultCryptoClient builds the KmsCryptoClient based on the vault extension OCID in the keyId
func buildVaultCryptoClient(configProvider common.ConfigurationProvider, keyId string, region string) (client keymanagement.KmsCryptoClient, err error) {
	re := regexp.MustCompile("ocid1\\.key\\.([\\w-]+)\\.([\\w-]+)\\.([\\w-]+)\\.([\\w]{60})")
	vaultExt := re.FindStringSubmatch(keyId)[3]
	cryptoEndpointTemplate := common.StringToRegion(region).EndpointForTemplate("kms", "https://{vaultExt}-crypto.kms.{region}.{secondLevelDomain}")
	cryptoEndpoint := strings.Replace(cryptoEndpointTemplate, "{vaultExt}", vaultExt, 1)
	common.Debugf(fmt.Sprintf("Built vault crypto client with endpoint: %s", cryptoEndpoint))

	return keymanagement.NewKmsCryptoClientWithConfigurationProvider(configProvider, cryptoEndpoint)
}

// filterItemByTrustedKeys filters out ContainerImageSignatureSummary items if its key is not in the trustedKeys list
func filterItemByTrustedKeys(items []artifacts.ContainerImageSignatureSummary, trustedKeys []string) (ret []artifacts.ContainerImageSignatureSummary) {
	for _, item := range items {
		if isTrustedKey(*item.KmsKeyId, trustedKeys) {
			ret = append(ret, item)
		}
	}
	return
}

// isTrustedKey returns true if a key is in the trustedKeys key list
func isTrustedKey(key string, trustedKeys []string) bool {
	for _, k := range trustedKeys {
		if k == key {
			return true
		}
	}
	return false
}

// Message defines the struct of container image signature payload
type Message struct {
	Description      string `mandatory:"true" json:"description"`
	ImageDigest      string `mandatory:"true" json:"imageDigest"`
	KmsKeyId         string `mandatory:"true" json:"kmsKeyId"`
	KmsKeyVersionId  string `mandatory:"true" json:"kmsKeyVersionId"`
	Metadata         string `mandatory:"true" json:"metadata"`
	Region           string `mandatory:"true" json:"region"`
	RepositoryName   string `mandatory:"true" json:"repositoryName"`
	SigningAlgorithm string `mandatory:"true" json:"signingAlgorithm"`
}

// mappingCreateContainerImageSignatureDetailsSigningAlgorithm duplicated from the artifacts package since they're not exported
var mappingCreateContainerImageSignatureDetailsSigningAlgorithm = map[string]artifacts.CreateContainerImageSignatureDetailsSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": artifacts.CreateContainerImageSignatureDetailsSigningAlgorithm224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": artifacts.CreateContainerImageSignatureDetailsSigningAlgorithm256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": artifacts.CreateContainerImageSignatureDetailsSigningAlgorithm384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": artifacts.CreateContainerImageSignatureDetailsSigningAlgorithm512RsaPkcsPss,
}

// mappingSignDataDetailsSigningAlgorithm duplicated from the keymanagement package since they're not exported
var mappingSignDataDetailsSigningAlgorithm = map[string]keymanagement.SignDataDetailsSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": keymanagement.SignDataDetailsSigningAlgorithmSha224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": keymanagement.SignDataDetailsSigningAlgorithmSha256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": keymanagement.SignDataDetailsSigningAlgorithmSha384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": keymanagement.SignDataDetailsSigningAlgorithmSha512RsaPkcsPss,
}

// mappingVerifyDataDetailsSigningAlgorithm duplicated from the keymanagement package since they're not exported
var mappingVerifyDataDetailsSigningAlgorithm = map[string]keymanagement.VerifyDataDetailsSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS": keymanagement.VerifyDataDetailsSigningAlgorithmSha224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS": keymanagement.VerifyDataDetailsSigningAlgorithmSha256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS": keymanagement.VerifyDataDetailsSigningAlgorithmSha384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS": keymanagement.VerifyDataDetailsSigningAlgorithmSha512RsaPkcsPss,
}
