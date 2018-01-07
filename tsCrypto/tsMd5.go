package tsCrypto

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(data []byte)(string) {
	md5d := md5.New()
	md5d.Write(data)
	des_md5 := hex.EncodeToString(md5d.Sum(nil))
	return des_md5
}