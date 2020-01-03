package hwpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/itcuihao/gopush/client"
)

func (p *PushReceptV2) Push(appId, secret string) (res string, err error) {
	var (
		token string
	)
	token, err = GetToken(appId, secret)
	if err != nil {
		return
	}
	p.AccessToken = token
	fmt.Printf("%+v\n", p.Payload)
	if len(p.DeviceTokenList) == 0 {
		err = fmt.Errorf("not found device")
		return
	}
	for _, devices := range toUidsList(p.DeviceTokenList...) {
		if er := p.SetDeviceList(devices...).send(appId); er != nil {
			err = er
		}
	}
	return "", nil
}

func pushUrl(appId string) string {
	nc := struct {
		Ver   string `json:"ver"`
		AppId string `json:"appId"`
	}{
		Ver:   "1",
		AppId: appId,
	}
	nb, _ := json.Marshal(nc)
	fmt.Println(string(nb))
	return PushUrl + "?nsp_ctx=" + url.QueryEscape(string(nb))
}

// func getDeviceList(deviceTokenList []string) [][]string {
// 	deviceList := make([][]string, 0, len(deviceTokenList)/100)
// 	deviceLen := len(deviceTokenList)
// 	if deviceLen > 100 {
// 		i := 100
// 		for i < deviceLen {
// 			deviceList = append(deviceList, deviceTokenList[:i])
// 			i += 100
// 		}
// 		deviceList = append(deviceList, deviceTokenList[deviceLen-i:])
// 	} else {
// 		deviceList = append(deviceList, deviceTokenList)
// 	}
// 	return deviceList
// }

func (p *PushReceptV2) send(appId string) error {
	var (
		req     *http.Request
		err     error
		resByte []byte
	)

	reqUrl := pushUrl(appId)
	device, _ := json.Marshal(p.DeviceTokenList)
	payload, _ := json.Marshal(p.Payload)

	param := make(url.Values)
	param.Add("access_token", p.AccessToken)
	param.Add("nsp_svc", p.NspSvc)
	param.Add("nsp_ts", p.NspTs)
	param.Add("device_token_list", string(device))
	param.Add("payload", string(payload))
	if p.ExpireTime != "" {
		param.Add("expire_time", p.ExpireTime)
	}

	req, err = http.NewRequest("POST", reqUrl, strings.NewReader(param.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", client.ContentTypeFORM)

	c := client.NewClient(&http.Client{}, 10*time.Second)
	resByte, err = c.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(string(resByte))
	res := &Res{}
	err = json.Unmarshal(resByte, res)
	if err != nil {
		return err
	}
	if res.Code != SuccessCode {
		return fmt.Errorf("Code: %s, Msg: %s, Rid: %s", res.Code, res.Msg, res.RequestId)
	}
	return nil
}

func toUidsList(uidStr ...string) [][]string {
	lenUid := len(uidStr)
	num := lenUid / 100
	if lenUid%100 > 0 {
		num += 1
	}

	uids := make([][]string, 0, num)
	i := 100
	for j := 0; j < lenUid; j += 100 {
		if i <= lenUid {
			uids = append(uids, uidStr[j:i])
			i += 100
		}
	}
	if lenUid%100 > 0 {
		uids = append(uids, uidStr[(num-1)*100:])
	}

	return uids
}
