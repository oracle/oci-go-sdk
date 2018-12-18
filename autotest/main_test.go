package autotest

import (
	"os"
	"testing"
)

var testClient *OCITestClient
var testConfig *TestConfiguration

func TestMain(m *testing.M) {

	var err error
	testConfig, err = newTestConfiguration()
	if err != nil {
		panic(err)
	}

	testClient = NewOCITestClient()
	err = testClient.startSession()
	if err != nil {
		panic(err)
	}
	runResult := m.Run()

	err = testClient.endSession()
	if err != nil {
		panic(err)
	}
	os.Exit(runResult)
}
