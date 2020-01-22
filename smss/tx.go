package smss

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/itcuihao/gopush/client"
)

// 文档地址：https://cloud.tencent.com/document/product/382/5976

const (
	defaultTxSign    = "签名标题"
	defaultTxNation  = "86"
	smsLoginDuration = 10
	txUrl            = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=%v&random=%v"
)

var smsTx SmsTxConf

type SmsTxConf struct {
	appId  string
	appKey string
}

func RegistrySmsTx(appId, appKey string) {
	smsTx = SmsTxConf{
		appId:  appId,
		appKey: appKey,
	}
}

type TxTel struct {
	Mobile     string `json:"mobile" form:"mobile"`
	NationCode string `json:"nation_code" form:"nation_code"`
}

type reqSmsTx struct {
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    TxTel    `json:"tel"`
	Time   int64    `json:"time"`
	TplId  int      `json:"tpl_id"`
}

type respSmsTx struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

func NewSmsTx() *reqSmsTx {
	return &reqSmsTx{
		Sign: defaultTxSign,
		Time: time.Now().Unix(),
	}
}

func (s *reqSmsTx) SetTel(nation, mobile string) *reqSmsTx {
	if nation == "" {
		nation = defaultTxNation
	}
	s.Tel = TxTel{
		NationCode: nation,
		Mobile:     mobile,
	}
	return s
}

func (s *reqSmsTx) SetTplId(tid int) *reqSmsTx {
	s.TplId = tid
	return s
}

func (s *reqSmsTx) SetParams(p ...string) *reqSmsTx {
	s.Params = append(s.Params, p...)
	return s
}

func (s *reqSmsTx) SetSign(sign string) *reqSmsTx {
	s.Sign = sign
	return s
}

func (s *reqSmsTx) setSig(tNow, random int64) *reqSmsTx {
	str := fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v", smsTx.appKey, random, tNow, s.Tel.Mobile)
	b := sha256.Sum256([]byte(str))
	s.Sig = hex.EncodeToString(b[:])
	return s
}

func (s *reqSmsTx) Send() error {
	now := time.Now()
	random := now.UnixNano()
	tNow := now.Unix()

	s.setSig(tNow, random)
	url := fmt.Sprintf(txUrl, smsTx.appId, random)

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", client.ContentTypeJSON)
	req.Header.Add("Content-Type", client.ContentTypeJSON+";charset=utf-8")

	resp, err := client.NewClient(&http.Client{}, 10*time.Second).Do(req)
	if err != nil {
		return err
	}

	txRes := &respSmsTx{}
	if err := json.Unmarshal(resp, txRes); err != nil {
		return err
	}

	if txRes.Result != 0 {
		return fmt.Errorf("Err: %v", txRes.ErrMsg)
	}
	return nil
}
