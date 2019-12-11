package xmpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"tools/push/client"
)

func (p *PushRecept) SetSearchMsgId(m string) *PushRecept {
	if p.Search == nil {
		p.Search = &SearchRecept{}
	}
	p.Search.MsgId = m

	return p
}

func (p *PushRecept) SetSearchJobKey(m string) *PushRecept {
	if p.Search == nil {
		p.Search = &SearchRecept{}
	}
	p.Search.JobKey = m
	return p
}

func (p PushRecept) SearchStatus(secret string) (result SingleStatusResult, err error) {
	p.SetURI(MessageStatusURL)
	url := p.getUrl()
	var searchStr string
	searchStr, err = FormatSearch(p)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url+searchStr, nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", secret))
	req.Header.Add("Content-Type", ContentTypeFORM)
	fmt.Printf("req:%+v\n", req)

	client := client.NewClient(&http.Client{}, 10*time.Second)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &result)
	fmt.Printf("r:%+v\n", result)
	return
}
