package util

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	src := []byte("exampleplaintexts")
	if dst, err := AesEncrypt(src, presetKey, iv, ""); err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("%x\n", dst)
	}
}
func TestAesDecrypt(t *testing.T) {
	src, _ := hex.DecodeString("b5b468f1465a1ae5b2e8b03ed783d46be9ea0681e21ae86869d1e059848f2fc4")
	if dst, err := AesDecrypt(src, presetKey, iv, ""); err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("%s\n", dst)
	}
}

func TestWeapi(t *testing.T) {
	obj, _ := Weapi(map[string]string{"username": "<your email>"})
	if d, err := json.Marshal(obj); err == nil {
		fmt.Printf("%s", d)
	} else {
		t.Fatal(err)
	}
}
