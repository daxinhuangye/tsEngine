/****使用方法******
	oAes := new(tsCrypto.AesGA)
	oAes.Key = beego.AppConfig.String("AesKey")
	oAes.Iv = beego.AppConfig.String("AesKey")
	oAes.PadType = 1

	token, _ := oAes.EncryptCBC([]byte("asdfsadfdddddddddd"))
	token1 := tsCrypto.ToHexString(token)

	beego.Trace(token)
	aa := tsCrypto.FromHexString(token1)

**********************/
package tsCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

const (
	PadString int = 1
	PadByte   int = 2
)

type AesGA struct {
	Key     string // 密码
	Iv      string // 密码
	PadType int    // 数据补齐方式
}

func (this *AesGA) EncryptCBC(data []byte) ([]byte, error) {
	key := this.getKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 补齐数据
	block_size := block.BlockSize()
	pad_data := this.padding(data, block_size)

	block_mode := cipher.NewCBCEncrypter(block, this.getIv())

	en_data := make([]byte, len(pad_data))
	block_mode.CryptBlocks(en_data, pad_data)
	return en_data, nil
}

func (this *AesGA) DecryptCBC(data []byte) ([]byte, error) {
	key := this.getKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	block_mode := cipher.NewCBCDecrypter(block, this.getIv())
	de_data := make([]byte, len(data))
	block_mode.CryptBlocks(de_data, data)

	return this.unpadding(de_data), nil
}

func (this *AesGA) getKey() []byte {
	keyLen := len(this.Key)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(this.Key)
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

func (this *AesGA) getIv() []byte {
	ivLen := len(this.Iv)
	if ivLen < 16 {
		panic("res iv 长度不能小于16")
	}
	arriv := []byte(this.Iv)
	if ivLen >= 32 {
		//取前32个字节
		return arriv[:32]
	}
	if ivLen >= 24 {
		//取前24个字节
		return arriv[:24]
	}
	//取前16个字节
	return arriv[:16]
}

// 补齐数据
func (this *AesGA) padding(data []byte, block_size int) []byte {
	// 必须添加，最后一位表示补的大小
	if this.PadType == PadByte {
		pad_size := block_size - len(data)%block_size
		pad_data := make([]byte, pad_size)
		pad_data[pad_size-1] = byte(pad_size)
		return append(data, pad_data...)
	}

	// 完整数据无需补齐，如果需要补齐就0填充数据
	if this.PadType == PadString {
		if len(data)%block_size == 0 {
			return data
		}
		pad_size := block_size - len(data)%block_size
		pad_data := make([]byte, pad_size)
		return append(data, pad_data...)
	}
	return nil
}

func (this *AesGA) unpadding(data []byte) []byte {
	if this.PadType == PadByte {
		pad_size := int(data[len(data)-1])
		return data[:len(data)-pad_size]
	}

	if this.PadType == PadString {
		index := bytes.IndexByte(data, 0)
		if index == -1 {
			return data
		}
		return data[:index]
	}
	return nil
}
