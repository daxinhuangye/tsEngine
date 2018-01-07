package tsVideo

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"net/url"
	"tsEngine/tsCrypto"
	"tsEngine/tsTime"
)

const (
	//Acfun_jsonKey = "qwer3as2jin4fdsa"
	Acfun_jsonKey = "ELSH6ruK0qva88DD"
	Acfun_apiKey  = "78554907b127c3853f8e956243dc74c4"
	Acfun_btanKey = "zx26mfbsuebv72ja"
	Acfun_time    = 3600
)

type acHttpData struct {
	BlankNum int64  `json:"blank_num"`
	Data     string `json:"data"`
}

type acF4vFileId struct {
	Url     string `json:"url"`
	Seconds int64  `json:"seconds"`
	Fileid  string `json:"fileid"`
	Id      int64  `json:"id"`
	Size    int64  `json:"size"`
}

type acM3u8FileId struct {
	Url     string `json:"url"`
	Seconds int64  `json:"seconds"`
	Id      int64  `json:"id"`
	Size    int64  `json:"size"`
}

type acFunMovie struct {
	Mp4   []acF4vFileId `json:"mp4"`
	Hd2   []acF4vFileId `json:"hd2"`
	Flvhd []acF4vFileId `json:"flvhd"`
	Flv   []acF4vFileId `json:"flv"`
	Gphd3 []acF4vFileId `json:"3gphd"`

	M3u8     []acM3u8FileId `json:"m3u8"`
	M3u8_mp4 []acM3u8FileId `json:"m3u8_mp4"`
	M3u8_flv []acM3u8FileId `json:"m3u8_flv"`
}

type acSid_data struct {
	Token string `json:"token"`
	Oip   int64  `json:"oip"`
	Sid   string `json:"sid"`
}

type Acfun struct {
	Videocode string

	Results acFunMovie `json:"results"`

	Sid_data      acSid_data `json:"sid_data"`
	Viddecode     string     `json:"viddecode"`
	Status        string     `json:"status"`
	Weburl        string     `json:"weburl"`
	Show_videoseq int        `json:"show_videoseq"`
	Paystate      string     `json:"paystate"`
	Totalseconds  int64      `json:"totalseconds"`
}

func (this *Acfun) GetVideo() (string, error) {

	url := "http://acfun.api.mobile.youku.com/common/partner/play?"

	value := fmt.Sprintf("%d", tsTime.CurrSe())

	str := "GET:/common/partner/play:" + value + ":" + Acfun_apiKey

	time := tsCrypto.GetMd5([]byte(str))

	url += "_t_=" + value
	url += "&e=md5"
	url += "&_s_=" + time
	url += "&point=1"
	url += "&id=" + this.Videocode
	url += "&local_time="
	url += "&local_vid="
	url += "&format=1,2,3,4,5,6,7,8,9"
	url += "&did=" + time
	url += "&language=default"
	url += "&ctype=87"
	url += "&local_point="
	url += "&audiolang=1"
	url += "&pid=4e308edfc33936d7"
	url += "&guid=5663956e2a6d8361610deb9384db5ebe"
	url += "&mac=2C:44:01:E1:98:9E"
	url += "&imei=357980042532616"
	url += "&ver=260"

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return "", err
	}

	var oHttpData acHttpData

	json.Unmarshal(content, &oHttpData)

	//将json的data 值，进行aes解密
	oAes := new(tsCrypto.AesGB)
	oAes.Strkey = Acfun_jsonKey

	//data, err := oAes.Decrypt(oHttpData.Data)
	oHttpData.Data = "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"

	data, err := oAes.Decrypt(oHttpData.Data)

	fmt.Println(string(data))
	if err != nil {
		return "", err
	}

	json.Unmarshal(data, this)

	result := this.getData(time)

	return result, nil

}

type VideoJson struct {
	Super  VideoType `json:"Super"`
	Height VideoType `json:"Height"`
	Fluent VideoType `json:"Fluent"`
}

type VideoType struct {
	F4v  []VideoFileId `json:"F4v"`
	M3u8 string        `json:"M3u8"`
}

type VideoFileId struct {
	Url     string `json:"Url"`
	Seconds string `json:"Seconds"`
	Size    string `json:"Size"`
}

func (this *Acfun) getData(time string) string {

	var vJson VideoJson

	//标清
	if len(this.Results.Flvhd) > 0 {
		vJson.Fluent.M3u8 = this.getEpData(this.Results.Gphd3[0].Url, this.Results.Gphd3[0].Fileid, time)

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
		vJson.Height.M3u8 = this.getEpData(this.Results.Gphd3[0].Url, this.Results.Gphd3[0].Fileid, time)
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
		vJson.Super.M3u8 = this.getEpData(this.Results.Gphd3[0].Url, this.Results.Gphd3[0].Fileid, time)
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
func (this *Acfun) getGphd(vtype string, time string) string {

	url := "http://pl.youku.com/partner/m3u8?vid=" + this.Videocode + "&type=" + vtype + "&ts=" + time + "&keyframe=0"

	m3u8 := this.getEpData(url, this.Videocode, time)
	//fmt.Println(m3u8)
	return m3u8
}

//组装成m3u8数据
func (this *Acfun) getM3u8(vtype string, time string) string {

	url := "http://pl.youku.com/partner/m3u8?vid=" + this.Videocode + "&type=" + vtype + "&ts=" + time + "&keyframe=0"

	m3u8 := this.getEpData(url, this.Videocode, time)
	//fmt.Println(m3u8)
	return m3u8
}

//优酷EP算法
func (this *Acfun) getEpData(s string, s2 string, s6 string) string {

	s3 := this.Sid_data.Token
	s4 := fmt.Sprintf("%d", this.Sid_data.Oip)
	s5 := this.Sid_data.Sid
	str := s5 + "_" + s2 + "_" + s3
	oAes := tsCrypto.AesGB{Strkey: Acfun_btanKey}
	data, _ := oAes.Encrypt([]byte(str))
	urlstr := s + "&oip=" + s4 + "&sid=" + s5 + "&token=" + s3 + "&did=" + s6 + "&ev=1&ctype=87&ep=" + url.QueryEscape(string(data))
	return urlstr
}

/*
<?php
//使用方法 *.php?id=A站的VID
//BY相对湿度
error_reporting(0);
header("Content-Type: text/html; charset=utf-8");
$id=$_GET["id"];
if(strlen(floor($id))=="7"){
$html=file_get_contents('http://www.acfun.tv/video/getVideo.aspx?id='.$id);
$json = json_decode($html);
$sourceId=$json->sourceId;
echo 'VIP：'.$id.' 的C值为 '.$sourceId;
}elseif(strlen(floor($id))=="6"){
$newid=base64_encode($id*4);
echo 'VIP：'.$id.' 的C值为 C'.$newid;
}else{
echo '非正常值，请查询';
}
?>

*/
