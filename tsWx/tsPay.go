package tsWx

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"sort"
	"strings"
	"tsEngine/tsCrypto"
	"tsEngine/tsTime"
)

//post的参数
type UnifyOrderReq struct {
	Appid            string `xml:"appid"`
	Body             string `xml:"body"`
	Mch_id           string `xml:"mch_id"`
	Nonce_str        string `xml:"nonce_str"`
	Notify_url       string `xml:"notify_url"`
	Trade_type       string `xml:"trade_type"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Total_fee        int    `xml:"total_fee"`
	Out_trade_no     string `xml:"out_trade_no"`
	Sign             string `xml:"sign"`
	Openid           string `xml:"openid"`
}

//回调的参数
type UnifyOrderResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
	Code_url    string `xml:"code_url"`
}

type WXPayNotifyReq struct {
	Return_code    string `xml:"return_code"`
	Return_msg     string `xml:"return_msg"`
	Appid          string `xml:"appid"`
	Mch_id         string `xml:"mch_id"`
	Nonce          string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Result_code    string `xml:"result_code"`
	Openid         string `xml:"openid"`
	Is_subscribe   string `xml:"is_subscribe"`
	Trade_type     string `xml:"trade_type"`
	Bank_type      string `xml:"bank_type"`
	Total_fee      int    `xml:"total_fee"`
	Fee_type       string `xml:"fee_type"`
	Cash_fee       int    `xml:"cash_fee"`
	Cash_fee_Type  string `xml:"cash_fee_type"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	Time_end       string `xml:"time_end"`
}

type WXPayNotifyResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}

//订单检测
type WXOrderQuery struct {
	Appid          string `xml:"appid"`
	Mch_id         string `xml:"mch_id"`
	Nonce_str      string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Transaction_id string `xml:"transaction_id"`
}

//微信支付计算签名的函数
func WxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.

	upperSign := strings.ToUpper(tsCrypto.GetMd5([]byte(signStrings)))

	return upperSign
}

//统一下单
func UnifiedOrder(app_id string, mch_id string, pay_key string, callback_url string, body string, trade_type string, total_fee int, nonce_str string, openid ...string) (result UnifyOrderResp, err error) {
	//请求UnifiedOrder的代码

	var yourReq UnifyOrderReq
	yourReq.Appid = app_id
	yourReq.Body = body
	yourReq.Mch_id = mch_id
	yourReq.Nonce_str = nonce_str
	yourReq.Notify_url = callback_url
	yourReq.Trade_type = trade_type
	yourReq.Total_fee = total_fee //单位是分
	yourReq.Out_trade_no = tsCrypto.GetMd5([]byte(fmt.Sprintf("%s-%d", nonce_str, tsTime.CurrMs())))
	yourReq.Openid = openid[0]

	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["appid"] = yourReq.Appid
	m["body"] = yourReq.Body
	m["mch_id"] = yourReq.Mch_id
	m["notify_url"] = yourReq.Notify_url
	m["trade_type"] = yourReq.Trade_type
	m["total_fee"] = yourReq.Total_fee
	m["out_trade_no"] = yourReq.Out_trade_no
	m["nonce_str"] = yourReq.Nonce_str
	m["openid"] = yourReq.Openid

	yourReq.Sign = WxpayCalcSign(m, pay_key) //这个是计算wxpay签名的函数上面已贴出

	xmlResp := UnifyOrderResp{}

	bytes_req, err := xml.Marshal(yourReq)
	if err != nil {
		return xmlResp, err
	}

	str_req := string(bytes_req)
	//wxpay的unifiedorder接口需要http body中xmldoc的根节点是<xml></xml>这种，所以这里需要replace一下
	str_req = strings.Replace(str_req, "UnifyOrderReq", "xml", -1)

	//发送unified order请求.

	//建立一个请求对象
	curl := httplib.Post("https://api.mch.weixin.qq.com/pay/unifiedorder")
	curl.Header("Accept", "application/xml")
	curl.Header("Content-Type", "application/xml;charset=utf-8")
	curl.Body([]byte(str_req))

	//获取请求的内容
	content, err := curl.Bytes()
	if err != nil {
		return xmlResp, err
	}

	xml.Unmarshal(content, &xmlResp)
	return xmlResp, err

}

//具体的微信支付回调函数的范例
func WxpayCallback(body []byte, pay_key string) (res WXPayNotifyReq, res2 string, err error) {

	var mr WXPayNotifyReq
	err = xml.Unmarshal(body, &mr)
	if err != nil {
		return mr, "", err
	}

	var reqMap map[string]interface{}
	reqMap = make(map[string]interface{}, 0)

	reqMap["return_code"] = mr.Return_code
	reqMap["return_msg"] = mr.Return_msg
	reqMap["appid"] = mr.Appid
	reqMap["mch_id"] = mr.Mch_id
	reqMap["nonce_str"] = mr.Nonce
	reqMap["result_code"] = mr.Result_code
	reqMap["openid"] = mr.Openid
	reqMap["is_subscribe"] = mr.Is_subscribe
	reqMap["trade_type"] = mr.Trade_type
	reqMap["bank_type"] = mr.Bank_type
	reqMap["total_fee"] = mr.Total_fee
	reqMap["fee_type"] = mr.Fee_type
	reqMap["cash_fee"] = mr.Cash_fee
	reqMap["cash_fee_type"] = mr.Cash_fee_Type
	reqMap["transaction_id"] = mr.Transaction_id
	reqMap["out_trade_no"] = mr.Out_trade_no
	reqMap["attach"] = mr.Attach
	reqMap["time_end"] = mr.Time_end

	//进行签名校验
	signCalc := WxpayCalcSign(reqMap, pay_key)

	var resp WXPayNotifyResp

	if mr.Sign != signCalc {
		resp.Return_code = "FAIL"
		resp.Return_msg = "failed to verify sign, please retry!"
		err = errors.New("failed to verify sign, please retry!")
	} else {
		resp.Return_code = "SUCCESS"
		resp.Return_msg = "OK"
	}

	bytes, _err := xml.Marshal(resp)
	if err != nil {
		return mr, "", _err
	}
	strResp := strings.Replace(string(bytes), "WXPayNotifyResp", "xml", -1)

	return mr, strResp, err

}
