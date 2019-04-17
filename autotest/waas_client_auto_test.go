package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/waas"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createWaasClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := waas.NewWaasClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientAcceptRecommendations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "AcceptRecommendations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AcceptRecommendations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "AcceptRecommendations", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "AcceptRecommendations")
	assert.NoError(t, err)

	type AcceptRecommendationsRequestInfo struct {
		ContainerId string
		Request     waas.AcceptRecommendationsRequest
	}

	var requests []AcceptRecommendationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.AcceptRecommendations(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientCancelWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "CancelWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "CancelWorkRequest", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "CancelWorkRequest")
	assert.NoError(t, err)

	type CancelWorkRequestRequestInfo struct {
		ContainerId string
		Request     waas.CancelWorkRequestRequest
	}

	var requests []CancelWorkRequestRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CancelWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientCreateCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "CreateCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "CreateCertificate", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "CreateCertificate")
	assert.NoError(t, err)

	type CreateCertificateRequestInfo struct {
		ContainerId string
		Request     waas.CreateCertificateRequest
	}

	var requests []CreateCertificateRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientCreateWaasPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "CreateWaasPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateWaasPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "CreateWaasPolicy", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "CreateWaasPolicy")
	assert.NoError(t, err)

	type CreateWaasPolicyRequestInfo struct {
		ContainerId string
		Request     waas.CreateWaasPolicyRequest
	}

	var requests []CreateWaasPolicyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateWaasPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientDeleteCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "DeleteCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "DeleteCertificate", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "DeleteCertificate")
	assert.NoError(t, err)

	type DeleteCertificateRequestInfo struct {
		ContainerId string
		Request     waas.DeleteCertificateRequest
	}

	var requests []DeleteCertificateRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientDeleteWaasPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "DeleteWaasPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteWaasPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "DeleteWaasPolicy", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "DeleteWaasPolicy")
	assert.NoError(t, err)

	type DeleteWaasPolicyRequestInfo struct {
		ContainerId string
		Request     waas.DeleteWaasPolicyRequest
	}

	var requests []DeleteWaasPolicyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteWaasPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetCertificate", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetCertificate")
	assert.NoError(t, err)

	type GetCertificateRequestInfo struct {
		ContainerId string
		Request     waas.GetCertificateRequest
	}

	var requests []GetCertificateRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetDeviceFingerprintChallenge(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetDeviceFingerprintChallenge")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDeviceFingerprintChallenge is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetDeviceFingerprintChallenge", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetDeviceFingerprintChallenge")
	assert.NoError(t, err)

	type GetDeviceFingerprintChallengeRequestInfo struct {
		ContainerId string
		Request     waas.GetDeviceFingerprintChallengeRequest
	}

	var requests []GetDeviceFingerprintChallengeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDeviceFingerprintChallenge(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetHumanInteractionChallenge(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetHumanInteractionChallenge")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetHumanInteractionChallenge is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetHumanInteractionChallenge", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetHumanInteractionChallenge")
	assert.NoError(t, err)

	type GetHumanInteractionChallengeRequestInfo struct {
		ContainerId string
		Request     waas.GetHumanInteractionChallengeRequest
	}

	var requests []GetHumanInteractionChallengeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetHumanInteractionChallenge(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetJsChallenge(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetJsChallenge")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJsChallenge is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetJsChallenge", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetJsChallenge")
	assert.NoError(t, err)

	type GetJsChallengeRequestInfo struct {
		ContainerId string
		Request     waas.GetJsChallengeRequest
	}

	var requests []GetJsChallengeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetJsChallenge(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetPolicyConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetPolicyConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPolicyConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetPolicyConfig", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetPolicyConfig")
	assert.NoError(t, err)

	type GetPolicyConfigRequestInfo struct {
		ContainerId string
		Request     waas.GetPolicyConfigRequest
	}

	var requests []GetPolicyConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPolicyConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetProtectionRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetProtectionRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetProtectionRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetProtectionRule", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetProtectionRule")
	assert.NoError(t, err)

	type GetProtectionRuleRequestInfo struct {
		ContainerId string
		Request     waas.GetProtectionRuleRequest
	}

	var requests []GetProtectionRuleRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetProtectionRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetProtectionSettings(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetProtectionSettings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetProtectionSettings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetProtectionSettings", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetProtectionSettings")
	assert.NoError(t, err)

	type GetProtectionSettingsRequestInfo struct {
		ContainerId string
		Request     waas.GetProtectionSettingsRequest
	}

	var requests []GetProtectionSettingsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetProtectionSettings(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetWaasPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetWaasPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWaasPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetWaasPolicy", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetWaasPolicy")
	assert.NoError(t, err)

	type GetWaasPolicyRequestInfo struct {
		ContainerId string
		Request     waas.GetWaasPolicyRequest
	}

	var requests []GetWaasPolicyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetWaasPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetWafAddressRateLimiting(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetWafAddressRateLimiting")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWafAddressRateLimiting is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetWafAddressRateLimiting", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetWafAddressRateLimiting")
	assert.NoError(t, err)

	type GetWafAddressRateLimitingRequestInfo struct {
		ContainerId string
		Request     waas.GetWafAddressRateLimitingRequest
	}

	var requests []GetWafAddressRateLimitingRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetWafAddressRateLimiting(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetWafConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetWafConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWafConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetWafConfig", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetWafConfig")
	assert.NoError(t, err)

	type GetWafConfigRequestInfo struct {
		ContainerId string
		Request     waas.GetWafConfigRequest
	}

	var requests []GetWafConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetWafConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientGetWorkRequest(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "GetWorkRequest", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     waas.GetWorkRequestRequest
	}

	var requests []GetWorkRequestRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListAccessRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListAccessRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAccessRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListAccessRules", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListAccessRules")
	assert.NoError(t, err)

	type ListAccessRulesRequestInfo struct {
		ContainerId string
		Request     waas.ListAccessRulesRequest
	}

	var requests []ListAccessRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListAccessRulesRequest)
				return c.ListAccessRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListAccessRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListAccessRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListCaptchas(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListCaptchas")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCaptchas is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListCaptchas", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListCaptchas")
	assert.NoError(t, err)

	type ListCaptchasRequestInfo struct {
		ContainerId string
		Request     waas.ListCaptchasRequest
	}

	var requests []ListCaptchasRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListCaptchasRequest)
				return c.ListCaptchas(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListCaptchasResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListCaptchasResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListCertificates(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListCertificates")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCertificates is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListCertificates", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListCertificates")
	assert.NoError(t, err)

	type ListCertificatesRequestInfo struct {
		ContainerId string
		Request     waas.ListCertificatesRequest
	}

	var requests []ListCertificatesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListCertificatesRequest)
				return c.ListCertificates(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListCertificatesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListCertificatesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListEdgeSubnets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListEdgeSubnets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListEdgeSubnets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListEdgeSubnets", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListEdgeSubnets")
	assert.NoError(t, err)

	type ListEdgeSubnetsRequestInfo struct {
		ContainerId string
		Request     waas.ListEdgeSubnetsRequest
	}

	var requests []ListEdgeSubnetsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListEdgeSubnetsRequest)
				return c.ListEdgeSubnets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListEdgeSubnetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListEdgeSubnetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListGoodBots(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListGoodBots")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListGoodBots is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListGoodBots", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListGoodBots")
	assert.NoError(t, err)

	type ListGoodBotsRequestInfo struct {
		ContainerId string
		Request     waas.ListGoodBotsRequest
	}

	var requests []ListGoodBotsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListGoodBotsRequest)
				return c.ListGoodBots(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListGoodBotsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListGoodBotsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListProtectionRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListProtectionRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProtectionRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListProtectionRules", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListProtectionRules")
	assert.NoError(t, err)

	type ListProtectionRulesRequestInfo struct {
		ContainerId string
		Request     waas.ListProtectionRulesRequest
	}

	var requests []ListProtectionRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListProtectionRulesRequest)
				return c.ListProtectionRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListProtectionRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListProtectionRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListRecommendations(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListRecommendations")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRecommendations is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListRecommendations", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListRecommendations")
	assert.NoError(t, err)

	type ListRecommendationsRequestInfo struct {
		ContainerId string
		Request     waas.ListRecommendationsRequest
	}

	var requests []ListRecommendationsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListRecommendationsRequest)
				return c.ListRecommendations(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListRecommendationsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListRecommendationsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListThreatFeeds(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListThreatFeeds")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListThreatFeeds is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListThreatFeeds", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListThreatFeeds")
	assert.NoError(t, err)

	type ListThreatFeedsRequestInfo struct {
		ContainerId string
		Request     waas.ListThreatFeedsRequest
	}

	var requests []ListThreatFeedsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListThreatFeedsRequest)
				return c.ListThreatFeeds(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListThreatFeedsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListThreatFeedsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListWaasPolicies(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListWaasPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWaasPolicies is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListWaasPolicies", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListWaasPolicies")
	assert.NoError(t, err)

	type ListWaasPoliciesRequestInfo struct {
		ContainerId string
		Request     waas.ListWaasPoliciesRequest
	}

	var requests []ListWaasPoliciesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListWaasPoliciesRequest)
				return c.ListWaasPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListWaasPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListWaasPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListWafBlockedRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListWafBlockedRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWafBlockedRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListWafBlockedRequests", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListWafBlockedRequests")
	assert.NoError(t, err)

	type ListWafBlockedRequestsRequestInfo struct {
		ContainerId string
		Request     waas.ListWafBlockedRequestsRequest
	}

	var requests []ListWafBlockedRequestsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListWafBlockedRequestsRequest)
				return c.ListWafBlockedRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListWafBlockedRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListWafBlockedRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListWafLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListWafLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWafLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListWafLogs", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListWafLogs")
	assert.NoError(t, err)

	type ListWafLogsRequestInfo struct {
		ContainerId string
		Request     waas.ListWafLogsRequest
	}

	var requests []ListWafLogsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListWafLogsRequest)
				return c.ListWafLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListWafLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListWafLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListWafRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListWafRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWafRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListWafRequests", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListWafRequests")
	assert.NoError(t, err)

	type ListWafRequestsRequestInfo struct {
		ContainerId string
		Request     waas.ListWafRequestsRequest
	}

	var requests []ListWafRequestsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListWafRequestsRequest)
				return c.ListWafRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListWafRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListWafRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListWafTraffic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListWafTraffic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWafTraffic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListWafTraffic", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListWafTraffic")
	assert.NoError(t, err)

	type ListWafTrafficRequestInfo struct {
		ContainerId string
		Request     waas.ListWafTrafficRequest
	}

	var requests []ListWafTrafficRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListWafTrafficRequest)
				return c.ListWafTraffic(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListWafTrafficResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListWafTrafficResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListWhitelists(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListWhitelists")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWhitelists is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListWhitelists", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListWhitelists")
	assert.NoError(t, err)

	type ListWhitelistsRequestInfo struct {
		ContainerId string
		Request     waas.ListWhitelistsRequest
	}

	var requests []ListWhitelistsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListWhitelistsRequest)
				return c.ListWhitelists(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListWhitelistsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListWhitelistsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientListWorkRequests(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "ListWorkRequests", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     waas.ListWorkRequestsRequest
	}

	var requests []ListWorkRequestsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*waas.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]waas.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(waas.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateAccessRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateAccessRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAccessRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateAccessRules", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateAccessRules")
	assert.NoError(t, err)

	type UpdateAccessRulesRequestInfo struct {
		ContainerId string
		Request     waas.UpdateAccessRulesRequest
	}

	var requests []UpdateAccessRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAccessRules(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateCaptchas(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateCaptchas")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCaptchas is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateCaptchas", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateCaptchas")
	assert.NoError(t, err)

	type UpdateCaptchasRequestInfo struct {
		ContainerId string
		Request     waas.UpdateCaptchasRequest
	}

	var requests []UpdateCaptchasRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateCaptchas(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateCertificate(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCertificate is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateCertificate", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateCertificate")
	assert.NoError(t, err)

	type UpdateCertificateRequestInfo struct {
		ContainerId string
		Request     waas.UpdateCertificateRequest
	}

	var requests []UpdateCertificateRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateDeviceFingerprintChallenge(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateDeviceFingerprintChallenge")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDeviceFingerprintChallenge is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateDeviceFingerprintChallenge", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateDeviceFingerprintChallenge")
	assert.NoError(t, err)

	type UpdateDeviceFingerprintChallengeRequestInfo struct {
		ContainerId string
		Request     waas.UpdateDeviceFingerprintChallengeRequest
	}

	var requests []UpdateDeviceFingerprintChallengeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDeviceFingerprintChallenge(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateGoodBots(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateGoodBots")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateGoodBots is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateGoodBots", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateGoodBots")
	assert.NoError(t, err)

	type UpdateGoodBotsRequestInfo struct {
		ContainerId string
		Request     waas.UpdateGoodBotsRequest
	}

	var requests []UpdateGoodBotsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateGoodBots(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateHumanInteractionChallenge(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateHumanInteractionChallenge")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateHumanInteractionChallenge is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateHumanInteractionChallenge", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateHumanInteractionChallenge")
	assert.NoError(t, err)

	type UpdateHumanInteractionChallengeRequestInfo struct {
		ContainerId string
		Request     waas.UpdateHumanInteractionChallengeRequest
	}

	var requests []UpdateHumanInteractionChallengeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateHumanInteractionChallenge(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateJsChallenge(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateJsChallenge")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateJsChallenge is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateJsChallenge", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateJsChallenge")
	assert.NoError(t, err)

	type UpdateJsChallengeRequestInfo struct {
		ContainerId string
		Request     waas.UpdateJsChallengeRequest
	}

	var requests []UpdateJsChallengeRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateJsChallenge(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdatePolicyConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdatePolicyConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePolicyConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdatePolicyConfig", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdatePolicyConfig")
	assert.NoError(t, err)

	type UpdatePolicyConfigRequestInfo struct {
		ContainerId string
		Request     waas.UpdatePolicyConfigRequest
	}

	var requests []UpdatePolicyConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdatePolicyConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateProtectionRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateProtectionRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateProtectionRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateProtectionRules", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateProtectionRules")
	assert.NoError(t, err)

	type UpdateProtectionRulesRequestInfo struct {
		ContainerId string
		Request     waas.UpdateProtectionRulesRequest
	}

	var requests []UpdateProtectionRulesRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateProtectionRules(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateProtectionSettings(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateProtectionSettings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateProtectionSettings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateProtectionSettings", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateProtectionSettings")
	assert.NoError(t, err)

	type UpdateProtectionSettingsRequestInfo struct {
		ContainerId string
		Request     waas.UpdateProtectionSettingsRequest
	}

	var requests []UpdateProtectionSettingsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateProtectionSettings(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateThreatFeeds(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateThreatFeeds")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateThreatFeeds is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateThreatFeeds", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateThreatFeeds")
	assert.NoError(t, err)

	type UpdateThreatFeedsRequestInfo struct {
		ContainerId string
		Request     waas.UpdateThreatFeedsRequest
	}

	var requests []UpdateThreatFeedsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateThreatFeeds(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateWaasPolicy(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateWaasPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateWaasPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateWaasPolicy", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateWaasPolicy")
	assert.NoError(t, err)

	type UpdateWaasPolicyRequestInfo struct {
		ContainerId string
		Request     waas.UpdateWaasPolicyRequest
	}

	var requests []UpdateWaasPolicyRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateWaasPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateWafAddressRateLimiting(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateWafAddressRateLimiting")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateWafAddressRateLimiting is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateWafAddressRateLimiting", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateWafAddressRateLimiting")
	assert.NoError(t, err)

	type UpdateWafAddressRateLimitingRequestInfo struct {
		ContainerId string
		Request     waas.UpdateWafAddressRateLimitingRequest
	}

	var requests []UpdateWafAddressRateLimitingRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateWafAddressRateLimiting(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateWafConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateWafConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateWafConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateWafConfig", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateWafConfig")
	assert.NoError(t, err)

	type UpdateWafConfigRequestInfo struct {
		ContainerId string
		Request     waas.UpdateWafConfigRequest
	}

	var requests []UpdateWafConfigRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateWafConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="oci_waas_dev_us_grp@oracle.com" jiraProject="WAAS" opsJiraProject="WAF"
func TestWaasClientUpdateWhitelists(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("waas", "UpdateWhitelists")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateWhitelists is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("waas", "Waas", "UpdateWhitelists", createWaasClientWithProvider)
	assert.NoError(t, err)
	c := cc.(waas.WaasClient)

	body, err := testClient.getRequests("waas", "UpdateWhitelists")
	assert.NoError(t, err)

	type UpdateWhitelistsRequestInfo struct {
		ContainerId string
		Request     waas.UpdateWhitelistsRequest
	}

	var requests []UpdateWhitelistsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateWhitelists(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
