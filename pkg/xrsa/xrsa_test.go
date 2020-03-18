package xrsa

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Demo(t *testing.T)  {
	message := "appId=test&data={mobile=13800138000&username=tom}&nonce=12345678&timestamp=1536566968643"
	priKey := `-----BEGIN PRIVATE KEY-----
MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBALM/wbHfBxLw7cvU
/hOWczFpw2E9gXuheLT+5rODL8JuetEb+8Byf2hMLmHd/ZC7ekKnZcQIXHO6qwev
KymUX3uoZo8l8hnbIgaTfJU4OveB+/n7LCrxWrP7NPvMJ3vCKtNfvVHX2Syp6OyO
g22Ea9T+IMJFjXkGHxhMnjxIBldPAgMBAAECgYBXRhKj7Qg/s6jEVscCeH1LsMgd
P1oc/MiNr5JaEuSdETehOrO/xr+742Ot6Oz9nBpx+5kz8jI5csi1LBei3ggCl4/R
f5aHyf03WE1MXUDQnzzIEaYW92BOvz1bwhB85B01x2sJGaxRcff5vGxciX1pEQ9M
IFpQgNDxAq0elBsoSQJBANzw+75oFGD/veXi3xV3urJIiEu8hS0aeu0sK+KTIwG/
qHfvaz6SQAEeQ5VnweGhWaxLNhbXnaP2kA7aG9y72XUCQQDPsSknsX+nAWdabDmL
igchWI0YhpRu/4zK5g6t9cCFnE7mGNBMMtF95fbMmMMzOeHOrqrjehaREnqW6iL+
HFEzAkEAispzShQ1oQ1mbFANVX4F8NLxk6oUetXknLKfytAlMIcGPHlRBFh5DnrF
d8hbGfLy0vHYQ5ck9wf/TOUklZAHtQJBALZJ5zZIiyLYj48EqCk6F3IEF6rsDAG0
WC2JaF708GoUvAcmxkPq4oYevdPrTIB8kB1onuKTOJVR47jrfZUM45MCQQCRl2gJ
P/Gpzdl0zsvonwle4WVqNfIIUacoIQ/3yLTut4Gx3jzqdU5WKVkiqs3dOhU75jcg
U0CMWXDVnJ/KGuno
-----END PRIVATE KEY-----`
	pubKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCzP8Gx3wcS8O3L1P4TlnMxacNh
PYF7oXi0/uazgy/CbnrRG/vAcn9oTC5h3f2Qu3pCp2XECFxzuqsHrysplF97qGaP
JfIZ2yIGk3yVODr3gfv5+ywq8Vqz+zT7zCd7wirTX71R19ksqejsjoNthGvU/iDC
RY15Bh8YTJ48SAZXTwIDAQAB
-----END PUBLIC KEY-----`

	handle,err := NewXRsa([]byte(priKey),[]byte(pubKey))
	assert.NoError(t,err)

	result,err := handle.Encrypt(message)
	assert.NoError(t,err)
	fmt.Println(result)

	ok,err := handle.Verify(message,result)
	assert.NoError(t,err)
	fmt.Println(ok)
}
