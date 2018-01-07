package tsCrypto

import (
	"encoding/base64"
)

func Base64Encode(str string) string {
	data := base64.StdEncoding.EncodeToString([]byte(str))
	return data
}

func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}
