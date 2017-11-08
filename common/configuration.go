package common

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ConfigurationProvider wraps information about the account owner
type ConfigurationProvider interface {
	KeyProvider
	TenancyOCID() (string, error)
	UserOCID() (string, error)
	KeyFingerPrint() (string, error)
	Region() (string, error)
}


// Test to each part of the configration provider does not return an error
// If it does the function return false and the first error that caused the failure
func IsConfigurationProviderValid(conf ConfigurationProvider) (ok bool, err error) {
	baseFn := []func()(string, error){conf.TenancyOCID, conf.UserOCID, conf.KeyFingerPrint, conf.Region, conf.KeyID}
	for _, fn := range baseFn {
		_, err = fn()
		ok = err == nil
		if err != nil {
			return
		}
	}

	_, err = conf.PrivateRSAKey()
	ok = err == nil
	if err != nil {
		return
	}
	return true, nil
}

// environmentConfigurationProvider reads configuration from environment variables
type environmentConfigurationProvider struct {
	PrivateKeyPassword        string
	EnvironmentVariablePrefix string
}

// Create a ConfigurationProvider from a uniform set of environment variables starting with a prefix
// The env variables should look like: prefix_private_key_path, prefix_tenancy_ocid, prefix_user_ocid, prefix_fingerprint
// prefix_region
func ConfigurationProviderEnvironmentVariables(environmentVariablePrefix, privateKeyPassword string) ConfigurationProvider {
	return environmentConfigurationProvider{EnvironmentVariablePrefix: environmentVariablePrefix,
		PrivateKeyPassword: privateKeyPassword}
}

func (p environmentConfigurationProvider) String() string {
	return fmt.Sprintf("Configuration provided by environment variables prefixed with: %s", p.EnvironmentVariablePrefix)
}


// PrivateKeyFromBytes is a helper function that will produce a RSA private
// key from bytes.
func privateKeyFromBytes(pemData []byte, password *string) (key *rsa.PrivateKey, e error) {
	if pemBlock, _ := pem.Decode(pemData); pemBlock != nil {
		decrypted := pemBlock.Bytes
		if x509.IsEncryptedPEMBlock(pemBlock) {
			if password == nil {
				e = fmt.Errorf("private_key_password is required for encrypted private keys")
				return
			}
			if decrypted, e = x509.DecryptPEMBlock(pemBlock, []byte(*password)); e != nil {
				return
			}
		}

		key, e = x509.ParsePKCS1PrivateKey(decrypted)

	} else {
		e = fmt.Errorf("PEM data was not found in buffer")
		return
	}
	return
}

func (p environmentConfigurationProvider) PrivateRSAKey() (key *rsa.PrivateKey, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "private_key_path")
	var ok bool
	var value string
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		return nil, fmt.Errorf("Can not read PrivateKey from env variable: %s", environmentVariable)
	}

	pemFileContent, err := ioutil.ReadFile(value)
	if err != nil {
		Debugln("Can not read PrivateKey location from environment variable: " + environmentVariable)
		return
	}

	key, err = privateKeyFromBytes(pemFileContent, &p.PrivateKeyPassword)
	return
}

func (p environmentConfigurationProvider) KeyID() (keyID string, err error) {
	ocid, err := p.TenancyOCID()
	if err != nil {
		return
	}

	userocid, err := p.UserOCID()
	if err != nil {
		return
	}

	fingerPrint, err := p.KeyFingerPrint()
	if err != nil {
		return
	}

	return fmt.Sprintf("%s/%s/%s", ocid, userocid, fingerPrint), nil
}

func (p environmentConfigurationProvider) TenancyOCID() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "tenancy_ocid")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read Tenancy from environment variable %s", environmentVariable)
	}
	return
}

func (p environmentConfigurationProvider) UserOCID() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "user_ocid")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read user id from environment variable %s", environmentVariable)
	}
	return
}

func (p environmentConfigurationProvider) KeyFingerPrint() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "fingerprint")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read fingerprint from environment variable %s", environmentVariable)
	}
	return
}

func (p environmentConfigurationProvider) Region() (value string, err error) {
	environmentVariable := fmt.Sprintf("%s_%s", p.EnvironmentVariablePrefix, "region")
	var ok bool
	if value, ok = os.LookupEnv(environmentVariable); !ok {
		err = fmt.Errorf("Can not read region from environment variable %s", environmentVariable)
	}
	return
}

var configurationFileInfo *configFileInfo

// fileConfigurationProvider reads configuration information from a file
// Does not support multiple profiles
type fileConfigurationProvider struct {
	//The path to the configuration file
	ConfigPath string

	//The password for the private key
	PrivateKeyPassword string
}

// Creates a ConfigurationProvider from a file
func ConfigurationProviderFromFile(configFilePath, privateKeyPassword string) (ConfigurationProvider, error) {
	if configFilePath == "" {
		return nil, fmt.Errorf("config file path can not be empty")
	}

	return fileConfigurationProvider{ConfigPath: configFilePath, PrivateKeyPassword: privateKeyPassword}, nil
}

type configFileInfo struct {
	UserOcid, Fingerprint, KeyFilePath, TenancyOcid, Region string
	PresentConfiguration                                    byte
}

const (
	hasTenancy = 1 << iota
	hasUser
	hasFingerPrint
	hasRegion
	hasKeyFile
	none
)

func parseConfigFile(data []byte) (info *configFileInfo, err error) {
	var configurationPresent byte

	if len(data) == 0 {
		return nil, fmt.Errorf("configuration file content is empty")
	}

	info = &configFileInfo{}
	content := string(data)
	for _, line := range strings.Split(content, "\n") {
		if !strings.Contains(line, "=") {
			continue
		}

		splits := strings.Split(line, "=")
		switch key, value := strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1]); strings.ToLower(key) {
		case "user":
			configurationPresent = configurationPresent | hasUser
			info.UserOcid = value
		case "fingerprint":
			configurationPresent = configurationPresent | hasFingerPrint
			info.Fingerprint = value
		case "key_file":
			configurationPresent = configurationPresent | hasKeyFile
			info.KeyFilePath = value
		case "tenancy":
			configurationPresent = configurationPresent | hasTenancy
			info.TenancyOcid = value
		case "region":
			configurationPresent = configurationPresent | hasRegion
			info.Region = value
		}
	}
	info.PresentConfiguration = configurationPresent
	return
}

func openConfigFile(configFilePath string) (data []byte, err error) {
	data, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		err = fmt.Errorf("can not read config file: %s due to: %s", configFilePath, err.Error())
	}

	return
}

func (p fileConfigurationProvider) String() string {
	return fmt.Sprintf("Configuration provided by file: %s", p.ConfigPath)
}

func (p fileConfigurationProvider) readAndParseConfigFile() (info *configFileInfo, err error) {
	if configurationFileInfo != nil {
		return configurationFileInfo, nil
	}

	if p.ConfigPath == "" {
		return nil, fmt.Errorf("configuration path can not be empty")
	}

	data, err := openConfigFile(p.ConfigPath)
	if err != nil {
		err = fmt.Errorf("error while parsing config file: %s. Due to: %s", p.ConfigPath, err.Error())
		return
	}

	return parseConfigFile(data)
}

func presentOrError(value string, expectedConf, presentConf byte, confMissing string) (string, error) {
	if presentConf&expectedConf == expectedConf {
		return value, nil
	} else {
		return "", errors.New(confMissing + " configuration is missing from file")
	}
}

func (p fileConfigurationProvider) TenancyOCID() (value string, err error) {
	info, err := p.readAndParseConfigFile()
	if err != nil {
		err = fmt.Errorf("can not read tenancy configuration due to: %s", err.Error())
		return
	}

	value, err = presentOrError(info.TenancyOcid, hasTenancy, info.PresentConfiguration, "tenancy")
	return
}

func (p fileConfigurationProvider) UserOCID() (value string, err error) {
	info, err := p.readAndParseConfigFile()
	if err != nil {
		err = fmt.Errorf("can not read tenancy configuration due to: %s", err.Error())
		return
	}

	value, err = presentOrError(info.UserOcid, hasUser, info.PresentConfiguration, "user")
	return
}

func (p fileConfigurationProvider) KeyFingerPrint() (value string, err error) {
	info, err := p.readAndParseConfigFile()
	if err != nil {
		err = fmt.Errorf("can not read tenancy configuration due to: %s", err.Error())
		return
	}
	value, err = presentOrError(info.Fingerprint, hasFingerPrint, info.PresentConfiguration, "fingerprint")
	return
}

func (p fileConfigurationProvider) KeyID() (keyID string, err error) {
	info, err := p.readAndParseConfigFile()
	if err != nil {
		err = fmt.Errorf("can not read tenancy configuration due to: %s", err.Error())
		return
	}

	return fmt.Sprintf("%s/%s/%s", info.TenancyOcid, info.UserOcid, info.Fingerprint), nil
}

func (p fileConfigurationProvider) PrivateRSAKey() (key *rsa.PrivateKey, err error) {
	info, err := p.readAndParseConfigFile()
	if err != nil {
		err = fmt.Errorf("can not read tenancy configuration due to: %s", err.Error())
		return
	}

	filePath, err := presentOrError(info.KeyFilePath, hasKeyFile, info.PresentConfiguration, "key file path")
	if err != nil {
		return
	}
	pemFileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("can not read PrivateKey  from configuration file due to: %s", err.Error())
		return
	}

	key, err = privateKeyFromBytes(pemFileContent, &p.PrivateKeyPassword)
	return
}

func (p fileConfigurationProvider) Region() (value string, err error) {
	info, err := p.readAndParseConfigFile()
	if err != nil {
		err = fmt.Errorf("can not read region configuration due to: %s", err.Error())
		return
	}

	value, err = presentOrError(info.Region, hasRegion, info.PresentConfiguration, "region")
	return
}

// A configuration provider that look for information in  multiple configuration providers
type composingConfigurationProvider struct {
	Providers []ConfigurationProvider
}
// Creates a composing configuration provider with the given slice of configuration providers
// A composing provider will return the configuration of the first provider that has the configuration
// if no provider has the configuration it will return an error.
func ComposingConfigurationProvider(providers []ConfigurationProvider) (ConfigurationProvider, error) {
	if len(providers) == 0 {
		return nil, fmt.Errorf("providers can not be an empty slice")
	}
	return composingConfigurationProvider{Providers: providers}, nil
}

func (c composingConfigurationProvider) TenancyOCID() (string, error) {
	for _, p := range c.Providers {
		val, err := p.TenancyOCID()
		if err == nil {
			return val, nil
		}
	}
	return "", fmt.Errorf("did not find a proper configuration for tenancy")
}

func (c composingConfigurationProvider) UserOCID() (string, error) {
	for _, p := range c.Providers {
		val, err := p.UserOCID()
		if err == nil {
			return val, nil
		}
	}
	return "", fmt.Errorf("did not find a proper configuration for user")
}

func (c composingConfigurationProvider) KeyFingerPrint() (string, error) {
	for _, p := range c.Providers {
		val, err := p.KeyFingerPrint()
		if err == nil {
			return val, nil
		}
	}
	return "", fmt.Errorf("did not find a proper configuration for keyFingerprint")
}
func (c composingConfigurationProvider) Region() (string, error) {
	for _, p := range c.Providers {
		val, err := p.Region()
		if err == nil {
			return val, nil
		}
	}
	return "", fmt.Errorf("did not find a proper configuration for keyFingerprint")
}

func (c composingConfigurationProvider) KeyID() (string, error) {
	for _, p := range c.Providers {
		val, err := p.KeyID()
		if err == nil {
			return val, nil
		}
	}
	return "", fmt.Errorf("did not find a proper configuration for key id")
}

func (c composingConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	for _, p := range c.Providers {
		val, err := p.PrivateRSAKey()
		if err == nil {
			return val, nil
		}
	}
	return nil, fmt.Errorf("did not find a proper configuration for private key")
}
