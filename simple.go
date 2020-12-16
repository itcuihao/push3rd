package gopush

import (
	"fmt"
	"sync"

	"github.com/itcuihao/gopush/apple"
	"github.com/itcuihao/gopush/hwpush"
	"github.com/itcuihao/gopush/jpush"
	"github.com/itcuihao/gopush/wechat"
	"github.com/itcuihao/gopush/xmpush"
)

func checkTitle(t string) string {
	if t == "" {
		t = "温馨提醒"
	}
	return t
}

func sendJAndroid(wg *sync.WaitGroup, title, alert string, devices ...string) {
	jConf := GetJPush()
	defer wg.Done()
	if len(devices) > 0 {
		na := &jpush.NotificationAndroid{
			Alert: alert,
			Title: checkTitle(title),
		}
		result, err := jpush.PushReceptor().
			SetPlatform(jpush.PlatformA).
			SetRegistrationId(devices...).
			SetNotifyAndroid(na).
			Push(jConf.Secret, jConf.AppKey)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result)
	}
}

func sendJIos(wg *sync.WaitGroup, title, alert string, iosProd bool, devices ...string) {
	jConf := GetJPush()
	defer wg.Done()
	if len(devices) > 0 {
		na := &jpush.NotificationApple{
			Alert: map[string]string{"title": checkTitle(title), "body": alert},
		}
		result, err := jpush.PushReceptor().
			SetPlatform(jpush.PlatformI).
			SetRegistrationId(devices...).
			SetNotifyApple(na).
			SetOptionApns(iosProd).
			Push(jConf.Secret, jConf.AppKey)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result)
	}
}

func sendHuaWei(wg *sync.WaitGroup, title, alert string, devices ...string) {
	hwConf := GetHwPush()
	defer wg.Done()
	if len(devices) > 0 {
		title = checkTitle(title)
		hwPkgName := "com.putao.KidReading.bookbook"
		result, err := hwpush.NewPusherV2().
			SetDeviceList(devices...).
			SetMsgType(hwpush.MsgTypeNotify).
			SetMsgTitle(title).
			SetMsgContent(alert).
			SetAction(hwpush.MsgActionApp, hwPkgName).
			Push(hwConf.AppId, hwConf.Secret)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result)
	}
}

func sendXiaoMi(wg *sync.WaitGroup, title, alert string, devices ...string) {
	miConf := GetMiPush()
	defer wg.Done()
	if len(devices) > 0 {
		title = checkTitle(title)
		xmPkgName := "com.putao.KidReading.bookbook"
		result, err := xmpush.NewPushRecptor(xmpush.Version03, false).
			SetRestrictedPackageName(xmPkgName).
			SetPassThrough(xmpush.PassThroughNotify).
			SetTitle(title).
			SetDescription(alert).
			SetRegistrationId(devices...).
			Push(miConf.Secret)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result)
	}
}

func SimpleSend(aDevices,
	iDevices,
	hDevices,
	mDevices []string, title, alert string) {
	var (
		// aDevices []string // 安卓
		// iDevices []string // 苹果
		// hDevices []string // 华为
		// mDevices []string // 小米
		wg = &sync.WaitGroup{}
	)

	wg.Add(4)
	go sendJAndroid(wg, title, alert, aDevices...)
	go sendJIos(wg, title, alert, true, iDevices...)
	go sendHuaWei(wg, title, alert, hDevices...)
	go sendXiaoMi(wg, title, alert, mDevices...)
	wg.Wait()

}

func sendApplePushExample() {
	conf := GetApplePush()
	if err := apple.InitApns(conf.P12Path); err != nil {
		fmt.Println(err.Error())
		return
	}
	apns := apple.NewApns(conf.PackageName, false)
	title := "1"
	body := "1"
	sound := ""
	link := ""
	payload := apple.Payload{
		Aps:  apple.Aps{Alert: apple.Alert{Title: title, Body: body}, Badge: 1, Sound: sound},
		Link: link,
	}
	apns.QuickPush("f9a4040ecd89b0f08511fd479de65ba016f0ac15b25ae57d585b972a49a8e242", payload)
}

func sendWxPushExample() {
	token := "30_EN51TrSEfCrboPg26FmnjkunBPrFxkEGUyiR59vQhY0WkY064uoU0Lv8FdpCsp4uawWszWYI3HieT10BMWJqr3AOrih_xYjOhDe6U2x7hjQTD1Of_-0zvgxL80wuMxxD2n6k88geRGyjKXwbEWXfAFAUMN"
	user := "oFqEd6HUL3u0grQ6Zsa27SJ2w558"
	tid := "vy-_mZpvXezlnGeT3Qipyz_zNdy56bdi0smvrUyvfJ4"
	url := wechat.UrlTemplateMessageSend
	result := &wechat.TemplateSendResp{}
	data := struct {
		Name wechat.TemplateDataItem `json:"name"`
	}{
		Name: wechat.TemplateDataItem{Value: "平安度疫", Color: "#ff0000"},
	}
	template := wechat.NewTemplate().
		SetToUser(user).
		SetTemplateId(tid).
		SetData(data)
	wechat.PushReceptor(wechat.PushTypeTemplate, token).
		SetTemplate(template).
		Push(url, result)
}
