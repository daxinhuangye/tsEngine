package tsWx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"net/url"
	"strings"
)

const (
	UrlPrefix      = "https://api.weixin.qq.com/"
	MediaUrlPrefix = "http://file.api.weixin.qq.com/cgi-bin/media/"
	retryNum       = 3
)

type WxData struct {
	ErrCode      int64  `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

func GetWxData(this *context.Context, appid, secret, snsapi, subUrl string) WxData {

	var data WxData

	if code := this.Input.Query("code"); code == "" {
		encodUrl := url.QueryEscape(subUrl)
		urlStr := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + appid + "&redirect_uri=" + encodUrl + "&response_type=code&scope=" + snsapi + "&state=STATE#wechat_redirect"
		this.Redirect(302, urlStr)

	} else {
		beego.Trace("获取token", code, appid, secret)
		if data, err := GetFromoauth2(code, appid, secret); err == nil {
			return data
		}
	}
	return data
}

type UserInfo struct {
	Subscribe     int64  `json:"subscribe"`
	Openid        string `json:"openid"`
	Nickname      string `json:"nickname"`
	Sex           int64  `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	Headimgurl    string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`
	Unionid       string `json:"unionid"`
}

func GetUserInfo(openId, token string) (UserInfo, error) {
	var uinf UserInfo

	url := fmt.Sprintf("%ssns/userinfo?lang=zh_CN&openid=%s&access_token=%s", UrlPrefix, openId, token)

	beego.Trace(url + token)

	// retry
	for i := 0; i < retryNum; i++ {

		//建立一个请求对象
		curl := httplib.Get(url)

		//获取请求的内容
		data, err := curl.Bytes()

		if err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}

		// has error?
		var rtn WxData
		if err := json.Unmarshal(data, &rtn); err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}
		// yes
		if rtn.ErrCode != 0 {
			if i < retryNum-1 {
				continue
			}
			return uinf, errors.New(fmt.Sprintf("%d %s", rtn.ErrCode, rtn.ErrMsg))
		}
		// no
		if err := json.Unmarshal(data, &uinf); err != nil {
			if i < retryNum-1 {
				continue
			}
			return uinf, err
		}
		break // success
	}
	return uinf, nil
}

func GetJsApiTicket(timestamp uint64, noncestr string, web_url string) (string, error) {

	api_url := "http://mp.wememe.cn/api/getjsapiticket?"

	api_url += "timestamp=" + fmt.Sprintf("%d", timestamp)

	api_url += "&noncestr=" + noncestr

	api_url += "&url=" + url.QueryEscape(web_url)

	fmt.Println(api_url)
	//建立一个请求对象
	curl := httplib.Get(api_url)

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return "", err
	}

	return string(content), nil

}

func GetFromoauth2(code, appid, secret string) (WxData, error) {

	var data WxData

	url := strings.Join([]string{"https://api.weixin.qq.com/sns/oauth2/access_token",
		"?appid=", appid,
		"&secret=", secret,
		"&code=", code,
		"&grant_type=authorization_code"}, "")

	//建立一个请求对象
	curl := httplib.Get(url)

	//获取请求的内容
	content, err := curl.Bytes()

	if err != nil {
		return data, err
	}

	beego.Trace(string(content))

	err = json.Unmarshal(content, &data)

	if err != nil {
		beego.Trace(err)
		return data, err
	}

	return data, nil
}
