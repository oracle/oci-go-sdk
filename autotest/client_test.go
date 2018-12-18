package autotest

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

type creature interface {
	Habitat() string
}

type bird struct {
	Wings   int
	Element string
}

func (b bird) Habitat() string {
	return b.Element
}

type requestWrapper struct {
	ContainerId string
	Request     struct {
		RequestId string
		Creature  creature
	}
}

func TestConditionalStructCopy(t *testing.T) {
	containerId := "aaa"
	reqId := "id-1"

	testData := []struct {
		name         string
		srcData      map[string]interface{}
		polyInfo     map[string]PolymorphicRequestUnmarshallingInfo
		expectedType interface{}
	}{
		{
			name: "simple polymorphic type. unmarshaling the content of src data ",
			srcData: map[string]interface{}{
				"RequestId": reqId,
				"Creature": map[string]interface{}{
					"type":    "bird",
					"wings":   2,
					"element": "air",
				},
			},
			polyInfo: map[string]PolymorphicRequestUnmarshallingInfo{
				"Creature": {
					DiscriminatorName: "type",
					DiscriminatorValuesAndTypes: map[string]interface{}{
						"bird": &bird{},
					},
				},
			},
			expectedType: &bird{},
		},
	}

	for _, tIO := range testData {
		t.Run(tIO.name, func(t *testing.T) {
			rw := &requestWrapper{}
			reqJsonInfo := map[string]interface{}{
				"ContainerId": containerId,
				"Request":     tIO.srcData,
			}
			err := conditionalStructCopy(reqJsonInfo, rw, tIO.polyInfo, log.New(ioutil.Discard, "", 0))
			assert.NoError(t, err)
			assert.Equal(t, containerId, rw.ContainerId)
			assert.Equal(t, reqId, rw.Request.RequestId)
			assert.IsType(t, tIO.expectedType, rw.Request.Creature)
		})
	}
}
