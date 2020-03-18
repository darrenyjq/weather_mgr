package xrsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type XRsa struct {
	privateKey	*rsa.PrivateKey
	publicKey	*rsa.PublicKey
}


func NewXRsa(prvKey, pubKey []byte) (handle *XRsa,err error) {
	handle = new(XRsa)
	//格式化私钥
	block, _ := pem.Decode(prvKey)
	if block == nil {
		err = errors.New("private block is nil")
		return
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes);
	if  err != nil {
		return
	}
	var ok bool
	handle.privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		err = errors.New("privateKey interface not ok")
		return
	}

	//格式化公钥
	block, _ = pem.Decode(pubKey)
	if block == nil {
		err = errors.New("public block is nil")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if  err != nil {
		return
	}
	handle.publicKey,ok = pubInterface.(*rsa.PublicKey)
	if !ok {
		err = errors.New("publicKey interface not ok")
		return
	}

	return
}

func (x *XRsa) Encrypt(message string) (base64Result string,err error) {
	digest := x.sha256(message)
	b,err := rsa.SignPKCS1v15(rand.Reader, x.privateKey, crypto.SHA256, digest)
	if err != nil {
		return
	}
	base64Result = base64.StdEncoding.EncodeToString(b)
	return
}

func (x *XRsa) Verify(message string, SignBase64 string) (ok bool,err error) {
	sign,err := base64.StdEncoding.DecodeString(SignBase64)
	if err != nil {
		return
	}
	digest := x.sha256(message)
	err = rsa.VerifyPKCS1v15(x.publicKey,crypto.SHA256,digest,sign)
	if err != nil {
		return
	}
	ok = true
	return
}

//hash字符串
func (x *XRsa) sha256(message string) (digest []byte) {
	messageBytes := bytes.NewBufferString(message)
	hash := sha256.New()
	hash.Write(messageBytes.Bytes())
	digest = hash.Sum(nil)
	return
}