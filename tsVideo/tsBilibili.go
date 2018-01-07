package tsVideo

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
)

type Bili struct {
	Videocode string
	Videoid   int64
	Cacheid   string
	Src       string `json:"src"`
}

func (this *Bili) GetVideo() (string, error) {

	url := "http://www.bilibili.com/m/html5?page=1&aid=" + this.Videocode

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return "", err
	}

	json.Unmarshal(content, this)

	return this.Src, nil
}
