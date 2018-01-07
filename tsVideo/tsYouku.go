package tsVideo

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/utils/pagination"
	"net/url"
	"tsEngine/tsCrypto"
	"tsEngine/tsTime"
)

const (
	Youku_jsonKey = "qwer3as2jin4fdsa"
	Youku_apiKey  = "631l1i1x3fv5vs2dxlj5v8x81jqfs2om"
	Youku_btanKey = "9e3633aadde6bfec"
	Youku_time    = 3600
	//Youku_cookie  = "v=UMTU1ODcwODc2__null____fc206c59c6dd97df7b69fece;ykss=fc206c59c6dd97df7b69fece;u=ninowolf;k=ninowolf;_1=1;YOUKUSESSID=yktk%3Dnull;P_sck=MS4wLjB8MjAxNjEwMTJBUFAwMDA3MDd8UGN3cFdJYlZOQlNwZmJ2eHxXTWU2Sk4rNUx0UURBQVVl%0AQ2tLS25telF8MTUwMDI1ODU1NjAxN3wzODQ2MTUwMDl8QW5kcm9pZHw0LjQuMnxDQUE1OTE5QUZB%0AMzg0REIwQ0Q5NTdGMEJBRTIzQ0YzQQ%3D%3D%0A"
)

type HttpData struct {
	BlankNum int64  `json:"blank_num"`
	Data     string `json:"data"`
}

type MovieFileId struct {
	Url     string `json:"url"`
	Seconds int64  `json:"seconds"`
	Fileid  string `json:"fileid"`
	Id      int64  `json:"id"`
	Size    int64  `json:"size"`
}

type MovieNoFileId struct {
	Url     string `json:"url"`
	Seconds int64  `json:"seconds"`
	Id      int64  `json:"id"`
	Size    int64  `json:"size"`
}

type tYouKuMovie struct {
	Mp4      []MovieFileId   `json:"mp4"`
	Hd2      []MovieFileId   `json:"hd2"`
	M3u8     []MovieNoFileId `json:"m3u8"`
	M3u8_mp4 []MovieNoFileId `json:"m3u8_mp4"`
	M3u8_flv []MovieNoFileId `json:"m3u8_flv"`
	Flvhd    []MovieFileId   `json:"flvhd"`
	Flv      []MovieFileId   `json:"flv"`
	Gphd3    []MovieFileId   `json:"3gphd"`
}

type tstreamlogos struct {
	Hd2   int64 `json:"hd2"`
	Mp4   int64 `json:"mp4"`
	Gphd3 int64 `json:"3gphd"`
	Flvhd int64 `json:"flvhd"`
	Flv   int64 `json:"flv"`
}

type tstream_milliseconds struct {
	Mp4   int64 `json:"mp4"`
	Gphd3 int64 `json:"3gphd"`
	Flvhd int64 `json:"flvhd"`
	Flv   int64 `json:"flv"`
}

type tsid_data struct {
	Token string `json:"token"`
	Oip   int64  `json:"oip"`
	Sid   string `json:"sid"`
}

type next_video struct {
	Videoid      string `json:"videoid"`
	ShowVideoseq int64  `json:"show_videoseq"`
	Title        string `json:"title"`
}

type points struct {
	Start float64 `json:"start"`
	Type  string  `json:"type"`
	Title string  `json:"title"`
	Desc  string  `json:"desc"`
}
type Youku struct {
	Videocode string
	Videoid   int64
	Cacheid   string

	NextVideo           next_video           `json:"next_video"`
	Exclusive           bool                 `json:"exclusive"`
	Uid                 string               `json:"uid"`
	Siddecode           string               `json:"siddecode"`
	Results             tYouKuMovie          `json:"results"`
	Drm_type            string               `json:"drm_type"`
	Cs                  string               `json:"cs"`
	Sct                 string               `json:"sct"`
	Streamlogos         tstreamlogos         `json:"streamlogos"`
	Scs                 string               `json:"scs"`
	Ct                  string               `json:"ct"`
	Copyright           int64                `json:"copyright"`
	Title               string               `json:"title"`
	Pcs                 string               `json:"pcs"`
	Interact            bool                 `json:"interact"`
	Pct                 string               `json:"pct"`
	Stream_milliseconds tstream_milliseconds `json:"stream_milliseconds"`
	Type_arr            []string             `json:"type_arr"`
	Sid_data            tsid_data            `json:"sid_data"`
	Show_icon           int64                `json:"show_icon"`
	Viddecode           string               `json:"viddecode"`
	Status              string               `json:"status"`
	Weburl              string               `json:"weburl"`
	Img_hd              string               `json:"img_hd"`
	Show_videoseq       int                  `json:"show_videoseq"`
	Paid                int64                `json:"paid"`
	Trailers            string               `json:"trailers"`
	Play_u_state        int64                `json:"play_u_state"`
	Video_type          int64                `json:"video_type"`
	Paystate            string               `json:"paystate"`
	Cid                 int                  `json:"cid"`
	Videoid_play        string               `json:"videoid_play"`
	Is_phone_stream     string               `json:"is_phone_stream"`
	Points              []points             `json:"points"`
	Panorama            bool                 `json:"panorama"`
	Totalseconds        int64                `json:"totalseconds"`
}

//获取优酷视频数据
func (this *Youku) GetVideo() (string, error) {

	url := "http://a.play.api.3g.youku.com/common/v3/play?"

	value := fmt.Sprintf("%d", tsTime.CurrSe())

	str := "GET:/common/v3/play:" + value + ":" + Youku_apiKey

	time := tsCrypto.GetMd5([]byte(str))

	url += "_t_=" + value
	url += "&e=md5"
	url += "&_s_=" + time
	url += "&point=1"
	url += "&id=" + this.Videocode
	url += "&local_time="
	url += "&local_vid="
	url += "&format=1,5,6,7,8"
	url += "&did=" + time
	url += "&language=default"
	url += "&ctype=20"
	url += "&local_point="
	url += "&audiolang=1"
	url += "&pid=4e308edfc33936d7"
	url += "&guid=5663956e2a6d8361610deb9384db5ebe"
	url += "&mac=2C:44:01:E1:98:9E"
	url += "&imei=357980042532616"
	url += "&ver=5.4.1"

	//建立一个请求对象
	curl := httplib.Get(url)

	curl.Header("User-Agent", "Youku;6.8.0;iOS;10.3.2;iPhone6,2")
	curl.Header("Cookie", this.getVipCookie())

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return "", err
	}

	var oHttpData HttpData

	json.Unmarshal(content, &oHttpData)

	if oHttpData.Data == "" {
		return "", nil
	}
	//将json的data 值，进行aes解密
	oAes := tsCrypto.AesGB{Strkey: Youku_jsonKey}
	data, err := oAes.Decrypt(oHttpData.Data)

	if err != nil {
		return "", err
	}

	json.Unmarshal(data, this)

	result := this.getData(time)

	return result, nil

}

//优酷会员电影数据结构
type YmovieField struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Img    string  `json:"thumburl"`
	Vid    string  `json:"showid"`
	Score  float64 `json:"reputation"`
	Visits string  `json:"showweek_all_vv"`
}

type YmovieResult struct {
	TotalSize int64         `json:"totalSize"`
	List      []YmovieField `json:"result"`
}

type YmovieList struct {
	State  string       `json:"code"`
	Result YmovieResult `json:"result"`
}

//获取优酷会员电影列表数据（带分页默认 500条）
func (this *Youku) GetYoukuData(ctx *context.Context, category string, order string, page string) []YmovieField {

	var oHot YmovieList

	url := "http://vip.youku.com/?app=newvip&c=svip&a=listShow&pt=1&pl=50&mg=%s&r=0&ar=0&o=%s&pn=%s"

	url = fmt.Sprintf(url, category, order, page)

	//建立一个请求对象
	curl := httplib.Get(url)
	curl.Header("Cookie", this.getVipCookie())

	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oHot)

	total := oHot.Result.TotalSize
	if category == "0" {
		total = oHot.Result.TotalSize/2 - 50
	}

	pagination.SetPaginator(ctx, 50, total)

	return oHot.Result.List

}

//优酷视频查询数据结构

type YsearchField struct {
	Id         string
	Name       string  `json:"showname"`
	Img        string  `json:"show_vthumburl"`
	Vid        string  `json:"showid"`
	Score      float64 `json:"reputation"`
	Visits     string  `json:"showtotal_vv"`
	Isyouku    int     `json:"is_youku"`
	Updatetime string  `json:"summary"`
	Type       int
}

type YsearchList struct {
	Status string         `json:"status"`
	Page   int            `json:"pg"`
	Psize  int            `json:"pz"`
	Total  int64          `json:"total"`
	List   []YsearchField `json:"results"`
}

//优酷视频查询列表
func (this *Youku) GetYoukuSearch(ctx *context.Context, s string, p string) []YsearchField {

	if p == "" {
		p = "1"
	}

	//url := "http://search.api.3g.youku.com/videos/search/" + url.QueryEscape(s) + "?"
	url := "http://search.api.3g.youku.com/layout/ios/v3/search/direct_all/" + url.QueryEscape(s) + "?"

	value := fmt.Sprintf("%d", tsTime.CurrSe())

	str := "GET:/layout/ios/v3/search/direct_all/" + s + ":" + value + ":" + Youku_apiKey
	//str := "GET:/videos/search/" + s + ":" + value + ":" + Youku_apiKey

	md5 := tsCrypto.GetMd5([]byte(str))

	url += "area_code=1"
	url += "&brand=apple"
	url += "&btype=iPhone6%2C2"
	url += "&deviceid=0f607264fc6318a92b9e13c65db7cd3c"
	url += "&guid=5663956e2a6d8361610deb9384db5ebe"
	url += "&idfa=7066707c5bdc38af1621eaf94a6fe779"
	url += "&image_hd=1"
	url += "&network=WIFI"
	url += "&operator=%E4%B8%AD%E5%9B%BD%E7%A7%BB%E5%8A%A8_46002"
	url += "&os=ios"
	url += "&os_ver=9.3.1"
	url += "&ouid=8c2b7eecc41f84f66793ce5a33d4e39639b59079"
	url += "&pg=" + p
	url += "&pid=69b81504767483cf"
	url += "&pz=10"
	url += "&relationship=1"
	url += "&scale=2"
	url += "&vdid=D1F3D755-3C2D-4443-9750-9543CBE2050D"
	url += "&_s_=" + md5
	url += "&_t_=" + value
	url += "&ver=5.5.2"

	var oHot YsearchList

	//建立一个请求对象
	curl := httplib.Get(url)
	curl.Header("Cookie", this.getVipCookie())
	curl.Header("User-Agent", "Youku;5.5.2;iPhone OS;9.3.1;iPhone6,2")

	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oHot)

	var nHot []YsearchField
	for _, v := range oHot.List {

		if v.Isyouku == 1 && v.Vid != "" {
			var info YsearchField
			info.Id = v.Vid
			info.Img = v.Img
			info.Name = v.Name
			info.Score = v.Score
			info.Type = v.Type
			info.Vid = v.Vid
			info.Updatetime = v.Updatetime
			nHot = append(nHot, info)
		}
	}

	pagination.SetPaginator(ctx, oHot.Psize, oHot.Total)

	return nHot

}

//获取排行列表数据结构
type YoukuTopField struct {
	Name          string `json:"title"`
	Show_thumburl string `json:"img"`
	Total_vv      string `json:"subtitle"`
	Tid           string `json:"tid"`
}

type YoukuTop struct {
	State string          `json:"status"`
	List  []YoukuTopField `json:"results"`
}

//获取优酷排行列表数据 （97 电视剧   96 电影  85 综艺 99游戏 91资讯）
func (this *Youku) GetYoukuTop(category string) ([]YoukuTopField, error) {

	var oTop YoukuTop

	url := "http://api.mobile.youku.com/layout/phone/channel/rank?"

	value := fmt.Sprintf("%d", tsTime.CurrSe())

	str := "GET:/layout/phone/channel/rank:" + value + ":" + Youku_apiKey

	md5 := tsCrypto.GetMd5([]byte(str))

	url += "brand=apple"
	url += "&btype=iPhone6%2C2"
	url += "&cid=" + category
	url += "&deviceid=0f607264fc6318a92b9e13c65db7cd3c"
	url += "&guid=7066707c5bdc38af1621eaf94a6fe779"
	url += "&idfa=19A0FBA9-C4C8-423D-AE48-7C9D6EBD9688"
	url += "&image_hd=1"
	url += "&network=WIFI"
	url += "&operator=%E4%B8%AD%E5%9B%BD%E7%A7%BB%E5%8A%A8_46002&os=ios"
	url += "&os_ver=9.3.2"
	url += "&ouid=8c2b7eecc41f84f66793ce5a33d4e39639b59079"
	url += "&pid=69b81504767483cf&scale=2&vdid=D1F3D755-3C2D-4443-9750-9543CBE2050D"
	url += "&ver=5.7.3"
	url += "&_s_=" + md5
	url += "&_t_=" + value

	//建立一个请求对象
	//fmt.Println("地址：", url)
	curl := httplib.Get(url)
	curl.Header("Cookie", this.getVipCookie())
	curl.Header("User-Agent", "Youku;5.5.2;iPhone OS;9.3.1;iPhone6,2")
	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oTop)

	return oTop.List, nil
}

//排行单个视频数据结构
type YoukuTopViewField struct {
	Title   string `json:"title"`
	Episode int    `json:"show_videoseq"`
	Code    string `json:"videoid"`
}

type YoukuTopView struct {
	State string              `json:"status"`
	Pz    int                 `json:"pz"`
	Total int                 `json:"total"`
	List  []YoukuTopViewField `json:"results"`
}

//获取优酷排行单个视频的数据
func (this *Youku) GetYoukuTopView(tid string, category string) ([]YoukuTopViewField, error) {

	var oTop YoukuTopView

	url := "http://api.mobile.youku.com/shows/" + tid + "/reverse/videos?"

	value := fmt.Sprintf("%d", tsTime.CurrSe())

	str := "GET:/shows/" + tid + "/reverse/videos:" + value + ":" + Youku_apiKey

	md5 := tsCrypto.GetMd5([]byte(str))

	url += "pid=0865e0628a79dfbb"
	url += "&guid=d5ea03455ee9b2d749459ed8256df739"
	url += "&mac=4e%3A18%3A20%3Add%3A57%3A87"
	url += "&imei=753346791151434"
	url += "&ver=4.8.1"
	url += "&e=md5"
	url += "&_s_=" + md5
	url += "&_t_=" + value
	url += "&network=WIFI"
	url += "&operator=%E4%B8%AD%E5%9B%BD%E7%A7%BB%E5%8A%A8_46002"
	url += "&fields=vid|titl|lim|is_new"
	url += "&pz=100"
	url += "&area_code=1"
	url += "&pg=1"

	//建立一个请求对象
	curl := httplib.Get(url)
	curl.Header("Cookie", this.getVipCookie())
	curl.Header("User-Agent", "Youku;5.7.3;iPhone OS;9.3.2;iPhone6,2")
	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oTop)

	return oTop.List, nil
}

//手机版获取m3u8文件（必须使用此方法否则只有前10分钟数据）
func (this *Youku) GetM3u8File(m3u8url string) string {

	//先去获取cookie
	url := "http://user-mobile.youku.com/user/updatecookie?"
	value := fmt.Sprintf("%d", tsTime.CurrSe())
	str := "GET:/user/updatecookie:" + value + ":" + Youku_apiKey
	md5 := tsCrypto.GetMd5([]byte(str))

	url += "brand=apple&btype=iPhone6%2C2&deviceid=0f607264fc6318a92b9e13c65db7cd3c&guid=7066707c5bdc38af1621eaf94a6fe779"
	url += "&idfa=19A0FBA9-C4C8-423D-AE48-7C9D6EBD9688&network=WIFI&operator=%E4%B8%AD%E5%9B%BD%E7%A7%BB%E5%8A%A8_46002&os=ios&os_ver=9.3.1"
	url += "&ouid=8c2b7eecc41f84f66793ce5a33d4e39639b59079&pid=69b81504767483cf&scale=2&vdid=D1F3D755-3C2D-4443-9750-9543CBE2050D&ver=5.5.2"
	url += "&_s_=" + md5
	url += "&_t_=" + value

	//建立一个请求对象
	curl := httplib.Get(url)
	curl.Header("Cookie", this.getVipCookie())

	//获取请求的内容
	temp, _ := curl.Response()

	//建立一个请求对象
	curl = httplib.Get(m3u8url)

	curl.Header("Cookie", fmt.Sprintf("%s", temp.Cookies()))

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return ""
	}
	return string(content)
}

//第二种view的解析
type YoukuViewField struct {
	Tid string `json:"showid"`
	Vid string `json:"videoid"`
}

type YoukuViewInfo struct {
	State  string         `json:"status"`
	Detail YoukuViewField `json:"detail"`
}

func (this *Youku) GetYoukuView(tid string, category string) ([]YoukuViewListField, error) {

	var oView YoukuViewInfo

	url := "http://api.mobile.youku.com/layout/ios5_0/play/detail?"

	value := fmt.Sprintf("%d", tsTime.CurrSe())

	str := "GET:/layout/ios5_0/play/detail:" + value + ":" + Youku_apiKey

	md5 := tsCrypto.GetMd5([]byte(str))

	url += "area_code=1&brand=apple&btype=iPhone6%2C2&deviceid=0f607264fc6318a92b9e13c65db7cd3c&guid=7066707c5bdc38af1621eaf94a6fe779"
	url += "&idfa=19A0FBA9-C4C8-423D-AE48-7C9D6EBD9688&network=WIFI&operator=%E4%B8%AD%E5%9B%BD%E7%A7%BB%E5%8A%A8_46002&os=ios&os_ver=9.3.1"
	url += "&ouid=8c2b7eecc41f84f66793ce5a33d4e39639b59079&pid=69b81504767483cf&scale=2&vdid=D1F3D755-3C2D-4443-9750-9543CBE2050D&ver=5.5.2"
	url += "&id=" + tid
	url += "&_s_=" + md5
	url += "&_t_=" + value

	//建立一个请求对象
	curl := httplib.Get(url)
	curl.Header("Cookie", this.getVipCookie())
	curl.Header("User-Agent", "Youku;5.5.2;iPhone OS;9.3.1;iPhone6,2")
	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oView)

	return this.getViewList(oView.Detail.Tid, category), nil

}

//获取view的列表
type YoukuViewListField struct {
	Title   string `json:"title"`
	Episode int    `json:"video_stage"`
	Code    string `json:"videoid"`
}

type YoukuViewList struct {
	State string               `json:"status"`
	List  []YoukuViewListField `json:"results"`
}

func (this *Youku) getViewList(tid string, category string) []YoukuViewListField {

	var oView YoukuViewList

	url := "http://api.mobile.youku.com/layout/phone3_0/shows/" + tid + "/series?"

	value := fmt.Sprintf("%d", tsTime.CurrSe())

	str := "GET:/layout/phone3_0/shows/" + tid + "/series:" + value + ":" + Youku_apiKey

	md5 := tsCrypto.GetMd5([]byte(str))

	url += "brand=apple&btype=iPhone6%2C2&deviceid=0f607264fc6318a92b9e13c65db7cd3c&guid=7066707c5bdc38af1621eaf94a6fe779"
	url += "&idfa=19A0FBA9-C4C8-423D-AE48-7C9D6EBD9688&network=WIFI&operator=%E4%B8%AD%E5%9B%BD%E7%A7%BB%E5%8A%A8_46002&os=ios&os_ver=9.3.1"
	url += "&ouid=8c2b7eecc41f84f66793ce5a33d4e39639b59079&pid=69b81504767483cf&scale=2&vdid=D1F3D755-3C2D-4443-9750-9543CBE2050D&ver=5.5.2"
	url += "&id=" + tid
	url += "&_s_=" + md5
	url += "&_t_=" + value

	//建立一个请求对象
	curl := httplib.Get(url)
	curl.Header("Cookie", this.getVipCookie())
	curl.Header("User-Agent", "Youku;5.5.2;iPhone OS;9.3.1;iPhone6,2")
	//获取请求的内容
	content, _ := curl.Bytes()

	json.Unmarshal(content, &oView)

	return oView.List

}

//拼装flv 和 m3u8数据的内部方法
func (this *Youku) getData(time string) string {

	var vJson VideoJson

	//标清
	if len(this.Results.Flvhd) > 0 {
		vJson.Fluent.M3u8 = this.getM3u8("flvhd", time)

		for _, v := range this.Results.Flvhd {

			var temp VideoFileId
			temp.Url = this.getEpData(v.Url, v.Fileid, time)
			temp.Size = fmt.Sprintf("%d", v.Size)
			temp.Seconds = fmt.Sprintf("%d", v.Seconds)

			vJson.Fluent.F4v = append(vJson.Fluent.F4v, temp)

		}
	}

	//高清
	if len(this.Results.Mp4) > 0 {
		vJson.Height.M3u8 = this.getM3u8("mp4", time)
		for _, v := range this.Results.Mp4 {
			var temp VideoFileId
			temp.Url = this.getEpData(v.Url, v.Fileid, time)
			temp.Size = fmt.Sprintf("%d", v.Size)
			temp.Seconds = fmt.Sprintf("%d", v.Seconds)

			vJson.Height.F4v = append(vJson.Height.F4v, temp)

		}
	}

	//超清
	if len(this.Results.Hd2) > 0 {
		vJson.Super.M3u8 = this.getM3u8("hd2", time)
		for _, v := range this.Results.Hd2 {
			var temp VideoFileId
			temp.Url = this.getEpData(v.Url, v.Fileid, time)
			temp.Size = fmt.Sprintf("%d", v.Size)
			temp.Seconds = fmt.Sprintf("%d", v.Seconds)

			vJson.Super.F4v = append(vJson.Super.F4v, temp)

		}
	}

	//转json字符串
	temp, err := json.Marshal(&vJson)
	if err != nil {
		return ""
	}
	vstr := string(temp)

	return vstr

}

//组装成m3u8数据
func (this *Youku) getM3u8(vtype string, time string) string {

	url := "http://pl.youku.com/playlist/m3u8?vid=" + this.Videocode + "&type=" + vtype + "&ts=" + time + "&keyframe=0"

	m3u8 := this.getEpData(url, this.Videocode, time)

	return m3u8
}

//优酷EP算法
func (this *Youku) getEpData(s string, s2 string, s6 string) string {

	s3 := this.Sid_data.Token
	s4 := fmt.Sprintf("%d", this.Sid_data.Oip)
	s5 := this.Sid_data.Sid
	str := s5 + "_" + s2 + "_" + s3
	oAes := tsCrypto.AesGB{Strkey: Youku_btanKey}

	data, _ := oAes.Encrypt([]byte(str))
	urlstr := s + "&oip=" + s4 + "&sid=" + s5 + "&token=" + s3 + "&did=" + s6 + "&ev=1&ctype=20&ep=" + url.QueryEscape(string(data))
	return urlstr
}

//获取vip登录cookie的内部方法
func (this *Youku) getVipCookie() string {

	time := fmt.Sprintf("%d", tsTime.CurrSe())

	cookie := beego.AppConfig.String("YkCookie") + time

	return cookie
}
