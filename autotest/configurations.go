package autotest

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"math"
	"os"
	"time"
)

type TestConfiguration struct {
	ConfigurationProvider common.ConfigurationProvider
	CompartmentID         string
}

func configurationProvider() common.ConfigurationProvider {
	fileConfig := os.Getenv("SDK_FILE_CONFIG")
	if fileConfig == "" {
		return common.ConfigurationProviderEnvironmentVariables("TF_VAR", "")
	}

	fileConfigProvider, _ := common.ConfigurationProviderFromFile(fileConfig, "")
	return fileConfigProvider
}

func newTestConfiguration() (*TestConfiguration, error) {

	confProvider := configurationProvider()

	compartmentID, err := getEnvOrError("INTEGRATION_TEST_COMPARTMENT")
	if err == nil {
		fmt.Println("INTEGRATION_TEST_COMPARTMENT variable found, will be overriden by service side configurations")
	}

	return &TestConfiguration{
		ConfigurationProvider: confProvider,
		CompartmentID:         compartmentID,
	}, nil
}

func getEnvOrError(name string) (string, error) {
	env := os.Getenv(name)
	if env == "" {
		return "", fmt.Errorf("%s not present", name)
	}
	return env, nil
}

func retryPolicyForTests() *common.RetryPolicy {
	shouldRetry := func(res common.OCIOperationResponse) bool {
		//Stop retrying on no errors and  non-service errors
		failure, ok := common.IsServiceError(res.Error)
		if !ok {
			return false
		}

		//Stop retrying on service errors other than rate limiting
		if failure.GetHTTPStatusCode() != 429 {
			return false
		}

		//Otherwise keep retrying
		return true
	}

	exponentialBackoff := func(res common.OCIOperationResponse) time.Duration {
		backoff := math.Pow(2, float64(res.AttemptNumber))
		duration := time.Duration(backoff) * time.Second
		return duration
	}
	policy := common.NewRetryPolicy(5, shouldRetry, exponentialBackoff)
	return &policy
}
