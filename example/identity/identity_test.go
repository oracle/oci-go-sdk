package identity

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	// create a compartment and return it's ID, if the compartment already exists
	// just return the id of it
	compartmentID, err := CreateCompartment()

	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Printf("Create compartment succeed with Compartment ID: %s", *compartmentID)
	}
}
