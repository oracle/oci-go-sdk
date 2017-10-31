package common

import (
	"reflect"
	"time"
	"fmt"
	"strings"
)

func String(value string) *string {
	return &value
}

func Int(value int) *int {
	return &value
}

func Uint(value uint) *uint {
	return &value
}

func Float32(value float32) *float32 {
	return &value
}

func Float64(value float64) *float64 {
	return &value
}

func Bool(value bool) *bool {
	return &value
}

//Extracts the values of the pointer fields in the structs
func PointerString(datastruct interface{}) (representation string) {
	val := reflect.ValueOf(datastruct)
	typ := reflect.TypeOf(datastruct)
	all := make([]string, 2)
	all = append(all, "{")
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)

		//unexported
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		sv := val.Field(i)
		stringValue := ""
		if isNil(sv) {
			stringValue = fmt.Sprintf("%s=<nil>", sf.Name)
		} else {
			if sv.Type().Kind() == reflect.Ptr {
				sv = sv.Elem()
			}
			stringValue = fmt.Sprintf("%s=%v", sf.Name, sv)
		}
		all = append(all, stringValue)
	}
	all = append(all, "}")
	representation = strings.TrimSpace(strings.Join(all, " "))
	return
}


// A time struct, which renders to/from json using RFC339
type SDKTime struct {
	time.Time
}

//Returns and SDKTime from a time.Time struct
func SDKTimeFromTime(t time.Time) SDKTime {
	return SDKTime{t}
}

//Returns the time.Now() as an SDKTime pointer
func Now() *SDKTime {
	t := SDKTime{time.Now()}
	return &t
}

var timeType = reflect.TypeOf(SDKTime{})

const sdkTimeFormat = time.RFC3339

func FormatTime(t SDKTime) string {
	return t.Format(sdkTimeFormat)
}

func (t *SDKTime) UnmarshalJSON(data []byte) (e error) {
	s := string(data)
	if s == "null" {
		t.Time = time.Time{}
	} else {
		t.Time, e = time.Parse(`"`+sdkTimeFormat+`"`, string(data))
	}
	return
}

func (t *SDKTime) MarshalJSON() (buff []byte, e error) {
	s := t.Format(sdkTimeFormat)
	buff = []byte(`"` + s + `"`)
	return
}
