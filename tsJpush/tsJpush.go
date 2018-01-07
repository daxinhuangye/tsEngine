package tsJpush

import (
	"fmt"
	"github.com/ylywyn/jpush-api-go-client"
)

const (
	appKey = "f77b0351d9336eb569824cde"
	secret = "41c728c969bc4f4c35bba47b"
)

func SendMessage(title string, content string) (string, error) {
	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	pf.Add(jpushclient.WINPHONE)
	//pf.All()

	//Audience
	var ad jpushclient.Audience
	/*
		s := []string{"1", "2", "3"}
		ad.SetTag(s)
		ad.SetAlias(s)
		ad.SetID(s)
	*/
	ad.All()

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert(title)
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: title})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: title})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: title})

	var msg jpushclient.Message
	msg.Title = title
	msg.Content = content

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()
	fmt.Printf("%s\r\n", string(bytes))

	//push
	c := jpushclient.NewPushClient(secret, appKey)
	str, err := c.Send(bytes)
	return str, err
	/*
		str, err := c.Send(bytes)

		if err != nil {
			fmt.Printf("err:%s", err.Error())
			return err
		} else {
			fmt.Printf("ok:%s", str)
			return err
		}
	*/
}
