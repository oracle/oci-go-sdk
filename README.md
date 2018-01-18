# Oracle Cloud Infrastructure Golang SDK Preview
[![wercker status](https://app.wercker.com/status/09bc4818e7b1d70b04285331a9bdbc41/s/master "wercker status")](https://app.wercker.com/project/byKey/09bc4818e7b1d70b04285331a9bdbc41)

This is a ***preview*** of the official go sdk for Oracle Cloud Infrastructure
> ***WARNING:***: To avoid breaking changes please consider using [go dep](https://github.com/golang/dep), or vendoring this sdk.
Although we make a conscious effort to not push breaking changes, sometimes they are needed. This is particularly true at the current stage 
of development.


## Dependencies
- Install [Golang](https://golang.org/dl/)
- Install [make](https://www.gnu.org/software/make/), through your favorite package manager or binary 
distribution of choice

Additionally, if you are planning to  build or develop the sdk you need the following dependencies
- Install [github.com/stretchr/testify](https://github.com/stretchr/testify)
```sh
go get github.com/stretchr/testify
```


- Install [go lint](https://github.com/golang/lint)
```
go get -u github.com/golang/lint/golint
```
 


## Installing
Simply clone this repo into your go sdk you can use the following command

```
git clone git@github.com:oracle/oci-go-sdk.git  $GOPATH/src/github.com/oracle
```

## Configuring the SDK
You can configure the sdk with your credentials by creating a settings file in:
 $HOME/.oci/config
 ```
 [DEFAULT]
 user=[user ocid]
 fingerprint=[fingerprint]
 key_file=[path to pem file containing private key]
 tenancy=[tenancy ocid]
 region=[region for your tenancy]
 ```

## Making a request
Here is a quick example showing how to make a request to the identity service
```
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListAvailabilityDomainsRequest{}
	response, err := c.ListAvailabilityDomains(context.Background(), request)
	fmt.Println(response.Items)
```

## Organization of the SDK

## Building and testing
Building is provided by the make file at the root of the project. To build the project execute

```
make build
```

To execute the tests:
```
make test
```

