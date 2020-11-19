package util

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestGenKeys(t *testing.T) {
	if pri, pub, err := GenerateKey64(2048); err == nil {
		fmt.Printf("public key:\n%s\n", pub)
		fmt.Printf("private key:\n%s\n", pri)
	}
	fmt.Printf("default public key:\n%s\n", publicKey)
}

func TestRsaEncrypt(t *testing.T) {
	pub := []byte("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2vPFQ3RZ3l9/3JXx04BnRppZgaGmtaNfmcxDR7X7u3ZDm6nN9Yj0BigtXpFOZbGfR/cucTGBXOFQstb4HplYwIlI2oxSogCl8qS0Oj539jBbNO1xNyMf8fNGFLD/1S9kclaNlOKIanMuB8Lrs1v8sm2jJ+IFo0mOvx6SPGlDYj88c6Dk3B3fW0TGbmDw1jjtZ3Df2XjaGxtMBjpxTYghx2ZXlQkLNl+HTrBWFJ9QLhkCAp6pYRwNI3hMVdBPhsyPUZNsX4odzqDoKC1D8jmJuAnwv2lN+LSQexFZGZldo3Z7Je0nKXBCjRo2sL921mCHObObCiWdcM7KftViWKeRFQIDAQAB")
	// block, _ := pem.Decode(pub)
	// if block == nil {
	// 	panic("failed to parse PEM block containing the public key")
	// }
	rsaKey, err := PublicKeyFrom64(string(pub))
	if err != nil {
		t.Fatal(err)
	}
	if cipher, err := PublicEncrypt(rsaKey, []byte("hello,rsa!")); err == nil {
		fmt.Printf("%s\n", ToBase64String(cipher))
	} else {
		t.Fatal(err)
	}
}

func TestRsaDecrypt(t *testing.T) {
	pri := []byte("MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDa88VDdFneX3/clfHTgGdGmlmBoaa1o1+ZzENHtfu7dkObqc31iPQGKC1ekU5lsZ9H9y5xMYFc4VCy1vgemVjAiUjajFKiAKXypLQ6Pnf2MFs07XE3Ix/x80YUsP/VL2RyVo2U4ohqcy4HwuuzW/yybaMn4gWjSY6/HpI8aUNiPzxzoOTcHd9bRMZuYPDWOO1ncN/ZeNobG0wGOnFNiCHHZleVCQs2X4dOsFYUn1AuGQICnqlhHA0jeExV0E+GzI9Rk2xfih3OoOgoLUPyOYm4CfC/aU34tJB7EVkZmV2jdnsl7ScpcEKNGjawv3bWYIc5s5sKJZ1wzsp+1WJYp5EVAgMBAAECggEAOdPaTMeFGMtGPN9JYF/wZuBLNYaZPi3pCPi+6EizNL62en1tjSTftmSUHv7noso31Ez/chvuG7bSlnlaTNKZcziPClG3XVwoeB5yD1QCYP5CiIcRuMtJG6Dx6HMbDpYWxHABssnl15+WD6y2jceExSdZySXT8V1zGx0a6GOwzViC46e4v5cryJDhdtrR3w5PioywuljmoKZLZxSfN8s3pHpzYGGSGx8ZEGOjZGyIAtyZDO64/RQzRMCMrLdkcKJR/xfNFoEJbWLaV9Cfa7xGOIdargG5+GhcaEqqFlGlOB0SF0NhYLkyE7h8qLQETP2MFU60DAP8T8bSXHw67RjTHQKBgQD9+MdgRUL4uthbgeCx2JnKMRYs50FilBRzS7V2AbAIKp6MVKlFlXnLsM5sDwLEChc5YZiCWiY1ksx+9WiJ5XqVXq6UVlLTAvVm3s31ACS+zNikgvTk6RUfFEJWg0L1NBBYiDmWzPw5xDRZX7ZhA2UZCIkhgsYk6NiTYS2FJwZw0wKBgQDcs2XHvvtFSVhRT7zUddE1J5coCbk0gX8WphROmiaVhwmrtPbjkWjE/INk//K1Hi98cBoiiLICj3I7uIXtLZ8kJ0Dk5haA9w9EfWdssqL7+uQbEaG2ZVsDLUoOp/id3DRkuOw+lfLCB9m+8ysOzH3Y9pI03O68lVhBkV0nWJcFdwKBgQCAlptr9OWJxiCRhFrd8Qs8wkm2Boimfs8z6RtBWm1kVR/dcWgbISl/pk99isQdufY0SWPMbBR4f5Emnt6FIzlDs6K03FgnGpJhuWGSMLZqHJF8CgowdIsLJ1jHMMKQIAI+sQpnffYe0Wan8bwHto5TdGzqGKp2OaaeSE+h4TZznQKBgQDTwMzG+dbcg6LZYUgsEV7JQVMmdy3b0uO9N1wEjqK2lPoSZW78qWd8mUr4fRrB7FRjcKuitsUU780Kv+C/0CYA6ii5dong8ysS679v63W1juONlT0zY8wPIEUOCtvfmogqm8MPyY9B24ZwT5/gcxPMN8fQMKpfBmvHfVGDjtxXwQKBgQCB2916kCvIN7gbh5W5zkHOatKaVs05qx9c8H8xtPefn2r6ctIfWngY+2rtkpPplxjwtCD7op6iTgtfNQX3/hpunKef7nZuAsC+VxUMPlQTYnf70s5PsqdWhRZMgS1bBFsPzn30FoaNr9Pw7bOtWDK512tytBaB7ZuGBIrJY9fcKw==")
	rsaKey, err := PrivateKeyFrom64(string(pri))
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("QKXMaJoWVtA2eXdO2NKT2GjKGRHJ2/rm/Esq/pOZCIZ5vGkUOoRPlPYviTACZN4RrXxb0+QHoKNkoYEz/Wju5bGv+NopMDZVIGgwdhl6I4eg4EHekae2Ue7OJpJYl1pcCeRwCwYirwLVSoH7npisJLQYcvbQVC4UtKRSqq2lAw1Iodn7LQFd6qNK523IilbavZpkgxvj4JlyNZw02oS9Wyg6efKxyHJOOdeHjnzfxKE+cVqohNTMKVOCAXzjkeWQsTGZhIG/oD5A8hHdevSLGd6ozkcIjdoQw8S70TQY0v58XGyqhIZPLuOa+N6KPk3n0AujvjO5vgTJCcHYE8HKIw==")
	dst, _ := base64.StdEncoding.DecodeString(string(data))
	if plain, err := PrivateDecrypt(rsaKey, dst); err == nil {
		fmt.Printf("%s\n", string(plain))
	}
}
