package tsCrypto

import (
	"encoding/hex"
	"strings"
)

func ToHexString(data []byte)(string) {
	res := hex.EncodeToString(data)
	
	return strings.ToUpper(res)
}

func FromHexString(data string)([]byte) {
	temp := strings.ToLower(data)
	res, err := hex.DecodeString(temp)
	if err!=nil {
		return nil
	}
	return res
}