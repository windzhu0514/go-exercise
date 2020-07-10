package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"
)

var pubPEM = `-----BEGIN PUBLIC KEY-----
MIIDOzCCAiOgAwIBAgIBMDANBgkqhkiG9w0BAQUFADA4MQswCQYDVQQGEwJDTjENMAsGA1UECwwEQ05DQjEaMBgGA1UEAwwRZS5iYW5rLmVjaXRpYy5jb20wHhcNMTgwMjExMDg0NTIyWhcNMzgwMjA2MDg0NTIyWjA4MQswCQYDVQQGEwJDTjENMAsGA1UECwwEQ05DQjEaMBgGA1UEAwwRZS5iYW5rLmVjaXRpYy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCkF+2AicVKj7SaHw3dbJt3i6fkL1WfLw1WRqe8r8Cc7qJOshaqNvCzW1qRX6E5H/umtl1Uj99V07uewUFk96xY/+s/GuBnbGoSrcu3OAHDgEGuY5atZo+umIk7LufAif2VUcNGY3nWxGcig20ExO/6nAf/G3Xxo4QL8fBdPG/prOXxSvtJiPls1Qg9zzSgAH+HMCAINMsuJmzDQiTt6Me8k7YHts+jWQF7KF25plITcW1Qmy3Aw8qYjVhbHn8KTAEeuQhmM5RS6KP1Hu71q4DYOWcx44QThSbiAYwG1JQBBwM8XnBfVYMpr6Qi0owibNYoZ/S6xwfRFGB0W1HeG9WfAgMBAAGjUDBOMB0GA1UdDgQWBBT0iLEXY9HIKNy5DG4d72l+R7Nf1zAfBgNVHSMEGDAWgBT0iLEXY9HIKNy5DG4d72l+R7Nf1zAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQB5MWz1RGFG537rJCtHp+LqxR9iJSFsHiW3ZoLIAeyD0oJ69RcL2gE/TNWmE9zYUkd9TdNtXqxlNPpj1P1/+x781neWnGou/n/XFS82T5S339X3DIjHc/IqOzwnxEOKH2V0NmK9iKgx6H05Q9MMvUXFsL3QK2hDMAVY28roRiC4S1yfJJaA08DfvXZf6cVx1xfWl+ks57+3knkoWap1rjwh1RdGk5ChPbzD0AnAcWTMWRCbjuJnttlmWZnI1I6mhcQUKUEMoj8sR8m11YJ5woscYPsIle/rJOOosuMghczD1vRcg3eLUaWn1A5rsBa82RyxhiuYocEQVX59Hy6v3npT
-----END PUBLIC KEY-----`

var privPEM = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-CBC,95E59EDB85FA21EA

uxKuKBuGSrNmemW/kYQm/q431gxfN7B1qP6/7hKBd5uN+zFO7Xl5Fv6XdxGOCEFR
zQnEIwkO6LOORFsvht6sjBORetj/UGNsERNwy/EhsgwjuzxlKk+Egd7cpFQ1tBE0
n4c9Am4t5Iy27zjLd2+Y2zNxPrDXY18WR6ws2dA6jaC3IPCH2HZxqx5fj0xgD24e
3nlimTFwVLH6Vdm2D61bfwIN+rlwTwoFnelsDqlGjjfxo9vkckkgS+Dk1brDqxGE
4dzT4Y1N/rd6pPFZ592EVnFAkxicyIEadqNaTtTkCR78gC2KjCgaGr9zo+peyJzb
Hv0h8cNlZEo1Py6HxnTyufE69N+x2chHJa4Lg/R4oQR3P0aEJtL2/DVMosbXPqAx
kdk770nEbB4po9/OC8VAn6v+Q7l6it5F5/grIavxTxsqsF1aEp+EEDsv60+k3VbE
5StWRO3rPppxsESJkYzYoeoFnk2HVDT8zOELm5P5gPjjMd62Z9fwpP0/8BJc07kp
pSKfNh7W0l3fow12Fb3NiK0MQzTeCBKak31a1yBDGWmUXWgwPAzgWSGuOaMsqVTe
Jgsr/JqE/6KTwIrh9Q8OlbKhhkuHpipHGYirNiFC365Pf9kRXYhSiP0wi3Rc4u2z
m43vf+CdZ0UKyhX92vKPoy+SWgumzzL0/Lg1ndl8srZdWKOVKMLOXavxQumDdJ16
g1vo26Mm93/PQBs3DAwFtaB8pGF/1I1OB2DL/58BqKsPwyIRvo5ZKdNNqeJpSfwB
+GcpRvx/Uwjj+EdWiaUyHKSACnhdqGht8lcsgWSMkvgXcFW/0pyqEYIkFvo2J9jj
X1GsX1arWvsrHZNMjv+d0K6wn+0NSk/miHHxRDzo1u1b4hZOMJybNszTptmbDQd8
3e2hoQXoxy7dSRXMuxxXK5Bz0n1IH8AxDKhqsxa5e+vjYFEm1Fbl12kQP1InGSnr
cHsNGcctEzlt9GC0Aievu/rNBihPOdd8IRrpmmpL61R/xa6DLHOGczvdslNNiov4
xXuLb1VGGYAFSxvV/gHU2Ghv+Jwf4O5l10cUUovpFbAivnvrZf+9rajNAminHktg
gZBew6EHJmu/eToJAVykmYDF9ykgC4UXk/HaMrqGR2Jzrlc1N3q/CkQW7761T6BH
VjW7IA2xIZYzRESR7mcKwtMWuYXNkqnyMKHN2wQOprxexhyw08s/BZeP3iHic8+a
11oMwKl7EHmVpQrIzm8b1V1L9NXjHZIo/ZCJaq8Nv3HlGaJR0tZB5p68Le8XirUb
2yE9xUssmdfx3z67He9Ibrmi7b5bgnz8OkOU2t9O781NOzbKRkh7l5Mjmo8F9uLQ
NUxgcipCMoNArhq+Zvq3bHCsWe5sSTwIADJeqCkf0l20qgZbub5LceXUSxYhYkb8
Kk9pgh+8MmzrSoa1K6fAUd/tWWy5L0Y9GfNrQW3NYtVgfnnf4U43JI/z1b44Fsju
0JJh9Cd5LVkQNaKRVUZ9GhDWkl1TA9+oLXf3n1MZEbcO06bT7y0NQVrfXF8P8TQm
yMVEoQZKvzi/ScMfCobwvQM2srFzBLrkRMJP32lyayqnB5rohrobCGDPkmTGTfnm
-----END RSA PRIVATE KEY-----`

func TestRSAEncrypt(t *testing.T) {
	cipherText, err := RSAEncrypt([]byte("hhliotwapntjcqlt"), []byte(pubPEM))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(cipherText))
}

var certPEM = []byte(`-----BEGIN CERTIFICATE-----
MIIDOzCCAiOgAwIBAgIBMDANBgkqhkiG9w0BAQUFADA4MQswCQYDVQQGEwJDTjENMAsGA1UECwwEQ05DQjEaMBgGA1UEAwwRZS5iYW5rLmVjaXRpYy5jb20wHhcNMTgwMjExMDg0NTIyWhcNMzgwMjA2MDg0NTIyWjA4MQswCQYDVQQGEwJDTjENMAsGA1UECwwEQ05DQjEaMBgGA1UEAwwRZS5iYW5rLmVjaXRpYy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCkF+2AicVKj7SaHw3dbJt3i6fkL1WfLw1WRqe8r8Cc7qJOshaqNvCzW1qRX6E5H/umtl1Uj99V07uewUFk96xY/+s/GuBnbGoSrcu3OAHDgEGuY5atZo+umIk7LufAif2VUcNGY3nWxGcig20ExO/6nAf/G3Xxo4QL8fBdPG/prOXxSvtJiPls1Qg9zzSgAH+HMCAINMsuJmzDQiTt6Me8k7YHts+jWQF7KF25plITcW1Qmy3Aw8qYjVhbHn8KTAEeuQhmM5RS6KP1Hu71q4DYOWcx44QThSbiAYwG1JQBBwM8XnBfVYMpr6Qi0owibNYoZ/S6xwfRFGB0W1HeG9WfAgMBAAGjUDBOMB0GA1UdDgQWBBT0iLEXY9HIKNy5DG4d72l+R7Nf1zAfBgNVHSMEGDAWgBT0iLEXY9HIKNy5DG4d72l+R7Nf1zAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQB5MWz1RGFG537rJCtHp+LqxR9iJSFsHiW3ZoLIAeyD0oJ69RcL2gE/TNWmE9zYUkd9TdNtXqxlNPpj1P1/+x781neWnGou/n/XFS82T5S339X3DIjHc/IqOzwnxEOKH2V0NmK9iKgx6H05Q9MMvUXFsL3QK2hDMAVY28roRiC4S1yfJJaA08DfvXZf6cVx1xfWl+ks57+3knkoWap1rjwh1RdGk5ChPbzD0AnAcWTMWRCbjuJnttlmWZnI1I6mhcQUKUEMoj8sR8m11YJ5woscYPsIle/rJOOosuMghczD1vRcg3eLUaWn1A5rsBa82RyxhiuYocEQVX59Hy6v3npT
-----END CERTIFICATE-----`)

func TestRSAEncrypt2(t *testing.T) {
	ret, err := RSAEncrypt2([]byte("hhliotwapntjcqlt"), certPEM)

	t.Log(base64.StdEncoding.EncodeToString(ret), err)
}

func RSAEncrypt2(plainText, certPEM []byte) ([]byte, error) {
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, errors.New("public key error")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse certificate: " + err.Error())
	}

	pub := cert.PublicKey.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

func TestRSADecrypt(t *testing.T) {
	cipherText, err := base64.StdEncoding.DecodeString("18MM0nuwJizgzlQr+y/sb7E9bVx1ipwDYcro0RsL+ysSllBF70r1vOVEedwG4leTlcgvom23G+8lyXvrCG84dIdfrjZaW3jms3Xop0DX4Tz6o8cIUpLRYDemPkmMlR+Rp53URhJbe4pnwQAlsQ+lwvN5JjB4ZQ9NApEWDugJg3BBgF/O8MaLRrpzVwvcMWp5LF9/gvC7nM22RIP+ot2d7IKRNSdzc/9mSXt4yTG7LcSxq4oc0es7GAsDGZ2OEVE7IOnRubl/JxLZWnfTqUtgWTAH4PCqxGJhO8Wb4wY1YI7dDBC1t0SMioQDFCg24hXN1jHLv13eQc8sBSiXPybAtw==")
	if err != nil {
		t.Error(err)
	}

	plainText, err := RSADecrypt(cipherText, []byte(privPEM))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(plainText))
}

func TestRSAPrivateEncrypt(t *testing.T) {
	cipherText, err := RSAPrivateEncrypt([]byte("123"), []byte(privPEM))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(cipherText))
}

func TestRSAPublicDecrypt(t *testing.T) {
	cipherText, err := base64.StdEncoding.DecodeString("Z1ClBI17FouEJJRmXakVYlAWkQT5d4/19s5aMCUk9tHkjo67ZnlyCM0I0EJZEQgW17oXHpKTpoSBZI83L9slJnoZPHPTunftfvKUQmZXYfwvWC/cm/ZOe9s2dc1ylxeIGCr/cEFRSf6QdSdwTd3GdCb0q79lG8z7l7b5Elk3Fp6c60IkwmgfUlvG388I7UaZJNrb9RAzmHCU37+VGCIqOmN8Z9ROvcu2F4AKTNQvkf5PSXH459KlAW/lvyGTEBbYH6mL63/nvQM7Qk25AdhRqa317EyYPn5/fqGSlKAyuPFZAYFi2jtaJPnSDqn01/8OnB3hC43joZVQLxmNmuxlEg==")
	if err != nil {
		t.Error(err)
	}

	plainText, err := RSAPublicDecrypt(cipherText, []byte(pubPEM))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(plainText))
}
