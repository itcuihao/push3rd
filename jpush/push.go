package jpush

import (
	"bytes"
	"net/http"
	"time"

	"github.com/itcuihao/gopush/client"
)

// PushReceptor 推送接收者
func PushReceptor() *PushRecept {
	return &PushRecept{}
}

// Push 推送
func (p *PushRecept) Push(secret, appKey string) (result string, err error) {

	pushClient := NewPushClient(secret, appKey)
	var pByte []byte
	pByte, err = FormatSendByte(p)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", pushClient.URL, bytes.NewReader(pByte))
	if err != nil {
		return
	}
	req.Header.Add("Charset", Charset)
	req.Header.Add("Authorization", pushClient.AuthB64)
	req.Header.Add("Content-Type", ContentTypeJSON)
	// client := NewClient(ContentTypeJSON, &http.Client{})
	// resp, err := client.DoPost(req)
	// if err != nil {
	// 	return
	// }

	client := client.NewClient(&http.Client{}, 10*time.Second)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	result = string(resp)
	return
}
