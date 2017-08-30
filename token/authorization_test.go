package token

import (
	"crypto/rsa"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDXmeFcJ5dPrTJBHw7n+q+aRV5tAHHkoxUIv+xF8wzvCroFPhgo
RYx4Xwy8eydOZZq4IyTeJ3tcYLYLJp6y/tDwMrP5CTt3tRyrUYWbYI1/IpWXYxKB
ExEan3CGLdKsRfoHl9XjL0EJjgbfFLEUavLhDOEpMrxStB34/sxlUyHtxQIDAQAB
AoGAKZy4o1m82ZL7FRiSfvGifOsQm4cvdEqJn8OzLZkIkQaH+sUOUo+B2iW2Rpaa
coYnw4x87CfApoa//Az7Kl8GIksxaLaakyXIKlu5baBvElL4FQ23beSOu5nGOjpF
2IjWTUmq0rN4fl1lTd2I85OZCJ/JYHIWv+MnrJmrCD2ehaECQQD7Z1R85YmJ+z3m
0ipPyWVwsFtMPinXwpweqHlhY97bZlvNVGhIE3WvRRk0xlAgx9C18SWEXtq+yj0/
F47vSnF5AkEA24r64u50Ynor5/uv2VFkWtcguv8XyI96Q0yx2xBfTb1ry/XyWKHu
/Ly6bDxtiQmlKKbJLoNQxyveNWHHirR3rQJANp0q6HsA1v3CY8tSL8UDiWh2XOjA
NEPtoGde/v4wIp9o4AEvKkE39pxfLmzKaWSe/XdmfAwTHxFWHseiYG9emQJAQSYk
/WK1tuN6VAUkUeENXseoNtWsQSASsBX6UX+ySGeFraj2mca0HNwElkn4V0o9cKrN
2LevKOujrUOu60JSoQJAJ6mjkZvXacflQ8KsysgMM6mstkSakwfFiUXkUNmpztnY
n06nXuam4qcYBpaAMUFf4miIwEa6NvWopHvtYyKP3A==
-----END RSA PRIVATE KEY-----`)

var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDXmeFcJ5dPrTJBHw7n+q+aRV5t
AHHkoxUIv+xF8wzvCroFPhgoRYx4Xwy8eydOZZq4IyTeJ3tcYLYLJp6y/tDwMrP5
CTt3tRyrUYWbYI1/IpWXYxKBExEan3CGLdKsRfoHl9XjL0EJjgbfFLEUavLhDOEp
MrxStB34/sxlUyHtxQIDAQAB
-----END PUBLIC KEY-----`)

func loadKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, nil, err
	}

	public, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, nil, err
	}

	return private, public, nil
}

func TestServerGenerateAuthorizationToken(t *testing.T) {
	private, public, err := loadKeys()
	if err != nil {
		t.Error(err)
	}

	authorization, err := NewAuthorization(private, public)
	if err != nil {
		t.Error(err)
	}

	authorization.AppId = 1
	authorization.ReleaseId = 2
	authorization.Admin = true

	token, err := authorization.GetSignedToken()
	if err != nil {
		t.Error(err)
	}

	if token != "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXBwSWQiOjEsInJlbGVhc2VJZCI6Mn0.BBTzsglHh70K5K6F3ixrAofhF-kiNvIQGN73W4WuEt8ISrMsGqoBRYywcyZ41G1tNkpahkevG8OG7eIgd373gjq-cTY65eVZNc5iSh0lGo0CL5IS_RhvhsSRET5lGPYzpTuSXQjmE4MgZkVBcK5IjkrLRdvD0usziLwXRgCWtAA" {
		t.Error("token not matched")
	}
}

func TestServerValidateAuthorizationToken(t *testing.T) {
	private, public, err := loadKeys()
	if err != nil {
		t.Error(err)
	}

	authorization, err := NewAuthorization(private, public)
	if err != nil {
		t.Error(err)
	}

	authorization.AppId = 1
	authorization.ReleaseId = 2
	authorization.Admin = true

	token, err := authorization.GetSignedToken()
	if err != nil {
		t.Error(err)
	}

	authorization2, err := NewAuthorization(private, public, token)
	if err != nil {
		t.Error(err)
	}

	if authorization2.AppId != authorization.AppId &&
		authorization2.ReleaseId != authorization.ReleaseId &&
		authorization2.Admin != authorization.Admin {
		t.Errorf("not matched")
	}
}
