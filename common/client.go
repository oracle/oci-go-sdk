package common

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
	"os/user"
)

const (
	DefaultHostUrlTemplate   = "%s.%s.oraclecloud.com"
	defaultScheme            = "https"
	defaultSDKMarker         = "Oracle-GoSDK"
	defaultUserAgentTemplate = "%s/%s (%s/%s; go/%s)" //SDK/SDKVersion (OS/OSVersion; Lang/LangVersion)
	defaultTimeout           = time.Second * 30
	defaultConfigFileName    = "config"
	defaultConfigDirName     = ".oci"
	secondorayConfigDirName  = ".oraclebmc"
)

type RequestInterceptor func(*http.Request) error

// HttpRequestor wraps the execution of a http request, it is generally implemented by
// http.Client.Do, but can be customized for testing
type HttpRequestDispatcher interface {
	Do(req *http.Request) (*http.Response, error)
}

//BaseClient struct implements all basic operations to call oci web services.
type BaseClient struct {
	//HttpClient performs the http network operations
	HttpClient HttpRequestDispatcher

	//Signer performs auth operation
	Signer HttpRequestSigner

	//A request interceptor can be used to customize the request before signing and dispatching
	Interceptor RequestInterceptor

	//The region of the client
	Region Region

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

func newBaseClient(signer HttpRequestSigner, dispatcher HttpRequestDispatcher, region Region) BaseClient {
	return BaseClient{
		UserAgent:   defaultUserAgent(),
		Region:      region,
		Interceptor: nil,
		Signer:      signer,
		HttpClient:  dispatcher,
	}
}

func defaultHttpDispatcher() http.Client {
	httpClient := http.Client{
		Timeout:   defaultTimeout,
		Transport: &http.Transport{},
	}
	return httpClient
}

func DefaultBaseClient(region Region, provider KeyProvider) BaseClient {
	dispatcher := defaultHttpDispatcher()
	signer := defaultRequestSigner(provider)
	return newBaseClient(signer, &dispatcher, region)
}

func DefaultBaseClientWithDispatcher(dispatcher HttpRequestDispatcher, region Region, provider KeyProvider) BaseClient {
	signer := defaultRequestSigner(provider)
	return newBaseClient(signer, dispatcher, region)
}

func DefaultBaseClientWithSigner(signer HttpRequestSigner, region Region) BaseClient {
	dispatcher := defaultHttpDispatcher()
	return newBaseClient(signer, &dispatcher, region)
}

// Create a new client with a configuration provider, the configuration provider
// will be used for the default signer as well as reading the region
func NewClientWithConfig(configProvider ConfigurationProvider) (client BaseClient, err error) {
	var ok bool
	if ok, err = IsConfigurationProviderValid(configProvider); !ok {
		err = fmt.Errorf("can not create client, bad configuration: %s", err.Error())
		return
	}

	regionstr, _ := configProvider.Region()
	region, err := StringToRegion(regionstr)
	if err != nil {
		err = fmt.Errorf("can not create client, bad region configuration: %s", err.Error())
		return
	}

	client = DefaultBaseClient(region, configProvider)
	return
}

func getHomeFolder() string {
	current, e := user.Current()
	if e != nil {
		//Give up and try to return something sensible
		home :=  os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return current.HomeDir
}

//Create a new default client for a given region, with a default config provider
func NewClientForRegion(region Region) (client BaseClient) {
	homeFolder := getHomeFolder()
	defaultConfigFile := path.Join(homeFolder, defaultConfigDirName, defaultConfigFileName)
	secondaryConfigFile := path.Join(homeFolder, secondorayConfigDirName, defaultConfigFileName)

	defaultFileProvider, _ := ConfigurationProviderFromFile(defaultConfigFile, "")
	secondaryFileProvider, _ := ConfigurationProviderFromFile(secondaryConfigFile, "")
	environmentProvider := environmentConfigurationProvider{EnvironmentVariablePrefix: "TF_VAR"}

	provider, _ := ComposingConfigurationProvider([]ConfigurationProvider{defaultFileProvider, secondaryFileProvider, environmentProvider})
	Debugf("Configuration provided by: %s", provider)
	return DefaultBaseClient(region, provider)
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

	clientUrl, err := url.Parse(client.Host)
	if err != nil {
		return fmt.Errorf("host is invalid. %s", err.Error())
	}
	request.URL.Host = clientUrl.Host
	request.URL.Scheme = clientUrl.Scheme
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
		if dump, e := httputil.DumpRequest(request, true); e == nil {
			Logf("Dump Request %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	//Execute the http request
	response, err = client.HttpClient.Do(request)

	IfDebug(func() {
		if err != nil {
			Logln(err)
			return
		}

		if dump, e := httputil.DumpResponse(response, true); e == nil {
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

//Closes the body of an http reponse if the response and the body are valid
func CloseBodyIfValid(httpResponse *http.Response) {
	if httpResponse != nil && httpResponse.Body != nil {
		httpResponse.Body.Close()
	}
}
