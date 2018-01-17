package integtest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"time"
	"math/rand"
	"io"
	"crypto/sha256"
	"encoding/hex"
)

const (
	GoSDK2_Test_Prefix   = "GOSDK2_Test_"
	ENV_TENANCY_OCID     = "tenancy_ocid"
	ENV_USER_OCID        = "user_ocid"
	ENV_COMPARTMENT_OCID = "compartment_ocid"
	ENV_GROUP_OCID       = "group_ocid"
	ENV_REGION           = "region"

	DEF_ROOT_COMPARTMENT_ID = "ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu"
	DEF_USER_ID             = "ocid1.user.oc1..aaaaaaaav6gsclr6pd4yjqengmriylyck55lvon5ujjnhkok5gyxii34lvra"
	DEF_COMPARTMENT_ID      = "ocid1.compartment.oc1..aaaaaaaa5dvrjzvfn3rub24nczhih3zb3a673b6tmbvpng3j5apobtxshlma"
	DEF_GROUP_ID            = "ocid1.group.oc1..aaaaaaaayvxomawkk23wkp32cgdufufgqvx62qanmbn6vs3lv65xuc42r5sq"
	DEF_REGION              = common.RegionPHX
)

func getEnvSetting(s string, defaultValue string) string {
	val := os.Getenv("TF_VAR_" + s)
	if val != "" {
		return val
	}
	val = os.Getenv("OCI_" + s)
	if val != "" {
		return val
	}
	val = os.Getenv(s)
	if val != "" {
		return val
	}
	return defaultValue
}

// For now we will just use the default value if the env var is not set
// Eventually, certain values (TenancyID, UserID, etc.) may be required
//func getRequiredEnvSetting(s string) string {
//	val := getEnvSetting(s, "")
//  if val == "" {
//		panic(fmt.Sprintf("Required env setting %s is missing", s))
//	}
//	return v
//}

func getTenancyID() string {
	return getEnvSetting(ENV_TENANCY_OCID, DEF_ROOT_COMPARTMENT_ID)
}

func getUserID() string {
	return getEnvSetting(ENV_USER_OCID, DEF_USER_ID)
}

func getCompartmentID() string {
	return getEnvSetting(ENV_COMPARTMENT_OCID, DEF_COMPARTMENT_ID)
}

func getGroupID() string {
	return getEnvSetting(ENV_GROUP_OCID, DEF_GROUP_ID)
}

func getRegion() common.Region {
	region := getEnvSetting(ENV_REGION, "")

	if region != "" {

		r, err := common.StringToRegion(region)
		if err != nil {
			panic(err)
		}
		return r
	}

	return DEF_REGION
}

//Panics on error
func panicIfError(t *testing.T, err error) {
	if err != nil {
		t.Fail()
		panic(err)
	}
}

//Fails the test on error
func failIfError(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
		t.FailNow()
	}
}

// Retries a function until the predicate is true or it reaches a timeout.
// The operation is retried at the give frequency
func retryUntilTrueOrError(operation func() (interface{}, error), predicate func(interface{}) (bool, error), frequency, timeout <-chan time.Time) error {
	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout reached")
		case <-frequency:
			result, err := operation()
			if err != nil {
				return err
			}

			isTrue, err := predicate(result)
			if err != nil {
				return err
			}

			if isTrue {
				return nil
			}
		}
	}
}

//Finds lifecycle value inside the struct based on reflection
func findLifecycleFieldValue(request interface{}) (string, error) {
	val := reflect.ValueOf(request)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "", fmt.Errorf("can not unmarshal to response a pointer to nil structure")
		}
		val = val.Elem()
	}

	var err error
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		if err != nil {
			return "", err
		}

		sf := typ.Field(i)

		//unexported
		if sf.PkgPath != "" {
			continue
		}

		sv := val.Field(i)

		if sv.Kind() == reflect.Struct {
			lif, err := findLifecycleFieldValue(sv.Interface())
			if err == nil {
				return lif, nil
			}
		}
		if !strings.Contains(strings.ToLower(sf.Name), "lifecycle") {
			continue
		}
		return sv.String(), nil
	}
	return "", fmt.Errorf("request does not have a lifecycle field")
}

//Returns a function that checks for that a struct has the given lifecycle
func checkLifecycleState(lifecycleState string) func(interface{}) (bool, error) {
	return func(request interface{}) (bool, error) {
		fieldLifecycle, err := findLifecycleFieldValue(request)
		if err != nil {
			return false, err
		}
		isEqual := fieldLifecycle == lifecycleState
		return isEqual, nil
	}
}

func removeFileFn(filename string) func() {
	return func() {
		os.Remove(filename)
	}
}

func writeTempFile(data string) (filename string) {
	f, _ := ioutil.TempFile("", "gosdkTestintegtest")
	defer f.Close()
	f.WriteString(data)
	filename = f.Name()
	return
}

func writeTempFileOfSize(filesize int64) (fileName string, fileSize int64, contentHash string) {
	hash := sha256.New()
	f, _ := ioutil.TempFile("", "gosdkTestintegtest")
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	defer f.Close()
	writer := io.MultiWriter(f, hash)
	written, _ := io.CopyN(writer, ra, filesize)
	fileName = f.Name()
	fileSize = written
	contentHash = hex.EncodeToString(hash.Sum(nil))
	return
}



func getUuid() string {
	output, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	trimmed := strings.TrimSuffix(string(output), "\n")
	return trimmed
}

func getUniqueName(base string) string {
	return fmt.Sprintf("%s%s%s", GoSDK2_Test_Prefix, base, getUuid())
}

func verifyResponseIsValid(t *testing.T, response interface{}, err error) {
	assert.NotEmpty(t, response, fmt.Sprint(response))
	assert.NoError(t, err)
}

func createTestUser(client identity.IdentityClient) (identity.User, error) {
	req := identity.CreateUserRequest{}
	req.CompartmentId = common.String(getTenancyID())
	req.Name = common.String(getUniqueName("AUTG_User_"))
	req.Description = common.String("GoSDK Test User")
	rsp, err := client.CreateUser(context.Background(), req)
	return rsp.User, err
}

func deleteTestUser(client identity.IdentityClient, userID *string) error {
	req := identity.DeleteUserRequest{UserId: userID}
	return client.DeleteUser(context.Background(), req)
}

func validAD() string {
	c, _ := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	req := identity.ListAvailabilityDomainsRequest{}
	req.CompartmentId = common.String(getCompartmentID())
	response, _ := c.ListAvailabilityDomains(context.Background(), req)
	return *response.Items[0].Name
}

func readSampleFederationMetadata(t *testing.T) string {
	bytes, e := ioutil.ReadFile("sampleFederationMetadata.xml")
	failIfError(t, e)
	return string(bytes)
}

func createTestGroup(client identity.IdentityClient) (identity.Group, error) {
	req := identity.CreateGroupRequest{}
	req.CompartmentId = common.String(getTenancyID())
	req.Name = common.String(getUniqueName("Test_Group_"))
	req.Description = common.String("Go SDK Test Group")
	rsp, err := client.CreateGroup(context.Background(), req)
	return rsp.Group, err
}

func deleteTestGroup(client identity.IdentityClient, groupId *string) error {
	req := identity.DeleteGroupRequest{GroupId: groupId}
	return client.DeleteGroup(context.Background(), req)
}

