package tsVideo

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"tsEngine/tsTime"
)

const (
	Letv_time = 3600
)

type dispatch struct {
	Mp4_1300 []string `json:"1300"`
	Mp4_1000 []string `json:"1000"`
	Mp4_350  []string `json:"350"`
	Mp4      []string `json:"mp4"`
}

type playurl struct {
	Dispatch dispatch `json:"dispatch"`
	Domain   []string `json:"domain"`
}

type Letv struct {
	Videocode string
	Videoid   int64
	Cacheid   string
	Playurl   playurl `json:"playurl"`
}

func (this *Letv) GetVideo() (string, error) {

	key := this.GetKey()
	api := "http://api.le.com/mms/out/video/playJsonH5?platid=3&splatid=301&tss=ios&id=" + this.Videocode + "&detect=0&dvtype=1000&accessyx=1&domain=m.le.com&tkey=" + fmt.Sprintf("%d", key) + "&devid=0512382421944a2aecd5b44b59743631"

	//建立一个请求对象
	curl := httplib.Get(api)
	fmt.Println(api)
	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return "", err
	}

	json.Unmarshal(content, this)

	result := this.getData()

	return result, nil

}

func (this *Letv) getData() string {

	var vJson VideoJson

	url := this.Playurl.Domain[0]

	//标清
	if len(this.Playurl.Dispatch.Mp4_350) > 0 {
		vJson.Fluent.M3u8 = url + this.Playurl.Dispatch.Mp4_350[0]

	}

	//高清
	if len(this.Playurl.Dispatch.Mp4_1000) > 0 {
		vJson.Height.M3u8 = url + this.Playurl.Dispatch.Mp4_1000[0]

	}

	//超清
	if len(this.Playurl.Dispatch.Mp4_1300) > 0 {
		vJson.Super.M3u8 = url + this.Playurl.Dispatch.Mp4_1300[0]

	}

	//转json字符串
	temp, err := json.Marshal(vJson)
	if err != nil {
		return ""
	}
	vstr := string(temp)

	return vstr

}

func (this *Letv) GetKey() int32 {
	var e, t int32

	t = int32(tsTime.CurrSe())

	for s := 0; s < 8; s++ {
		e = 1 & t
		t >>= 1
		e <<= 31
		t += e
	}

	return t ^ 185025305

}
