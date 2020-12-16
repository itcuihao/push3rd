package wechat

import "net/url"

const (
	AccessToken = "access_token"
	PushTypeTemplate = "template"
)

// Push 微信推送结构
type PushRecept struct {
	Template *Template
	Type     string
	UrlQuery url.Values
}

// Template 模板消息
type Template struct {
	ToUser      string       `json:"touser"`                // 接收者openid
	TemplateId  string       `json:"template_id"`           // 模板ID
	Url         string       `json:"url,omitempty"`         // 板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        interface{}  `json:"data"`                  // 模板数据
	Color       string       `json:"color"`                 // 模板内容字体颜色，不填默认为黑色
}

// MiniProgram 跳小程序所需数据
type MiniProgram struct {
	AppId    string `json:appid`              // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:pagepath,omitempty` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}

type TemplateDataItem struct {
	Value string `json:"value"` // 模板值
	Color string `json:"color"` // 模板颜色
}

// TemplateSendResp 模板发送返回
type TemplateSendResp struct {
	Error
	MsgID int `json:"msgid"`
}

// 错误信息
type Error struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}
