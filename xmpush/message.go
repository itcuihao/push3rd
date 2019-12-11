package xmpush

import "strings"

// SetPayload 消息的内容。（注意：需要对payload字符串做urlencode处理）
func (p *PushRecept) SetPayload(m string) *PushRecept {
	p.Message.Payload = m
	return p
}

// SetRestrictedPackageName App的包名。备注：V2版本支持一个包名，V3版本支持多包名（中间用逗号分割）。
func (p *PushRecept) SetRestrictedPackageName(m ...string) *PushRecept {
	p.Message.RestrictedPackageName = strings.Join(m, ",")
	return p
}

// SetPassThrough pass_through的值可以为：
// 0 表示通知栏消息
// 1 表示透传消息
func (p *PushRecept) SetPassThrough(m int64) *PushRecept {
	p.Message.PassThrough = m
	return p
}

// SetTitle 通知栏展示的通知的标题。
func (p *PushRecept) SetTitle(m string) *PushRecept {
	p.Message.Title = m
	return p
}

// SetDescription notify_type的值可以是DEFAULT_ALL或者以下其他几种的OR组合：
// -1 所有
// 1 使用默认提示音提示；
// 2 使用默认震动提示；
// 4 使用默认led灯光提示；
func (p *PushRecept) SetNotifyType(m int64) *PushRecept {
	switch m {
	case Notify_Type_All:
		fallthrough
	case Notify_Type_Sound:
		fallthrough
	case Notify_Type_Vibrate:
		fallthrough
	case Notify_Type_Lights:
		p.Message.NotifyType = m
	}
	return p
}

// SetDescription 通知栏展示的通知的描述。
func (p *PushRecept) SetDescription(m string) *PushRecept {
	p.Message.Description = m
	return p
}

// SetTimeToLive 可选项。如果用户离线，设置消息在服务器保存的时间，单位：ms。服务器默认最长保留两周。
func (p *PushRecept) SetTimeToLive(m int64) *PushRecept {
	p.Message.TimeToLive = m
	return p
}

func (p *PushRecept) SetTimeToSend(m int64) *PushRecept {
	p.Message.TimeToSend = m
	return p
}

func (p *PushRecept) SetNotifyId(m int64) *PushRecept {
	p.Message.NotifyID = m
	return p
}
func (p *PushRecept) SetExtra(key string, value string) *PushRecept {
	if p.Message.Extra == nil {
		p.Message.Extra = make(map[string]string)
	}
	p.Message.Extra[key] = value
	return p
}

func (p *PushRecept) SetAll() *PushRecept {
	p.SetURI(URIMessageAll)
	return p
}

func (p *PushRecept) SetRegistrationId(m ...string) *PushRecept {
	p.Message.RegistrationId = strings.Join(m, ",")
	p.SetURI(URIRegId)
	return p
}

func (p *PushRecept) SetAlias(m ...string) *PushRecept {
	p.Message.Alias = strings.Join(m, ",")
	p.SetURI(URIMessageAlisa)
	return p
}

func (p *PushRecept) SetTopic(m string) *PushRecept {
	p.Message.Topic = m
	return p
}

func (p *PushRecept) SetTopics(op string, m ...string) *PushRecept {
	p.Message.TopicOp = op
	p.Message.Topics = strings.Join(m, ",$,")
	return p
}

func (p *PushRecept) SetAccount() *PushRecept {
	p.SetURI(URIMessageUserAccount)
	return p
}
