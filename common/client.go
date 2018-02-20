// Package common Copyright (c) 2016, 2017, 2018 Oracle and/or its affiliates. All rights reserved.
package common

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

const (
	// DefaultHostURLTemplate The default url template for service hosts
	DefaultHostURLTemplate        = "%s.%s.oraclecloud.com"
	defaultScheme                 = "https"
	defaultSDKMarker              = "Oracle-GoSDK"
	defaultUserAgentTemplate      = "%s/%s (%s/%s; go/%s)" //SDK/SDKVersion (OS/OSVersion; Lang/LangVersion)
	defaultRequestTimeout         = 0
	defaultConnectionTimeout      = 10 * time.Second
	defaultTLSHandshakeTimeout    = 5 * time.Second
	defaultKeepaliveTimeout       = 30 * time.Second
	defaultConfigFileName         = "config"
	defaultConfigDirName          = ".oci"
	secondaryConfigDirName        = ".oraclebmc"
	maxBodyLenForDebug            = 1024 * 1000
	generatedRetryTokenLength     = 30
	absoluteMaximumRequestTimeout = 7 * 24 * time.Hour
)

// RequestInterceptor function used to customize the request before calling the underlying service
type RequestInterceptor func(*http.Request) error

// HTTPRequestDispatcher wraps the execution of a http request, it is generally implemented by
// http.Client.Do, but can be customized for testing
type HTTPRequestDispatcher interface {
	Do(req *http.Request) (*http.Response, error)
}

// BaseClient struct implements all basic operations to call oci web services.
type BaseClient struct {
	//HTTPClient performs the http network operations
	HTTPClient HTTPRequestDispatcher

	//Signer performs auth operation
	Signer HTTPRequestSigner

	//A request interceptor can be used to customize the request before signing and dispatching
	Interceptor RequestInterceptor

	//The host of the service
	Host string

	//The user agent
	UserAgent string

	//Base path for all operations of this client
	BasePath string
}

func defaultUserAgent() string {
	userAgent := fmt.Sprintf(defaultUserAgentTemplate, defaultSDKMarker, Version(), runtime.GOOS, runtime.GOARCH, runtime.Version())
	return userAgent
}

var clientCounter int64

func getNextSeed() int64 {
	newCounterValue := atomic.AddInt64(&clientCounter, 1)
	return newCounterValue + time.Now().UnixNano()
}

func newBaseClient(signer HTTPRequestSigner, dispatcher HTTPRequestDispatcher) BaseClient {
	rand.Seed(getNextSeed())
	return BaseClient{
		UserAgent:   defaultUserAgent(),
		Interceptor: nil,
		Signer:      signer,
		HTTPClient:  dispatcher,
	}
}

func defaultHTTPDispatcher() http.Client {
	httpClient := http.Client{
		Timeout: defaultRequestTimeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   defaultConnectionTimeout,
				KeepAlive: defaultKeepaliveTimeout,
				DualStack: true,
			}).DialContext,
			TLSHandshakeTimeout: defaultTLSHandshakeTimeout,
		},
	}
	return httpClient
}

func defaultBaseClient(provider KeyProvider) BaseClient {
	dispatcher := defaultHTTPDispatcher()
	signer := defaultRequestSigner(provider)
	return newBaseClient(signer, &dispatcher)
}

//DefaultBaseClientWithSigner creates a default base client with a given signer
func DefaultBaseClientWithSigner(signer HTTPRequestSigner) BaseClient {
	dispatcher := defaultHTTPDispatcher()
	return newBaseClient(signer, &dispatcher)
}

// NewClientWithConfig Create a new client with a configuration provider, the configuration provider
// will be used for the default signer as well as reading the region
func NewClientWithConfig(configProvider ConfigurationProvider) (client BaseClient, err error) {
	var ok bool
	if ok, err = IsConfigurationProviderValid(configProvider); !ok {
		err = fmt.Errorf("can not create client, bad configuration: %s", err.Error())
		return
	}

	regionstr, _ := configProvider.Region()
	_, err = StringToRegion(regionstr)
	if err != nil {
		err = fmt.Errorf("can not create client, bad region configuration: %s", err.Error())
		return
	}

	client = defaultBaseClient(configProvider)
	return
}

func getHomeFolder() string {
	current, e := user.Current()
	if e != nil {
		//Give up and try to return something sensible
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return current.HomeDir
}

// DefaultConfigProvider returns the default config provider. The default config provider
// will look for configurations in 3 places: file in $HOME/.oci/config, HOME/.obmcs/config and
// variables names starting with the string TF_VAR. If the same configuration are found in multiple
// places the provider will prefer the first one.
func DefaultConfigProvider() ConfigurationProvider {
	homeFolder := getHomeFolder()
	defaultConfigFile := path.Join(homeFolder, defaultConfigDirName, defaultConfigFileName)
	secondaryConfigFile := path.Join(homeFolder, secondaryConfigDirName, defaultConfigFileName)

	defaultFileProvider, _ := ConfigurationProviderFromFile(defaultConfigFile, "")
	secondaryFileProvider, _ := ConfigurationProviderFromFile(secondaryConfigFile, "")
	environmentProvider := environmentConfigurationProvider{EnvironmentVariablePrefix: "TF_VAR"}

	provider, _ := ComposingConfigurationProvider([]ConfigurationProvider{defaultFileProvider, secondaryFileProvider, environmentProvider})
	Debugf("Configuration provided by: %s", provider)
	return provider
}

func (client *BaseClient) prepareRequest(request *http.Request) (err error) {
	if client.UserAgent == "" {
		return fmt.Errorf("user agent can not be blank")
	}

	if request.Header == nil {
		request.Header = http.Header{}
	}
	request.Header.Set("User-Agent", client.UserAgent)

	if !strings.Contains(client.Host, "http") &&
		!strings.Contains(client.Host, "https") {
		client.Host = fmt.Sprintf("%s://%s", defaultScheme, client.Host)
	}

	clientURL, err := url.Parse(client.Host)
	if err != nil {
		return fmt.Errorf("host is invalid. %s", err.Error())
	}
	request.URL.Host = clientURL.Host
	request.URL.Scheme = clientURL.Scheme
	currentPath := request.URL.Path
	request.URL.Path = path.Clean(fmt.Sprintf("/%s/%s", client.BasePath, currentPath))
	return
}

func (client BaseClient) intercept(request *http.Request) (err error) {
	if client.Interceptor != nil {
		err = client.Interceptor(request)
	}
	return
}

func checkForSuccessfulResponse(res *http.Response) error {
	familyStatusCode := res.StatusCode / 100
	if familyStatusCode == 4 || familyStatusCode == 5 {
		return newServiceFailureFromResponse(res)
	}
	return nil

}

// GenerateRetryToken generates a retry token that must be included on any request passed to the Retry method.
func GenerateRetryToken() string {
	alphanumericChars := []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	retryToken := make([]rune, generatedRetryTokenLength)
	for i := range retryToken {
		retryToken[i] = alphanumericChars[rand.Intn(len(alphanumericChars))]
	}
	return string(retryToken)
}

// GetMaximumTimeout ensures that the policy MaximumTimeout (which can be set to 'unlimited') is still bounded by
// the OCI absolute maximum (currently set to 1 week).
func GetMaximumTimeout(policy *RetryPolicy) time.Duration {
	// even if a user says poll forever by specifying zero for the maximum timeout, we're going to stop them
	if policy.MaximumTimeout == 0 {
		return absoluteMaximumRequestTimeout
	}
	return policy.MaximumTimeout
}

type deadlineExceededByBackoffError struct{}

func (deadlineExceededByBackoffError) Error() string {
	return "now() + computed backoff duration exceeds request deadline"
}

// DeadlineExceededByBackoff is the error returned by Call() when GetNextDuration() returns a time.Duration that would
// force the user to wait past the request deadline before re-issuing a request. This enables us to exit early, since
// we cannot succeed based on the configured retry policy.
var DeadlineExceededByBackoff error = deadlineExceededByBackoffError{}

// OciRequest is any request made to an OCI service.
type OciRequest interface {
	// GetHttpRequest assembles an HTTP request.
	GetHttpRequest(method, path string) (http.Request, error)
}

// OciResponse is the response from issuing a request to an OCI service.
type OciResponse interface {
	// GetRawResponse returns the raw HTTP response.
	GetRawResponse() *http.Response
}

// OciOperation is the generalization of a request-response cycle undergone by an OCI service.
type OciOperation func(context.Context, OciRequest) (OciResponse, error)

// Retry executes the retryable request using the specified operation and retry policy options
func (client BaseClient) Retry(ctx context.Context, request OciRetryableRequest, operation OciOperation, options ...RetryPolicyOption) (OciResponse, error) {
	// Each operation defines the default retry behavior for a given request
	// Users can modify the default retry behavior by passing through a variadic number of retry policy options
	policy := request.GetRetryPolicy(options...)

	deadlineContext, deadlineCancel := context.WithTimeout(ctx, GetMaximumTimeout(&policy))
	defer deadlineCancel()

	for currentOperationAttempt := uint(1); ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		Debugln(fmt.Sprintf("operation attempt #%v", currentOperationAttempt))
		response, err := operation(deadlineContext, request)

		select {
		case <-deadlineContext.Done():
			// return why the request was aborted (could be user interrupted or deadline exceeded)
			// => include last received response for information (user may choose to re-issue request)
			return response, deadlineContext.Err()
		default:
			// non-blocking select
		}

		if policy.ShouldRetryOperation(response, err, currentOperationAttempt) {
			// this conditional is explicitly not added to the encompassing if condition to retry based on response
			// => it is only to determine if, on the last round of this loop, we still skip sleeping (if we're the
			//    last attempt, then there's no point sleeping before we round the loop again and fall out to the
			//    Maximum Number Attempts exceeded error)
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := deadlineContext.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					// we want to retry the operation, but the policy is telling us to wait for a duration that exceeds
					// the specified overall deadline for the operation => instead of waiting for however long that
					// time period is and then aborting, abort now and save the cycles
					return response, DeadlineExceededByBackoff
				}
				Debugln(fmt.Sprintf("waiting %v before retrying operation", duration))
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return response, err
		}
	}
	return nil, fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts)
}

// Call executes the http request with the given context
func (client BaseClient) Call(ctx context.Context, request *http.Request) (response *http.Response, err error) {
	Debugln("Atempting to call downstream service")
	request = request.WithContext(ctx)

	err = client.prepareRequest(request)
	if err != nil {
		return
	}

	//Intercept
	err = client.intercept(request)
	if err != nil {
		return
	}

	//Sign the request
	err = client.Signer.Sign(request)
	if err != nil {
		return
	}

	IfDebug(func() {
		dumpBody := true
		if request.ContentLength > maxBodyLenForDebug {
			Logln("not dumping body too big")
			dumpBody = false
		}
		if dump, e := httputil.DumpRequest(request, dumpBody); e == nil {
			Logf("Dump Request %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	//Execute the http request
	response, err = client.HTTPClient.Do(request)

	IfDebug(func() {
		if err != nil {
			Logln(err)
			return
		}

		dumpBody := true
		if response.ContentLength > maxBodyLenForDebug {
			Logln("not dumping body too big")
			dumpBody = false
		}

		if dump, e := httputil.DumpResponse(response, dumpBody); e == nil {
			Logf("Dump Response %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	if err != nil {
		return
	}

	err = checkForSuccessfulResponse(response)
	return
}

//CloseBodyIfValid closes the body of an http response if the response and the body are valid
func CloseBodyIfValid(httpResponse *http.Response) {
	if httpResponse != nil && httpResponse.Body != nil {
		httpResponse.Body.Close()
	}
}
