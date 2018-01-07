package tsVideo

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/utils/pagination"
	"net/url"
	"strings"
	"tsEngine/tsCrypto"
	"tsEngine/tsTime"
)

const (
	Mgtv_time   = 7200
	Mgtv_uid    = "bc3dc8c76766707f6bf1680745e80489"
	Mgtv_ticket = "9935ETIVBLDYGRBZA5BC"
)

type MgtvFileId struct {
	Definition int    `json:"definition"`
	Name       string `json:"name"`
	Url        string `json:"url"`
}

type tMgMovie struct {
	VideoDomains []string     `json:"videoDomains"`
	VideoSources []MgtvFileId `json:"videoSources"`
}

type Mgtv struct {
	Videocode  string
	Videoid    int64
	Cacheid    string
	Definition string
	Err_code   int      `json:"err_code"`
	Err_msg    string   `json:"err_msg"`
	Data       tMgMovie `json:"data"`
}

//获取优酷视频数据
func (this *Mgtv) GetVideo(definition string) (string, error) {

	this.Definition = definition

	err := this.setTicket()
	if err != nil {
		return "", err
	}

	did := "i421826490416727"
	dtime := fmt.Sprintf("%d", tsTime.CurrSe())
	seq_id := tsCrypto.GetMd5([]byte(did + "." + dtime))

	url := "http://mobile.api.hunantv.com/v5/video/getSource?"

	url += "uid=" + Mgtv_uid
	url += "&osVersion=4.4.2"
	url += "&ticket=" + Mgtv_ticket
	url += "&appVersion=4.6.8"
	url += "&videoId=" + this.Videocode
	url += "&userId="
	url += "&device=GT-N7100"
	url += "&mac=" + did
	url += "&osType=android"
	url += "&seqId=" + seq_id
	url += "&channel=baidu"

	//建立一个请求对象
	curl := httplib.Get(url)
	fmt.Println("api:", url)
	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return "", err
	}

	json.Unmarshal(content, this)

	result := this.getData()

	return result, nil

}

type MmovieList struct {
	Data MmovieData `json:"data"`
	Code int        `json:"code"`
}

type MmovieData struct {
	List      []MmovieField `json:"hitDocs"`
	TotalSize int64         `json:"hitCount"`
}

type MmovieField struct {
	Id        int64  `json:"id"`
	Name      string `json:"title"`
	Img       string `json:"imgsmall"`
	Vid       string `json:"indexfile"`
	Performer string `json:"player"`
}

func (this *Mgtv) GetVipList(ctx *context.Context, category string, order string, page string) []MmovieField {

	var oHot MmovieList

	url := "http://mgso.hunantv.com/list?src=pc&ic=1&ty=3&iv=1&pc=36&pn=%s&st=%s&vip=1&_=" + fmt.Sprintf("%d", tsTime.CurrSe())
	url = fmt.Sprintf(url, page, order)
	if category != "" {
		url += "&tid=" + category
	}

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oHot)

	pagination.SetPaginator(ctx, 36, oHot.Data.TotalSize)

	return oHot.Data.List

}

type MsearchData struct {
	Data MsearchList `json:"data"`
	Code int64       `json:"code"`
}

type MsearchList struct {
	List []MsearchField `json:"relatedVideo"`
}
type MsearchField struct {
	Vid     int64  `json:"videoId"`
	Name    string `json:"name"`
	Img     string `json:"image"`
	Visits  string `json:"playCount"`
	Channel string `json:"channel"`
}

func (this *Mgtv) GetMgtvSearch(s string) []MsearchField {

	var oHot MsearchData
	//http://m.api.hunantv.com/search/getResult?name=%E7%96%AF%E7%8B%82&more=1&sortOrder=2&pageNum=2
	url := "http://m.api.hunantv.com/search/getResult?name=" + url.QueryEscape(s)

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oHot)

	var list []MsearchField

	for _, v := range oHot.Data.List {
		if v.Channel == "电影" {
			list = append(list, v)
		}
	}

	return list

}

func (this *Mgtv) setTicket() error {

	did := "i421826490416727"
	dtime := fmt.Sprintf("%d", tsTime.CurrSe())
	seq_id := tsCrypto.GetMd5([]byte(did + "." + dtime))

	url := "http://mobile.api.hunantv.com/user/getUserInfo?"
	url += "uid=" + Mgtv_uid
	url += "&osVersion=4.4.2"
	url += "&appVersion=4.6.8"
	url += "&videoId=" + this.Videocode
	url += "&userId="
	url += "&device=GT-N7100"
	url += "&mac=" + did
	url += "&osType=android"
	url += "&seqId=" + seq_id
	url += "&ticket=" + Mgtv_ticket

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	_, err := curl.Bytes()
	return err
}

type M3data struct {
	Info string `json:"info"`
}

//拼装flv 和 m3u8数据的内部方法
func (this *Mgtv) getData() string {

	var m3data M3data

	url := this.Data.VideoDomains[0] + this.Data.VideoSources[0].Url
	if this.Definition == "3" {
		url = this.Data.VideoDomains[0] + this.Data.VideoSources[0].Url
	} else if this.Definition == "2" {
		url = this.Data.VideoDomains[0] + this.Data.VideoSources[1].Url
	} else if this.Definition == "1" {
		url = this.Data.VideoDomains[0] + this.Data.VideoSources[2].Url
	}

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return ""
	}

	json.Unmarshal(content, &m3data)

	//建立一个请求对象
	curl = httplib.Get(m3data.Info)

	//获取请求的内容
	content, err = curl.Bytes()

	if err != nil {
		return ""
	}

	fmt.Println("m3u8:", m3data.Info)

	temp := strings.Split(m3data.Info, "playlist.m3u8")

	murl := ",\n" + temp[0]

	m3content := strings.Replace(string(content), ",\n", murl, -1)

	return m3content

}
