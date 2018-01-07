package tsVideo

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

const (
	Tudou_jsonKey = "094b2a34e812a4282f25c7ca1987789f"
	Tudou_apiKey  = "6b72db72a6639e1d5a2488ed485192f6"
	Tudou_btanKey = "b45197d21f17bb8a"
	Tudou_time    = 0
)

type TudouMovieSegs struct {
	Url     string `json:"baseUrl"`
	Seconds int64  `json:"seconds"`
	Id      int64  `json:"k"`
	Size    int64  `json:"size"`
}

type Tudou struct {
	Videocode string
	Videoid   int64
	Cacheid   string
	F4v_99    []TudouMovieSegs `json:"99"`
	F4v_5     []TudouMovieSegs `json:"5"`
	F4v_3     []TudouMovieSegs `json:"3"`
	F4v_2     []TudouMovieSegs `json:"2"`
}

func (this *Tudou) GetVideo() (string, error) {

	url := "http://www.tudou.com/outplay/goto/getItemSegs.action?iid=" + this.Videocode

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return "", err
	}

	json.Unmarshal(content, this)

	result := this.getData()

	return result, nil
}

func (this *Tudou) getData() string {

	url := "http://vr.tudou.com/v2proxy/v?sid=95000&st=%d&id=%d"

	var vJson VideoJson
	//标清
	if len(this.F4v_2) > 0 {
		vJson.Fluent.M3u8 = this.getM3u8("2")
		for _, v := range this.F4v_2 {

			var temp VideoFileId
			temp.Url = fmt.Sprintf(url, 2, v.Id)
			temp.Size = fmt.Sprintf("%d", v.Size)
			temp.Seconds = fmt.Sprintf("%d", v.Seconds/1000)

			vJson.Fluent.F4v = append(vJson.Fluent.F4v, temp)

		}
	}

	//高清
	if len(this.F4v_3) > 0 {
		vJson.Height.M3u8 = this.getM3u8("3")
		for _, v := range this.F4v_3 {

			var temp VideoFileId
			temp.Url = fmt.Sprintf(url, 3, v.Id)
			temp.Size = fmt.Sprintf("%d", v.Size)
			temp.Seconds = fmt.Sprintf("%d", v.Seconds/1000)

			vJson.Height.F4v = append(vJson.Height.F4v, temp)

		}
	}

	//超清
	if len(this.F4v_5) > 0 {
		vJson.Super.M3u8 = this.getM3u8("5")
		for _, v := range this.F4v_5 {
			var temp VideoFileId
			temp.Url = fmt.Sprintf(url, 5, v.Id)
			temp.Size = fmt.Sprintf("%d", v.Size)
			temp.Seconds = fmt.Sprintf("%d", v.Seconds/1000)

			vJson.Super.F4v = append(vJson.Super.F4v, temp)
		}
	}

	//转json字符串
	temp, err := json.Marshal(vJson)
	if err != nil {
		return ""
	}
	vstr := string(temp)

	return vstr

}

//组装成m3u8数据
func (this *Tudou) getM3u8(st string) string {

	m3u8 := "http://vr.tudou.com/v2proxy/v2.m3u8?it=" + this.Videocode + "&st=" + st + "&pw="

	return m3u8
}
