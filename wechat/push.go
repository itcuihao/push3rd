package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/itcuihao/gopush/client"
)

// PushReceptor 推送接收者
func PushReceptor(pt, token string) *PushRecept {
	params := url.Values{}
	params.Set(AccessToken, token)

	return &PushRecept{
		Type:     pt,
		UrlQuery: params,
	}
}

func (p *PushRecept) SetTemplate(t *Template) *PushRecept {
	p.Template = t
	return p
}

// Push 推送
func (p *PushRecept) Push(url string, result interface{}) (err error) {
	var data string
	switch p.Type {
	case PushTypeTemplate:
		dbyte, derr := json.Marshal(p.Template)
		if derr != nil {
			err = derr
			return
		}
		data = string(dbyte)
	default:
		err = fmt.Errorf("not found push type")
		return
	}

	req, err := http.NewRequest(http.MethodPost, p.url(url), strings.NewReader(data))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", client.ContentTypeJSON)

	cli := client.NewClient(&http.Client{}, 20*time.Second)
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(resp), &result)

	return
}

func (p *PushRecept) url(target string) string {
	return target + "?" + p.UrlQuery.Encode()
}
