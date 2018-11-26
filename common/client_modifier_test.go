package common

import (
	"crypto/tls"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClientModifier_NoInitialModifier(t *testing.T) {
	modifier := NewClientModifier(nil)
	initialClient := DefaultBaseClientWithSigner(NewNOOPSigner())
	returnClient, err := modifier.Modify(&initialClient)
	assert.Nil(t, err)
	assert.ObjectsAreEqual(initialClient, returnClient)
}

func TestNewClientModifier_InitialModifier(t *testing.T) {
	modifier := NewClientModifier(setCustomCAPool)
	initialClient := DefaultBaseClientWithSigner(NewNOOPSigner())
	returnClient, err := modifier.Modify(&initialClient)
	assert.Nil(t, err)
	assert.Equal(t, returnClient.Transport.TLSClientConfig.ServerName, "test")
}

func TestNewClientModifier_ModifierFails(t *testing.T) {
	modifier := NewClientModifier(modifierGoneWrong)
	initialClient := DefaultBaseClientWithSigner(NewNOOPSigner())
	returnClient, err := modifier.Modify(&initialClient)
	assert.NotNil(t, err)
	assert.Nil(t, returnClient)
}

func setCustomCAPool(client *BaseClient) (*BaseClient, error) {
	client.Transport.TLSClientConfig = &tls.Config{
		ServerName: "test",
	}
	return client, nil
}

func modifierGoneWrong(client *BaseClient) (*BaseClient, error) {
	return nil, errors.New("uh oh")
}
