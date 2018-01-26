package example

import (
	"testing"
)

type ExampleConfig struct {
	CompartmentName string
}

var Config = ExampleConfig{
	CompartmentName: "OCI_GOSDK_Sample_Compartment",
}

func RunAllTests(m *testing.T) {}
