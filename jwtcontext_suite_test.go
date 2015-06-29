package jwtcontext

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

var clientID = "testclient"
var emails = []string{"test@test.com"}

func TestJwtcontext(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jwtcontext Suite")
}

var ctx context.Context
var _ = BeforeSuite(func() {
	ctx = context.Background()
	Expect(ctx).NotTo(BeNil())
})

type jwtTestServer struct {
	user []byte
}

type user struct {
	Email string
}

func jwtTestContext(email string) func(context.Context, *http.Request) context.Context {
	data, _ := json.Marshal(user{
		Email: email,
	})

	return Add(jwtTestServer{
		user: data,
	})
}

func (server jwtTestServer) GetClient(ctx context.Context, clientID string) ([]byte, error) {
	return testPublicKey, nil
}

func (server jwtTestServer) GetUser(ctx context.Context, userID string, user interface{}) error {
	return json.Unmarshal(server.user, user)
}

var testPublicKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwIh/3G6FS9gaPfCWoZpB
5ZV/lEzWrGTHcRrhk/OaQTvsWQiuQGUYJCbmiXNV8j8M6arWtw7lg9GxwKiuHydk
DsW6+U16h/jmktom/rbKQvbpdYBssGaRoU9v9nbWGCBeHy3og8h5lB2G3g1MtOLr
LAEWxXvwULbcsCj583NVhukHrBC+YxnyW0Q9KI/YB3EhIur/mV4aH+IYPu3XRZjE
A1CJKNBlpZMc0EsLoKXb9cQZWr3wtJxwSbPCaoxGeayb/TvOdgZyBMJHlgDH8qfu
q8dkTDSjyQX8IGoOGXq1RGamdzcRpZNI+ly+R7uXDU7B+G+R6pG4Bc7lCcwiFxu7
RQIDAQAB
-----END RSA PUBLIC KEY-----`)

var testPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEAwIh/3G6FS9gaPfCWoZpB5ZV/lEzWrGTHcRrhk/OaQTvsWQiu
QGUYJCbmiXNV8j8M6arWtw7lg9GxwKiuHydkDsW6+U16h/jmktom/rbKQvbpdYBs
sGaRoU9v9nbWGCBeHy3og8h5lB2G3g1MtOLrLAEWxXvwULbcsCj583NVhukHrBC+
YxnyW0Q9KI/YB3EhIur/mV4aH+IYPu3XRZjEA1CJKNBlpZMc0EsLoKXb9cQZWr3w
tJxwSbPCaoxGeayb/TvOdgZyBMJHlgDH8qfuq8dkTDSjyQX8IGoOGXq1RGamdzcR
pZNI+ly+R7uXDU7B+G+R6pG4Bc7lCcwiFxu7RQIDAQABAoIBAQCvQNDaTsQI6Lni
XchPa4HGQJHz5QmXvNiKjQR+z0Q/UxNta+hpd1xFNB+vSAwYyOmxHS/7S5UBq3PY
wfJVK175z2TFn5AEiz0euXK7cpmj4OtbaL4GyCaRpiRtPXjeV0A1s20t3+NDPqlQ
IN0M3m1hDdCzjICQiBNUSshSoU5d0NU3Gf8JLREQ+6vpYd5PX/3YGmc6wxboRIgq
boigIF+uiz+RPX/IW8oG7KDNYba4n11sDg1gmog9SlpojIotvHNSLNvRBqCDd4Mg
jOMTO2haPTyquAjjh5WghVjQ3zBS6A8Kvu2iomiXWfyUUJ9wk+2Y/BmT4Lb5GbW7
mRrgUdk1AoGBAO9jvKf0mpL6Ouf6hK+PIUySADkhhS0CdOJcBQyhHMwPVoH/VtOW
anf3y8zq2XdPE1rJHuI2HKxDdOpqLppolj6sSIsMugeIa0Pz/jZeJ1jrhAn0Tg2j
0S+e580W4zbVL1sU5qBF3fjXKh/rFFtEUNOiN9WPVwAazzQDhLyA8GODAoGBAM3k
dIVf/PL2XiFOToz2rGOsYFbll8AxD9igiDtaDiMEWkypmM6fn9R7FsGZrwFwKyVq
6U/olfLGaMNLztPQ9qNs66IjjtkQiXggFu/ectiLAUinloMCAJkp8JZSMR1P+BZn
arD9l0BP4qHwimWSxx8JCgsep+QL76sYy/aB94OXAoGBAO0BHaJbNT4Wp4NML1lw
p1MbqUTmvucU64u/9u/OAqi09ry1g8zvunETz54NVUgd7deQDxd+41xZXMNIZONS
cI2UHC4uZ29VzHz2b9R72xDZ+1uvmCdA7LXs6SnPYxzCa6QNSVecVGss+vOm1fyV
4j/k9spko4njAQlKAnxuW5DpAoGBAMplUcXq3hVY6p+DAuS/eCdjRDEn8U0U7abd
6A2wMUVq+flynqw/bjbJ6UzXbpnTUhauSjwrP3wwXrvcwaynNkzgvaHdoobvIrPb
l7Ck0/DadKCfXe3bAguOltquYWmXOyGK9+6U2yJ1ZSwb8XB4IRX5JM5ZjGE4+pkO
/zGMAWxxAoGBANKYOsl9b0UR8kS8cfV+3cR1QNuYQeXM9lnWt5HjgBIwRb3fkF55
rS6ErCsJG/E5vpUNA6YD23Ubky1PDP7fIbxtL+7h1+JqQIvsS2cwRmAK/CYvgMti
KDKQtCQRuQNltjA+ACNcoyh2tcswagPvcZXyDzpN5pGYFCN6WgSwY2z/
-----END RSA PRIVATE KEY-----
`)

func authorizeRequest(req *http.Request, email string) *http.Request {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["cid"] = clientID
	token.Claims["exp"] = time.Now().Add(60 * time.Minute)
	token.Claims["sub"] = email
	accessToken, err := token.SignedString(testPrivateKey)
	if err != nil {
		panic(err)
	}
	req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bearer %s", accessToken))
	return req
}
