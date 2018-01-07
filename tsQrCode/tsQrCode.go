package tsQrCode

import (
	qrcode "github.com/skip2/go-qrcode"
)

//tsQrCode.CreateQrCodeFile("test.png", "http://www.baidu.com", 256)
func CreateQrCodeFile(path string, content string, size int)(err error) {
	err = qrcode.WriteFile(content, qrcode.Highest, size, path)
	return
}

func CreateQrCode(content string, size int)(data []byte, err error) {
	data, err = qrcode.Encode(content, qrcode.Highest, size)
	return
}