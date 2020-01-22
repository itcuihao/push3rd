package smss

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/itcuihao/gopush/client"
)

// 文档地址：网联短信
const (
	tpUrl            = "http://ip:prot/smss/smsTP.jsp"
	defaultTpTitle   = "【签名标题】"
	defaultTpTuiDing = "退订回T"
)

var smsTp SmsTpConf

type SmsTpConf struct {
	cpid     string
	password string
	appId    string
	appKey   string
}

func RegistrySmsTp(cpid, password, appId, appKey string) {
	smsTp = SmsTpConf{
		cpid:     cpid,
		password: password,
		appId:    appId,
		appKey:   appKey,
	}
}

// reqSmsTp 发送网联短信请求
type reqSmsTp struct {
	Mobile string   `json:"mobile"`           // 接收短信的手机号码，多个手机之间用英文逗号隔开，每批最多200个号码
	Tpid   string   `json:"tpid"`             // 短信模板id
	Sms    string   `json:"sms"`              // 短信内容
	Cpid   string   `json:"cpid"`             // 商户编号
	Datas  []string `json:"datas,omitempty"`  // 内容数据
	STime  string   `json:"s_time,omitempty"` // 定时发送时间
	Nonce  string   `json:"nonce"`            // 18位随机字符串
	Sign   string   `json:"sign"`             // 签名，md5(cpid+md5(password)+appId+appKey+nonce)，sign小写，md5(password)为大写
	Cuid   string   `json:"cuid,omitempty"`   // 商户订单号 可空
}

// respSmsTp 返回数据
type respSmsTp struct {
	RetCode string `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	OrderId string `json:"orderID"`
}

func NewSmsTp(nonce string) *reqSmsTp {
	return &reqSmsTp{
		Cpid:  smsTp.cpid,
		Nonce: nonce,
		Sign:  signSmsTp(smsTp.cpid, smsTp.password, smsTp.appId, smsTp.appKey, nonce),
	}
}

func signSmsTp(cpid, password, appID, appKey, nonce string) string {
	pswMD5 := md5.New()
	pswMD5.Write([]byte(password))
	tpsw := hex.EncodeToString(pswMD5.Sum(nil))

	str := cpid + strings.ToUpper(tpsw) + appID + appKey + nonce

	sign := md5.New()
	sign.Write([]byte(str))
	return hex.EncodeToString(sign.Sum(nil))
}

func (s *reqSmsTp) SetMobile(m ...string) *reqSmsTp {
	s.Mobile = strings.Join(m, ",")
	return s
}

func (s *reqSmsTp) SetTpId(t string) *reqSmsTp {
	s.Tpid = t
	return s
}

func (s *reqSmsTp) SetSms(title, sms string) *reqSmsTp {
	if !strings.HasSuffix(sms, defaultTpTuiDing) {
		sms += defaultTpTuiDing
	}

	if title == "" {
		sms = defaultTpTitle + sms
	} else {
		sms = fmt.Sprintf("【%s】%s", title, sms)
	}
	s.Sms = url.QueryEscape(sms)
	return s
}

func (s *reqSmsTp) SetData(data ...string) *reqSmsTp {
	s.Datas = data
	return s
}

func (s *reqSmsTp) SetSTime(st string) *reqSmsTp {
	s.STime = st
	return s
}

func (s *reqSmsTp) SetCuId(c string) *reqSmsTp {
	s.Cuid = c
	return s
}

func (s *reqSmsTp) Send() error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	req, err := http.NewRequest("POST", tpUrl, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", client.ContentTypeJSON)
	req.Header.Add("Content-Type", client.ContentTypeJSON+";charset=utf-8")

	resp, err := client.NewClient(&http.Client{}, 10*time.Second).Do(req)
	if err != nil {
		return err
	}

	tpRes := &respSmsTp{}
	if err := json.Unmarshal(resp, tpRes); err != nil {
		return err
	}
	fmt.Printf("%+v\n", tpRes)
	if tpRes.RetCode != "0" {
		return fmt.Errorf("Err: %s", tpRes.RetMsg)
	}
	return nil
}
