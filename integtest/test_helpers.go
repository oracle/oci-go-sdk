package integtest

import (
	"testing"
	"os/exec"
	"fmt"
	"os"
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

var (
	ENV_TENANCY_OCID 		 = "tenancy_ocid"
	ENV_USER_OCID 			 = "user_ocid"
	ENV_COMPARTMENT_OCID	 = "compartment_ocid"
	ENV_GROUP_OCID			 = "group_ocid"
	ENV_REGION				 = "region"

	DEF_ROOT_COMPARTMENT_ID  = "ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu"
	DEF_USER_ID              = "ocid1.user.oc1..aaaaaaaav6gsclr6pd4yjqengmriylyck55lvon5ujjnhkok5gyxii34lvra"
	DEF_COMPARTMENT_ID     	 = "ocid1.compartment.oc1..aaaaaaaa5dvrjzvfn3rub24nczhih3zb3a673b6tmbvpng3j5apobtxshlma"
	DEF_GROUP_ID           	 = "ocid1.group.oc1..aaaaaaaayvxomawkk23wkp32cgdufufgqvx62qanmbn6vs3lv65xuc42r5sq"
	DEF_REGION				 = common.REGION_PHX
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
// Eventually, certain values (TenancyID, UserID, etc.) will be required
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

func panicIfError(t *testing.T, err error) {
	if err != nil {
		t.Fail()
		panic(err)
	}
}

func getUuid() string {
	output, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", output)
}

func getUniqueName(base string) string {
	return fmt.Sprintf("%s%s", base, getUuid())
}