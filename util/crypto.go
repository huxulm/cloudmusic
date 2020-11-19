package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var (
	iv          = []byte("0102030405060708")
	presetKey   = []byte("0CoJUm6Qyw8W8jud")
	linuxapiKey = []byte("rFgB&h#%2?^eDg:Q")
	base62      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	publicKey   = []byte("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----")
	publicKey64 = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB"
	eapiKey     = []byte("e82ckenh8dichen8")
)

func AesEncrypt(src, key, iv []byte, mode string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	src = PKCS7Padding(src)
	dst := make([]byte, len(src))
	var m cipher.BlockMode
	if mode == "" {
		m = cipher.NewCBCEncrypter(block, iv)
	} else if mode == "ecb" {
		m = NewECBEncrypter(block)
	}
	m.CryptBlocks(dst, src)
	return dst, nil
}

func AesDecrypt(src, key, iv []byte, mode string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// src, _ := hex.DecodeString("b5b468f1465a1ae5b2e8b03ed783d46b")
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	dst := make([]byte, len(src))
	var m cipher.BlockMode
	if mode == "" {
		m = cipher.NewCBCDecrypter(block, iv)
	} else if mode == "ecb" {
		m = NewECBDecrypter(block)
	}
	m.CryptBlocks(dst, src)
	return dst, nil
}

/* func encrypt_RSA(pub *rsa.PublicKey, data []byte) []byte {
	encrypted := new(big.Int)
	e := big.NewInt(int64(pub.E))
	payload := new(big.Int).SetBytes(data)
	encrypted.Exp(payload, e, pub.N)
	return encrypted.Bytes()
} */
func RsaEncrypt(src, key []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key pem parse error!")
	}
	rsaKey, _ := PublicKeyFrom(block.Bytes)
	if data, err := PublicEncryptNoPadding(rsaKey, src); err == nil {
		return data, nil
	} else {
		return nil, err
	}
}

func Weapi(obj interface{}) (url.Values, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var secretKey = make([]byte, 16)
	if _, err := rand.Read(secretKey); err != nil {
		return nil, err
	}
	for i, k := range secretKey {
		secretKey[i] = []byte(base62)[k%62]
	}
	// fmt.Printf("seckey: %s\n", secretKey)
	var encData []byte
	if encData, err = AesEncrypt(b, presetKey, iv, ""); err != nil {
		return nil, err
	}
	if encData, err = AesEncrypt(ToBase64(encData), secretKey, iv, ""); err != nil {
		return nil, err
	}
	var encKey []byte
	// reverse key
	for i, j := 0, len(secretKey)-1; i < j; i, j = i+1, j-1 {
		secretKey[i], secretKey[j] = secretKey[j], secretKey[i]
	}
	if encKey, err = RsaEncrypt(secretKey, []byte(publicKey)); err != nil {
		return nil, err
	}
	return url.Values{
		"params":    {ToBase64String(encData)},
		"encSecKey": {hex.EncodeToString(encKey)},
	}, nil
}

func LinuxApi(obj interface{}) (url.Values, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	enc, err := AesEncrypt(b, linuxapiKey, nil, "ecb")
	if err != nil {
		return nil, err
	}
	return url.Values{"eparams": {strings.ToUpper(hex.EncodeToString(enc))}}, nil
}

func Eapi(URL string, obj interface{}) (url.Values, error) {
	text, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	message := fmt.Sprintf(`nobody%suse%smd5forencrypt`, URL, text)
	digest := Md5String(message)
	data := fmt.Sprintf(`%s-36cd479b6b5-%s-36cd479b6b5-%s`, URL, text, digest)
	if enc, err := AesEncrypt([]byte(data), eapiKey, nil, "ecb"); err != nil {
		return nil, err
	} else {
		return url.Values{"params": {strings.ToUpper(hex.EncodeToString(enc))}}, nil
	}
}
