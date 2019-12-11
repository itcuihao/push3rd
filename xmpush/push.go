package xmpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/itcuihao/gopush/client"
)

func NewPushRecptor(version string, sandbox bool) *PushRecept {
	return &PushRecept{
		version:   version,
		isSandbox: sandbox,
		Message: &Message{
			PassThrough: PassThrough0,
			NotifyType:  Notify_Type_All,
		},
	}
}

func (p PushRecept) Push(secret string) (result SendResult, err error) {
	url := p.getUrl()

	form, err := FormatPush(p)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(form))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", secret))
	req.Header.Add("Content-Type", ContentTypeFORM)

	// fmt.Printf("req:%+v\n", req)
	// rb, _ := ioutil.ReadAll(req.Body)
	// req.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
	// fmt.Printf("r body:%+v\n", string(rb))
	client := client.NewClient(&http.Client{}, 10*time.Second)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &result)
	fmt.Printf("result:%+v\n", result)
	return
}

func (p PushRecept) getUrl() string {
	if p.IsSandbox() {
		return SandboxURL + p.GetVersion() + p.GetURI()
	}

	return ProductURL + p.GetVersion() + p.GetURI()
}
