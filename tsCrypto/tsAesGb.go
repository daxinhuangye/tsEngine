package tsCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//用户表模型
type AesGB struct {
	Strkey string
}

func (this *AesGB) getKey() []byte {

	keyLen := len(this.Strkey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(this.Strkey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func (this *AesGB) Encrypt(origData []byte) ([]byte, error) {
	key := this.getKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	//origData = this.PKCS5Padding(origData, blockSize)
	origData = ZeroPadding(origData, blockSize)
	blockMode := cipher.NewECBEncrypter(block)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return []byte(base64.StdEncoding.EncodeToString(crypted)), nil
}

//解密字符串
func (this *AesGB) Decrypt(data string) ([]byte, error) {

	crypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	key := this.getKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewECBDecrypter(block)

	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//origData = this.PKCS5UnPadding(origData)
	origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	if padding == blockSize {
		return ciphertext
	}
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {

	index := bytes.IndexByte(origData, 0)
	if index == -1 {
		return origData
	}
	rbyf_pn := origData[0:index]
	return rbyf_pn

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
