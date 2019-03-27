package autotest

import (
	"encoding/json"
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

type mamal struct {
	Legs    int
	Element string
}

func (m *mamal) UnmarshalJSON(data []byte) error {
	type fakeM mamal
	f := &fakeM{}
	e := json.Unmarshal(data, f)
	if e != nil {
		return e
	}

	m.Element = f.Element
	m.Legs = f.Legs
	return nil
}

func (m mamal) Habitat() string {
	return m.Element
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
			expectedType: bird{},
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

type aCreature struct {
	JsonData []byte
	Type     string `json:"type"`
}

func (m aCreature) Habitat() string {
	return "no habitat for this creature"
}

func (a *aCreature) UnmarshalJSON(data []byte) error {
	a.JsonData = data
	type fakeC aCreature
	aa := &fakeC{}
	e := json.Unmarshal(data, aa)
	if e != nil {
		return e
	}
	a.Type = aa.Type
	return nil
}

func (a *aCreature) UnmarshalPoly(data []byte) (interface{}, error) {
	switch a.Type {
	case "mamal":
		m := mamal{}
		err := json.Unmarshal(data, &m)
		return m, err
	}
	return *a, nil

}

type Details struct {
	DisplayName string                 `json:"displayName"`
	Variables   map[string]interface{} `mandatory:"false" json:"variables"`
	Creature    creature               `json:"creature"`
}

func (d *Details) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName string                 `json:"displayName"`
		Variables   map[string]interface{} `mandatory:"false" json:"variables"`
		Creature    aCreature              `json:"creature"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return e
	}
	d.DisplayName = model.DisplayName
	d.Variables = model.Variables

	n, e := model.Creature.UnmarshalPoly(model.Creature.JsonData)

	if e != nil {
		return e
	}

	if n != nil {
		d.Creature = n.(creature)
	} else {
		d.Creature = nil
	}
	return
}

func TestUnmarshalRequestWithEmbeddedJson(t *testing.T) {
	jsonEmbedded := `
		{
			"stackId": "someocid",
			"displayName": "Seattle",
			"description": "Northgate",
			"creature": {
				"type": "mamal",
				"legs": 4,
				"element": "forests"
			},
			"variables": {
				"settings": {
					"var_1": "val1",
					"var_2": "val2"
				}
			}
		}`

	type Request struct {
		StackId string `mandatory:"true" contributesTo:"path" name:"stackId"`
		Details `contributesTo:"body"`
	}

	l := log.New(ioutil.Discard, "TEST-LOG ", 0)
	r := Request{}
	err := unmarshalRequestEmbeddedJson(&r, []byte(jsonEmbedded), true, l)
	assert.NoError(t, err)
	assert.Equal(t, "someocid", r.StackId)
	assert.Equal(t, "Seattle", r.Details.DisplayName)
	assert.IsType(t, mamal{}, r.Details.Creature)

}

func TestUnmarshallingComplex(t *testing.T) {

	jsonValue1 := `[
	{
		"containerId": "55",
		"request": {
			"stackId": "someocid.dd.d",
			"displayName": "Seattle",
			"description": "Northgate",
			"creature": {
				"type": "mamal",
				"legs": 4,
				"element": "forests"
			},
			"configSource": {
				"configSourceType": "ZIP_UPLOAD",
				"zipFileBase64Encoded": "UEsDBBQACAgIAHVMRU4AAAAAAAAAAAAAAAAfAAAAcmVzb3VyY2VtYW5hZ2VyX25vX3Jlc291cmNlcy50Zl2RTU7EMAyF9zlFlD1Cs2E3Egg4AVukyJN4Wov8yUlDR6h3x20Z0LB8fs+f7aQDE5wCauNyLMAtYmo2O/JGfy2q/9qMA+X0r+iplgAXmyCiLYxnmreEYqx5YieRNIVgr9Joc4JKTkJK68K5UxUqsjYhOwh3OOOPqbUsFCF5fdQG3Zj1wUh5UTd0lkCOtjamNAi9QK2fmf3OCJiGNgrg8CCqFnQEQWTjCaWQOzKTR/vnmPvHd8HcDmmhynHUoaH9wIv4OEMsAfcpEIbM1Ma49r8+v7w9bf3bdX49DdiN1HF/GeWhgTBRACvwTCvHUKK2467OSqPUIchXKLV8A1BLBwhR4OTQ/AAAAKkBAABQSwECFAAUAAgICAB1TEVOUeDk0PwAAACpAQAAHwAAAAAAAAAAAAAAAAAAAAAAcmVzb3VyY2VtYW5hZ2VyX25vX3Jlc291cmNlcy50ZlBLBQYAAAAAAQABAE0AAABJAQAAAAA="
			},
			"variables": {
				"settings": {
					"var_1": "val1",
					"var_2": "val2"
				}
			}
		}
	}
]`

	type Request struct {
		StackId string `mandatory:"true" contributesTo:"path" name:"stackId"`
		Details `contributesTo:"body"`
	}

	type sampleReqInfo struct {
		ContainerId string
		Request     Request
	}

	l := log.New(ioutil.Discard, "TEST-LOG ", 0)
	var reqInfo []sampleReqInfo

	//umarshal to map
	var holder []map[string]interface{}
	err := json.Unmarshal([]byte(jsonValue1), &holder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(holder, &reqInfo, l)
	assert.NoError(t, err)

	assert.Len(t, reqInfo, 1)
	assert.Equal(t, "55", reqInfo[0].ContainerId)
	assert.Equal(t, "Seattle", reqInfo[0].Request.DisplayName)
	assert.IsType(t, mamal{}, reqInfo[0].Request.Creature)

}
