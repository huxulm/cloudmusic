package util

import (
	"encoding/base64"
)

func ToBase64(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

func ToBase64String(src []byte) string {
	return string(ToBase64(src))
}
