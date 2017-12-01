package common

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"reflect"
	"strings"
	"time"
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

//PointerString prints the values of pointers in a struct
//Producing a human friendly string for an struct with pointers.
//useful when debugging the values of a struct
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

//Returns a SDKTime from a time.Time struct
func SDKTimeFromTime(t time.Time) SDKTime {
	return SDKTime{t}
}

//Returns the time.Now() as a SDKTime pointer
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

// PrivateKeyFromBytes is a helper function that will produce a RSA private
// key from bytes.
func PrivateKeyFromBytes(pemData []byte, password *string) (key *rsa.PrivateKey, e error) {
	if pemBlock, _ := pem.Decode(pemData); pemBlock != nil {
		decrypted := pemBlock.Bytes
		if x509.IsEncryptedPEMBlock(pemBlock) {
			if password == nil {
				e = fmt.Errorf("private_key_password is required for encrypted private keys")
				return
			}
			if decrypted, e = x509.DecryptPEMBlock(pemBlock, []byte(*password)); e != nil {
				return
			}
		}

		key, e = x509.ParsePKCS1PrivateKey(decrypted)

	} else {
		e = fmt.Errorf("PEM data was not found in buffer")
		return
	}
	return
}
