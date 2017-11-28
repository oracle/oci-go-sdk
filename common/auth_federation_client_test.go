package common

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	leafCertBody = `MIIHFTCCBP2gAwIBAgIRANNLkf6VdDV7+ewb7W/mSLYwDQYJKoZIhvcNAQELBQAw
gZ8xczBxBgNVBAsTam9wYy1kZXZpY2U6MTA6OWQ6Zjc6ZWM6NjE6OGE6NzQ6OTM6
ZGM6MGU6ZGY6YmQ6NDA6Yjc6NWY6ZjE6MTM6OTg6MmY6NDQ6ZmI6ZDM6MDY6OGE6
OTk6NmM6ZTE6MDQ6NDE6ZGY6NmE6MjExKDAmBgNVBAMTH1BLSVNWQyBJZGVudGl0
eSBJbnRlcm1lZGlhdGUgcjIwHhcNMTcxMTI3MjEwMTI2WhcNMTcxMTI3MjIwMTI2
WjCCAbAxggFOMBwGA1UECxMVb3BjLWNlcnR0eXBlOmluc3RhbmNlMGcGA1UECxNg
b3BjLWluc3RhbmNlOm9jaWQxLmluc3RhbmNlLm9jMS5waHguYWJ5aHFsanJ6Y29i
dmJ1dDdrbmlxZHk2c2Y0ZnA1NjU0dmI3ZGdiZ3g2ZHEzaDM3b3luamk0dTVuMmxx
MGoGA1UECxNjb3BjLWNvbXBhcnRtZW50Om9jaWQxLmNvbXBhcnRtZW50Lm9jMS4u
YWFhYWFhYWEyM2tubjJxaTIzdXJxcXZ2b3picm8zdTVneDJ1YmRhYnl5Z3ZlNng1
NmNpZWh1MnpibHZxMFkGA1UECxNSb3BjLXRlbmFudDpvY2lkdjE6dGVuYW5jeTpv
YzE6cGh4OjE0NjA0MDY1OTI2NjA6YWFhYWFhYWFiNGZhb2ZyZmt4ZWNvaGhqdWl2
anEyNjJwdTFcMFoGA1UEAxNTb2NpZDEuaW5zdGFuY2Uub2MxLnBoeC5hYnlocWxq
cnpjb2J2YnV0N2tuaXFkeTZzZjRmcDU2NTR2YjdkZ2JneDZkcTNoMzdveW5qaTR1
NW4ybHEwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDa+Nz0Mp7IRFKQ
qxb8MmRcMdArGcQBnHoGXMcp3K0rQm23KUSLs/mYPkws0/nFtmVYPTqt/zBlsEoh
zprX5mj8BtkmOaqDXQGpPqmXJuQnn6ZRnvalo2kkhkZMIS+CpU7LfcR96sZYpI9R
gLVkGvkYaLd6ZClYXCqVl++Wyu7vC1+H1K5dGhVi/XVYtOIEKThZWgTN8gcLNtz+
9wJTGpcnj2H9UgY6Ut7CNupHqUei9zS0YpHjM/qxft0plYYwWbv+T1uXbUXrZNu+
GvjZc2Xnxe9fQKAnYDiEP/rkoxgmALda93DECo3FraEiYkXtARrnqjGd09M22wGZ
pIh21RwPo8rZlcOtSEeS+U3VYDb78Sukud978EIEn2t9MZBH9/PbhVkANCBOVAjd
lN6+0nLItYPzRUs0LUsXVxxj9AJPFtIIaGQhZqOzpSaPuUtg5IkCoHy072FRPMpd
UhSUGtCuaTLAUZgM4OMCa9XcyTAdjP0X0vabXxCp+KFnHgr3ctF/xZB/N/GyNlUB
IdF02HUaDv9v86BRQPKpeuXKOpVnpOpBkwFBH9sITmFhn0mIC0pqv0/Jc5h8aQBr
HcsoZkt9TFr1Es8HD22a4AlgWW30pKOlfKei3+PfOpkFXKJslsgViIxTeLk/zjjB
HntyN0Zbxg+YrCaWgfxJpxdrmY/WvQIDAQABozgwNjATBgNVHSUEDDAKBggrBgEF
BQcDAjAfBgNVHSMEGDAWgBQKKSDx5jhmP8SLpYrtwMQsrYdKfDANBgkqhkiG9w0B
AQsFAAOCAgEAVARAHp9pyvBV4zVi+Zk1n5nrL7rNbTZWN23msCizvXLnqkjklqpU
3T3m1fuahpnwZDOkFb72OawSwikNz2oRV3wtkkQRTcQqmHconkBI8zAlAbybT7t9
92SmRPyy5/W/ud7dpWBm+M1uSa514JnxjF+2D8sCWDe1iqKLOE+gQjf9L84dUaM0
xlbFQbdofv6cvEEQS+u34790Kc/KlaOT9WL1fBatUBtBCxQmWwKjysJzf3508MCo
C+CLcKbXNuBWc25vIRE/7n6GCHWmGirEG/Ry+gS1jcTwCYEzT4dcUTPMCl+VYbmk
rRz0AQJeAwyuAOQaMwkLefJOB3eW9LvDFJXWlFxQofFiTpppourXbVooNLlf5Mba
m5Dag4iw4FqYNQCAsvn+TSLQ+tPKaBTQdxPIOGm+MDsJa3nzUCcr2olJ8VBC5/ty
8cPc8So/XDruAXrvps0DSiivOi3CslU4EcCjuC/hHToaMhdW1PaHgDVF7DM+2XDv
+PxGCxpStO/QDl3zg+/Pak7ecwdeqPmjkLqsaaQjqVEO69eu/baa/dXnWUymU+a7
ZvjctG/CWczlK5u597lz8GVCJRJM0/koqEk4vFiKF+qdrKCcxu3Vxmm+1PmD9UuY
SOoA5ojbEJuWIf5Y2gCeMqkFrx+o1LPC5/5kUe4l5OysF28vJ9J2dxI=`
	leafCertFingerPrint    = `05:76:1f:8b:9b:13:38:df:8e:fc:ae:74:2a:74:cc:6a:09:48:fe:0f`
	leafCertPrivateKeyBody = `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`
	intermediateCertBody = `MIIFoTCCA4mgAwIBAgIRAPoRkIqtdkVTpRV7VlpI260wDQYJKoZIhvcNAQELBQAw
MjEwMC4GA1UEAxMnT3JhY2xlIEJNQ1MgSW5zdGFuY2UgSWRlbnRpdHkgUm9vdCAt
IHIyMB4XDTE3MTEyNjE2MzE0M1oXDTE4MDIyNDE2MzE0M1owgZ8xczBxBgNVBAsT
am9wYy1kZXZpY2U6MTA6OWQ6Zjc6ZWM6NjE6OGE6NzQ6OTM6ZGM6MGU6ZGY6YmQ6
NDA6Yjc6NWY6ZjE6MTM6OTg6MmY6NDQ6ZmI6ZDM6MDY6OGE6OTk6NmM6ZTE6MDQ6
NDE6ZGY6NmE6MjExKDAmBgNVBAMTH1BLSVNWQyBJZGVudGl0eSBJbnRlcm1lZGlh
dGUgcjIwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCcyId/8YJTISnL
aC2mdnZaIDGbcofJb5BEZEPZTm1Nhqn0+EzA1UT67lQ+9DWqsqgY2wLFQHDQAMqa
UpnPqhLlI6p1ZVxK2HNyYlO2vNdWYyyentVD5xHi+cu35DHJch0MM/0H/oaHqGF3
71B9cRhl8Htd63NLvl3PcwdP6Eo1dPV4+oCGs7JbTBTyNex+engNyh+xnhxOifbO
ZAmZA9eFBD3qqSo/PMHx/Mjs5CsQhpmVKXnuiRf1AMGrwo0yYs2wwK5LhMyzHCf2
JBwdnQuikjyTNaC/E9iIMDgHLYC8BqLNAIH0XAQjX0pMxIfQ19IWVJmUJjzk5HqR
l9V4Ze6v4yIThC4U8jAoSel4I3OJ/hcgdx2jDnCSBPBU74/fms9JVy8NUyC0+XRn
E/LOQ74UXbhYegL88dJpbtHhdma3AM5RidEyFsExztNxuSW36V4mCDBmBk+Y+O9r
mzy4e6g/LKy2+676Hte1eWsfkN/uL1gYfA1n1nbLHpQa0Rsxy1HCXY61/tGWMqsx
4PT/cmr1b/XKrCZJRJ9fB8o6d+SQeHAjW2j0T2auBYKQGjmYyjcEWfiFT6BEw99B
LUWBJXpNptsW+XIpb8/LuxVwRC5qS8WMylOC8gSl2cL4iNRLTJ8Bp1GavePkpWOX
ww056s9/NucggJZwY4ERVxM1H80RuQIDAQABo0QwQjAOBgNVHQ8BAf8EBAMCArww
DwYDVR0TAQH/BAUwAwEB/zAfBgNVHSMEGDAWgBQKKSDx5jhmP8SLpYrtwMQsrYdK
fDANBgkqhkiG9w0BAQsFAAOCAgEALswft4QmyvnZlA/FmV+9va1f4w1ePmDkZAYS
xgCdvv2zO2dy680vnVTV1M6CtF38RjbTuLYm7QUtGASDsCK/FZlC1pXw1y2QlSwq
WvvBkfp0rpXNIAyVhF83K7I/Xclnzen6h0c48I2atXyBvM56STxb1YVLkQTCkTnl
8eQUjhbe1cRqSfCc376wHl5mMxANL4z3wDpoOlHHEc5pJLBMkDGyGG2Jk85IaNZQ
auhL5h3tfxFQir2YNwOrkI83zJoGiPbpK9UihelpWn3rwFykPtWFxGmWDj4+o21T
pE8teyGclDcAiZlS+57eiWF4/psP4i+QWTd6G97JQp/Kty1LVelUqYR7kPzr/Fj7
wUd7Teyk+OK4t1J2fZuk85T904C/yuNhA4+arEvDh5Y4yNx1sMhz2Aoa4UP8P4bI
/4As4JFm2J2Z/XxK1sWidND9kzZBN3mW+NlFCXcUBYLzmVGXw1bTarHGXzhB7SML
o0Etqx22KF+NwxgIa0fo1DcS8aSIc8RjKPDqYIMBjdRlWY1Rz37kn/QJ0ZsZSXuo
PZW21eEc5YTiPRg6X34AZ+b4YdxvHhuupNLzmAbA7nsZ9BPkNEYCn8jaY+uFA6PT
MT32Ndh36mTSRbiv5/wqBdHxp/lYWgquCEQG/9fcw2YV7Qvc0tNiToo03Yn/z3Bo
62lMOVg=`
	sessionPrivateKeyBody = `MIIEpAIBAAKCAQEAyQn1VgH+aBOfN8PfN6WUZOyiGWbozd1WaxDyP/rYCPJJIinK
upp1ZcissN+A2dgcknQqJteX9XWYz31WzeAkNEVAi9R4ZnrP0u/4921ZWmiCKIqB
VxSWGE+PJcHUWJRSFNS1mQcX//UjwEYNPDKVtPcwQqt7CYTxL77YFy0Z+s9WUmZa
OJakgrCLSokeQBWdi0JibYp1mZPZv6pqsIm9X86ef1hXyNjvEQRxuf1Bx96Y32m7
FjsD251XeOEzzdESCa90Z+bHN6k7wsTRrU79dYZF0puZUEmHID4xIF5AprOHVarr
hawiddwayMQWH7GZuVzhJ2Z/Q4CK2DneR8LrfwIDAQABAoIBAQCBOJpunzd0lHA5
+vfmn9KjkIqdA03OzV930Fu2DjmAfqPNhsldalzdgMCnq/H5lHTBGlPhxaBSUTVY
4TrDeDgtNPpJJtgBHx2oe0EvYSUW/Kf1ARoj7AKSO8lKPq2MBkJAS30ykKG3j01c
kGp/cBBUUrK9l9WtB7uzsdraAs55LFp61B3GLmIRcdbkJ1KcLYROVHw7gjdpy8jg
RIiEDYwhkD8p/du5Vc16Dn9f8UklTkHuQtvY8K186KJYlBngG5VNegqtSv8Id1fn
LO9iLQmQ3mWC/UpYggSl5MUkM9N/bWo/lDgw5mDJF1NW37fRwJF9ROGpTShH4YD8
YRIKNghxAoGBANqQ0Rd0TwWont6XdeO6pDsnmiQFcjvqvN2u6+jrruzPQZe6iR1u
R2XfAu/Bibm2NJEhRqeNaP64Wz+bB8I/2pIo2zZ5qVofCPjiba13p0mwwXolGeZy
Jf2bB2y4dQJgoydyg8VNv1A0nbNhN6yXchC/ve96IuA3KPDwYUQL9SIJAoGBAOt4
qm+e0Cwu4FHs1k583HeFPaWp3WrN9B7YOWNivq0nBNO1BVpP9e59emqEH+EF0szo
G2senOc11EcN7ErJgQgiOSPSNs/uov5jxyua1L3WyzX1blEOpPk2fno1ywRvlvR9
72SQL6jkEwQ68RzMm5Jxu8fShjYPIgadTB7Y9WNHAoGBAJF1j7JsVasObaz2YB8P
N/2mfY87kKsPrmJa5+Y79E7mIGE6Y8aSfjHuGaY2x9Iw0QtFeiBInHfpedD6/E7q
8CwVxM6caXjw7qe1CglIeK0yVZFU38fecCo34tkYaccLSYoXTFsjQw+99LZNHSBg
Kim7ckzOnCdcjoLLd/AxhRMBAoGAXr0t1bsQBrbViiDAGNQj3/K72ut9KnuyvLZC
YLpZ/VC1oHg0ryF80XeqyTBfoym0pry8S10FxADkZ7IyX+SzBZK/aqopY3+gFLoQ
aHrjHX5ORGd6+yHpPluuh30dMau0EKqapttcUZdyD0GnwO6RqVoZM0yAsz9jcUXY
WpDwKgECgYA+q2SGMaDlUDjCdU7mZxjuaS60aGDZSJtfF1o5TGKF7CqqiXHOqfif
8Mf5JuwQeSRIpmBG2YR5EBueyGXk3TRgK/eFd6L2qIfkLMQfZILBeNcrCoYAf6Gl
TjsJ0+JBWCEaXdU6zk4tSuyu15i9i1/yyzaaphasegFtxAP7hdWFVg==`
	sessionPublicKeyBody = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyQn1VgH+aBOfN8PfN6WU
ZOyiGWbozd1WaxDyP/rYCPJJIinKupp1ZcissN+A2dgcknQqJteX9XWYz31WzeAk
NEVAi9R4ZnrP0u/4921ZWmiCKIqBVxSWGE+PJcHUWJRSFNS1mQcX//UjwEYNPDKV
tPcwQqt7CYTxL77YFy0Z+s9WUmZaOJakgrCLSokeQBWdi0JibYp1mZPZv6pqsIm9
X86ef1hXyNjvEQRxuf1Bx96Y32m7FjsD251XeOEzzdESCa90Z+bHN6k7wsTRrU79
dYZF0puZUEmHID4xIF5AprOHVarrhawiddwayMQWH7GZuVzhJ2Z/Q4CK2DneR8Lr
fwIDAQAB`
	tenancyId             = `ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu`
	expectedSecurityToken = `eyJraWQiOiJhc3ciLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJvY2lkMS5pbnN0YW5jZS5vYzEucGh4LmFieWhxbGpyemNvYnZidXQ3a25pcWR5NnNmNGZwNTY1NHZiN2RnYmd4NmRxM2gzN295bmppNHU1bjJscSIsIm9wYy1jZXJ0dHlwZSI6Imluc3RhbmNlIiwiaXNzIjoiYXV0aFNlcnZpY2Uub3JhY2xlLmNvbSIsImZwcmludCI6IjA1Ojc2OjFGOjhCOjlCOjEzOjM4OkRGOjhFOkZDOkFFOjc0OjJBOjc0OkNDOjZBOjA5OjQ4OkZFOjBGIiwicHR5cGUiOiJpbnN0YW5jZSIsImF1ZCI6Im9wYy5vcmFjbGUuY29tIiwidHR5cGUiOiJ4NTA5Iiwib3BjLWluc3RhbmNlIjoib2NpZDEuaW5zdGFuY2Uub2MxLnBoeC5hYnlocWxqcnpjb2J2YnV0N2tuaXFkeTZzZjRmcDU2NTR2YjdkZ2JneDZkcTNoMzdveW5qaTR1NW4ybHEiLCJleHAiOjE1MTE4Mzg3OTMsIm9wYy1jb21wYXJ0bWVudCI6Im9jaWQxLmNvbXBhcnRtZW50Lm9jMS4uYWFhYWFhYWEyM2tubjJxaTIzdXJxcXZ2b3picm8zdTVneDJ1YmRhYnl5Z3ZlNng1NmNpZWh1MnpibHZxIiwiaWF0IjoxNTExODE3MTkzLCJqdGkiOiIxYTBkMTAwZi1mYjBjLTQzYjktOTE4Ni1iN2Y0MzIyNzJiY2UiLCJ0ZW5hbnQiOiJvY2lkdjE6dGVuYW5jeTpvYzE6cGh4OjE0NjA0MDY1OTI2NjA6YWFhYWFhYWFiNGZhb2ZyZmt4ZWNvaGhqdWl2anEyNjJwdSIsImp3ayI6IntcImtpZFwiOlwiMDU6NzY6MWY6OGI6OWI6MTM6Mzg6ZGY6OGU6ZmM6YWU6NzQ6MmE6NzQ6Y2M6NmE6MDk6NDg6ZmU6MGZcIixcIm5cIjpcIkFNa0o5VllCX21nVG56ZkQzemVsbEdUc29obG02TTNkVm1zUThqXzYyQWp5U1NJcHlycWFkV1hJckxEZmdObllISkowS2liWGxfVjFtTTk5VnMzZ0pEUkZRSXZVZUdaNno5THYtUGR0V1Zwb2dpaUtnVmNVbGhoUGp5WEIxRmlVVWhUVXRaa0hGX18xSThCR0RUd3lsYlQzTUVLcmV3bUU4Uy0tMkJjdEdmclBWbEptV2ppV3BJS3dpMHFKSGtBVm5ZdENZbTJLZFptVDJiLXFhckNKdlZfT25uOVlWOGpZN3hFRWNibjlRY2ZlbU45cHV4WTdBOXVkVjNqaE04M1JFZ212ZEdmbXh6ZXBPOExFMGExT19YV0dSZEtibVZCSmh5QS1NU0JlUUthemgxV3E2NFdzSW5YY0dzakVGaC14bWJsYzRTZG1mME9BaXRnNTNrZkM2MzhcIixcImVcIjpcIkFRQUJcIixcImt0eVwiOlwiUlNBXCIsXCJhbGdcIjpcIlJTMjU2XCIsXCJ1c2VcIjpcInNpZ1wifSIsIm9wYy10ZW5hbnQiOiJvY2lkdjE6dGVuYW5jeTpvYzE6cGh4OjE0NjA0MDY1OTI2NjA6YWFhYWFhYWFiNGZhb2ZyZmt4ZWNvaGhqdWl2anEyNjJwdSJ9.dHmjMqDKK___WmIyemg4rtjXw-UlZ_DvWzzPntOGONpWhcoKqhg3ihCTZIzL2ywEetIhoJASCHTy-EtdotntVGONeVyFwQEOgv0ctnrl4rQX3Tdwdrz5e4m_nD68WkKuVSxdA5TzVU73xwVsQGYIVT_8G68NmO1i9cLSCBJPG5xygmF8KG8orCzeU3zmiEOgyVYVPWTBhL4vFDNt7qmlCcNl_LEP4lOY9mH_OxENjNCtcwKfRCS_Ebzf3tdDJbjfz1h5vrQnt7v8U0aGxPf9tIMwb6X5pqXsqJXRcPiZDRPgA-V8WPb-ZvQ42VpWGZAo0zam19VX8akyo3K5Iwx3QQ`
)

var (
	leafCertPem                     = fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----\n", leafCertBody)
	leafCertBodyNoNewLine           = strings.Replace(leafCertBody, "\n", "", -1)
	leafCertPrivateKeyPem           = fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----\n", leafCertPrivateKeyBody)
	leafCertPrivateKeyBodyNoNewLine = strings.Replace(leafCertPrivateKeyBody, "\n", "", -1)
	intermediateCertPem             = fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----\n", intermediateCertBody)
	intermediateCertBodyNoNewLine   = strings.Replace(intermediateCertBody, "\n", "", -1)
	sessionPrivateKeyPem            = fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----\n", sessionPrivateKeyBody)
	sessionPrivateKeyBodyNoNewLine  = strings.Replace(sessionPrivateKeyBody, "\n", "", -1)
	sessionPublicKeyPem             = fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----\n", sessionPublicKeyBody)
	sessionPublicKeyBodyNoNewLine   = strings.Replace(sessionPublicKeyBody, "\n", "", -1)
)

const whateverRegion = REGION_PHX

func TestX509FederationClient_VeryFirstSecurityToken(t *testing.T) {
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		expectedKeyId := fmt.Sprintf("%s/fed-x509/%s", tenancyId, leafCertFingerPrint)
		assert.True(t, strings.HasPrefix(r.Header.Get("Authorization"), fmt.Sprintf(`Signature version="1",headers="date (request-target) content-length content-type x-content-sha256",keyId="%s",algorithm="rsa-sha256",signature=`, expectedKeyId)))

		expectedBody := fmt.Sprintf(`{"certificate":"%s","publicKey":"%s","intermediateCertificates":["%s"]}`,
			leafCertBodyNoNewLine, sessionPublicKeyBodyNoNewLine, intermediateCertBodyNoNewLine)
		var buf bytes.Buffer
		buf.ReadFrom(r.Body)
		assert.Equal(t, expectedBody, buf.String())

		// Return response
		fmt.Fprintf(w, "\n{\n  \"token\" : \"%s\"\n}\n", expectedSecurityToken)
	}))
	defer authServer.Close()

	mockSessionKeySupplier := new(mockSessionKeySupplier)
	mockSessionKeySupplier.On("Refresh").Return(nil).Once()
	mockSessionKeySupplier.On("PublicKeyPemRaw").Return([]byte(sessionPublicKeyPem))

	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	mockLeafCertificateRetriever.On("Refresh").Return(nil).Once()
	mockLeafCertificateRetriever.On("CertificatePemRaw").Return([]byte(leafCertPem))
	mockLeafCertificateRetriever.On("Certificate").Return(createCertificate(leafCertPem))
	mockLeafCertificateRetriever.On("PrivateKey").Return(parsePrivateKey(leafCertPrivateKeyPem))

	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)
	mockIntermediateCertificateRetriever.On("Refresh").Return(nil).Once()
	mockIntermediateCertificateRetriever.On("CertificatePemRaw").Return([]byte(intermediateCertPem))

	federationClient := &x509FederationClient{
		tenancyId:                         tenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	// Overwrite with the authServer's URL
	federationClient.authClient.Host = authServer.URL
	federationClient.authClient.BasePath = ""

	actualSecurityToken, err := federationClient.SecurityToken()

	assert.NoError(t, err)
	assert.Equal(t, expectedSecurityToken, actualSecurityToken)
	mockSessionKeySupplier.AssertExpectations(t)
	mockLeafCertificateRetriever.AssertExpectations(t)
	mockIntermediateCertificateRetriever.AssertExpectations(t)
}

func TestX509FederationClient_RenewSecurityToken(t *testing.T) {
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		expectedKeyId := fmt.Sprintf("%s/fed-x509/%s", tenancyId, leafCertFingerPrint)
		assert.True(t, strings.HasPrefix(r.Header.Get("Authorization"), fmt.Sprintf(`Signature version="1",headers="date (request-target) content-length content-type x-content-sha256",keyId="%s",algorithm="rsa-sha256",signature=`, expectedKeyId)))

		expectedBody := fmt.Sprintf(`{"certificate":"%s","publicKey":"%s","intermediateCertificates":["%s"]}`,
			leafCertBodyNoNewLine, sessionPublicKeyBodyNoNewLine, intermediateCertBodyNoNewLine)
		var buf bytes.Buffer
		buf.ReadFrom(r.Body)
		assert.Equal(t, expectedBody, buf.String())

		// Return response
		fmt.Fprintf(w, "\n{\n  \"token\" : \"%s\"\n}\n", expectedSecurityToken)
	}))
	defer authServer.Close()

	mockSessionKeySupplier := new(mockSessionKeySupplier)
	mockSessionKeySupplier.On("Refresh").Return(nil).Once()
	mockSessionKeySupplier.On("PublicKeyPemRaw").Return([]byte(sessionPublicKeyPem))

	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	mockLeafCertificateRetriever.On("Refresh").Return(nil).Once()
	mockLeafCertificateRetriever.On("CertificatePemRaw").Return([]byte(leafCertPem))
	mockLeafCertificateRetriever.On("Certificate").Return(createCertificate(leafCertPem))
	mockLeafCertificateRetriever.On("PrivateKey").Return(parsePrivateKey(leafCertPrivateKeyPem))

	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)
	mockIntermediateCertificateRetriever.On("Refresh").Return(nil).Once()
	mockIntermediateCertificateRetriever.On("CertificatePemRaw").Return([]byte(intermediateCertPem))

	mockSecurityToken := new(mockSecurityToken)
	mockSecurityToken.On("Valid").Return(false)

	federationClient := &x509FederationClient{
		tenancyId:                         tenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	// Overwrite with the authServer's URL
	federationClient.authClient.Host = authServer.URL
	federationClient.authClient.BasePath = ""
	federationClient.securityToken = mockSecurityToken

	actualSecurityToken, err := federationClient.SecurityToken()

	assert.NoError(t, err)
	assert.Equal(t, expectedSecurityToken, actualSecurityToken)
	mockSessionKeySupplier.AssertExpectations(t)
	mockLeafCertificateRetriever.AssertExpectations(t)
	mockIntermediateCertificateRetriever.AssertExpectations(t)
}

func TestX509FederationClient_GetCachedSecurityToken(t *testing.T) {
	mockSessionKeySupplier := new(mockSessionKeySupplier)
	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)

	mockSecurityToken := new(mockSecurityToken)
	mockSecurityToken.On("Valid").Return(true)
	mockSecurityToken.On("String").Return(expectedSecurityToken)

	federationClient := &x509FederationClient{
		tenancyId:                         tenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	federationClient.securityToken = mockSecurityToken

	actualSecurityToken, err := federationClient.SecurityToken()

	assert.NoError(t, err)
	assert.Equal(t, expectedSecurityToken, actualSecurityToken)

	mockSessionKeySupplier.AssertNotCalled(t, "Refresh")
	mockSessionKeySupplier.AssertNotCalled(t, "PublicKeyPemRaw")
	mockLeafCertificateRetriever.AssertNotCalled(t, "Refresh")
	mockLeafCertificateRetriever.AssertNotCalled(t, "CertificatePemRaw")
	mockLeafCertificateRetriever.AssertNotCalled(t, "Certificate")
	mockLeafCertificateRetriever.AssertNotCalled(t, "PrivateKey")
	mockIntermediateCertificateRetriever.AssertNotCalled(t, "Refresh")
	mockIntermediateCertificateRetriever.AssertNotCalled(t, "CertificatePemRaw")
}

func TestX509FederationClient_RenewSecurityTokenSessionKeySupplierError(t *testing.T) {
	mockSessionKeySupplier := new(mockSessionKeySupplier)
	expectedErrorMessage := "TestSessionKeySupplierRefreshError"
	mockSessionKeySupplier.On("Refresh").Return(fmt.Errorf(expectedErrorMessage)).Once()

	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)

	mockSecurityToken := new(mockSecurityToken)
	mockSecurityToken.On("Valid").Return(false)

	federationClient := &x509FederationClient{
		tenancyId:                         tenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	federationClient.securityToken = mockSecurityToken

	actualSecurityToken, actualError := federationClient.SecurityToken()

	assert.Empty(t, actualSecurityToken)
	assert.EqualError(t, actualError, expectedErrorMessage)
}

func TestX509FederationClient_RenewSecurityTokenLeafCertificateRetrieverError(t *testing.T) {
	mockSessionKeySupplier := new(mockSessionKeySupplier)
	mockSessionKeySupplier.On("Refresh").Return(nil).Once()

	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	expectedErrorMessage := "TestLeafCertificateRetrieverError"
	mockLeafCertificateRetriever.On("Refresh").Return(fmt.Errorf(expectedErrorMessage)).Once()

	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)

	mockSecurityToken := new(mockSecurityToken)
	mockSecurityToken.On("Valid").Return(false)

	federationClient := &x509FederationClient{
		tenancyId:                         tenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	federationClient.securityToken = mockSecurityToken

	actualSecurityToken, actualError := federationClient.SecurityToken()

	assert.Empty(t, actualSecurityToken)
	assert.EqualError(t, actualError, expectedErrorMessage)
}

func TestX509FederationClient_RenewSecurityTokenIntermediateCertificateRetrieverError(t *testing.T) {
	mockSessionKeySupplier := new(mockSessionKeySupplier)
	mockSessionKeySupplier.On("Refresh").Return(nil).Once()

	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	mockLeafCertificateRetriever.On("Refresh").Return(nil).Once()
	mockLeafCertificateRetriever.On("Certificate").Return(createCertificate(leafCertPem))

	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)
	expectedErrorMessage := "TestLeafCertificateRetrieverError"
	mockIntermediateCertificateRetriever.On("Refresh").Return(fmt.Errorf(expectedErrorMessage)).Once()

	mockSecurityToken := new(mockSecurityToken)
	mockSecurityToken.On("Valid").Return(false)

	federationClient := &x509FederationClient{
		tenancyId:                         tenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	federationClient.securityToken = mockSecurityToken

	actualSecurityToken, actualError := federationClient.SecurityToken()

	assert.Empty(t, actualSecurityToken)
	assert.EqualError(t, actualError, expectedErrorMessage)
}

const previousTenancyId = `ocidv1:tenancy:oc1:phx:1460406592661:aaaaaaaab4faofrfkxecohhjuivjq262pu`

func TestX509FederationClient_RenewSecurityTokenUnexpectedTenancyIdUpdateError(t *testing.T) {
	mockSessionKeySupplier := new(mockSessionKeySupplier)
	mockSessionKeySupplier.On("Refresh").Return(nil).Once()

	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	mockLeafCertificateRetriever.On("Refresh").Return(nil).Once()
	mockLeafCertificateRetriever.On("Certificate").Return(createCertificate(leafCertPem))

	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)

	mockSecurityToken := new(mockSecurityToken)
	mockSecurityToken.On("Valid").Return(false)

	federationClient := &x509FederationClient{
		tenancyId:                         previousTenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	federationClient.securityToken = mockSecurityToken

	actualSecurityToken, actualError := federationClient.SecurityToken()

	assert.Empty(t, actualSecurityToken)
	assert.EqualError(t, actualError, fmt.Sprintf("Unexpected update of tenancy OCID in the leaf certificate. Previous tenancy: %s, Updated: %s", previousTenancyId, tenancyId))
}

func TestX509FederationClient_AuthServerInternalError(t *testing.T) {
	authServer := httptest.NewServer(http.HandlerFunc(internalServerError))
	defer authServer.Close()

	mockSessionKeySupplier := new(mockSessionKeySupplier)
	mockSessionKeySupplier.On("Refresh").Return(nil).Once()
	mockSessionKeySupplier.On("PublicKeyPemRaw").Return([]byte(sessionPublicKeyPem))

	mockLeafCertificateRetriever := new(mockCertificateRetriever)
	mockLeafCertificateRetriever.On("Refresh").Return(nil).Once()
	mockLeafCertificateRetriever.On("CertificatePemRaw").Return([]byte(leafCertPem))
	mockLeafCertificateRetriever.On("Certificate").Return(createCertificate(leafCertPem))
	mockLeafCertificateRetriever.On("PrivateKey").Return(parsePrivateKey(leafCertPrivateKeyPem))

	mockIntermediateCertificateRetriever := new(mockCertificateRetriever)
	mockIntermediateCertificateRetriever.On("Refresh").Return(nil).Once()
	mockIntermediateCertificateRetriever.On("CertificatePemRaw").Return([]byte(intermediateCertPem))

	federationClient := &x509FederationClient{
		tenancyId:                         tenancyId,
		sessionKeySupplier:                mockSessionKeySupplier,
		leafCertificateRetriever:          mockLeafCertificateRetriever,
		intermediateCertificateRetrievers: []x509CertificateRetriever{mockIntermediateCertificateRetriever},
	}
	federationClient.authClient = newAuthClient(whateverRegion, federationClient)
	// Overwrite with the authServer's URL
	federationClient.authClient.Host = authServer.URL
	federationClient.authClient.BasePath = ""

	_, err := federationClient.SecurityToken()

	assert.Error(t, err)
}

func createCertificate(certPem string) *x509.Certificate {
	var block *pem.Block
	block, _ = pem.Decode([]byte(certPem))
	cert, _ := x509.ParseCertificate(block.Bytes)
	return cert
}

func parsePrivateKey(privateKeyPem string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privateKeyPem))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

type mockSessionKeySupplier struct {
	mock.Mock
}

func (m *mockSessionKeySupplier) Refresh() error {
	args := m.Called()
	return args.Error(0)
}
func (m *mockSessionKeySupplier) PrivateKey() *rsa.PrivateKey {
	args := m.Called()
	return args.Get(0).(*rsa.PrivateKey)
}
func (m *mockSessionKeySupplier) PublicKeyPemRaw() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

type mockCertificateRetriever struct {
	mock.Mock
}

func (m *mockCertificateRetriever) Refresh() error {
	args := m.Called()
	return args.Error(0)
}
func (m *mockCertificateRetriever) CertificatePemRaw() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}
func (m *mockCertificateRetriever) Certificate() *x509.Certificate {
	args := m.Called()
	return args.Get(0).(*x509.Certificate)
}
func (m *mockCertificateRetriever) PrivateKeyPemRaw() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}
func (m *mockCertificateRetriever) PrivateKey() *rsa.PrivateKey {
	args := m.Called()
	return args.Get(0).(*rsa.PrivateKey)
}

type mockSecurityToken struct {
	mock.Mock
}

func (m *mockSecurityToken) String() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockSecurityToken) Valid() bool {
	args := m.Called()
	return args.Bool(0)
}
