package hwpush

import (
	"strconv"
	"time"
)

func NewPusherV2() *PushReceptV2 {
	return &PushReceptV2{
		NspSvc: NspSvc,
		NspTs:  strconv.Itoa(int(time.Now().Unix())),
	}
}

func (p *PushReceptV2) SetDeviceList(ds ...string) *PushReceptV2 {
	if len(ds) == 0 {
		return p
	}
	p.DeviceTokenList = ds
	return p
}

func (p *PushReceptV2) checkMsgNil() *PushReceptV2 {
	if p.Payload == nil {
		p.Payload = new(PayloadV2)
	}
	if p.Payload.Hps == nil {
		p.Payload.Hps = new(HpsV2)
	}
	if p.Payload.Hps.Msg == nil {
		p.Payload.Hps.Msg = new(MsgV2)
	}
	return p
}

// SetMsgType
// 取值含义和说明：
// 1 透传异步消息
// 3 系统通知栏异步消息
// 注意：2和4以后为保留后续扩展使用
func (p *PushReceptV2) SetMsgType(t int) *PushReceptV2 {
	p.checkMsgNil()

	switch t {
	case MsgTypeCustom:
		p.Payload.Hps.Msg.Type = MsgTypeCustom
	case MsgTypeNotify:
		p.Payload.Hps.Msg.Type = MsgTypeNotify
	}
	return p
}

func (p *PushReceptV2) SetMsgTitle(title string) *PushReceptV2 {
	p.checkMsgNil()
	if p.Payload.Hps.Msg.Body == nil {
		p.Payload.Hps.Msg.Body = new(BodyV2)
	}
	p.Payload.Hps.Msg.Body.Title = title
	return p
}

func (p *PushReceptV2) SetMsgContent(content string) *PushReceptV2 {
	p.checkMsgNil()

	if p.Payload.Hps.Msg.Body == nil {
		p.Payload.Hps.Msg.Body = new(BodyV2)
	}
	p.Payload.Hps.Msg.Body.Content = content
	return p
}

// SetActionType
// 1 自定义行为：行为由参数intent定义
// 2 打开URL：URL地址由参数url定义
// 3 打开APP：默认值，打开App的首页
// 注意：富媒体消息开放API不支持。
func (p *PushReceptV2) SetAction(ty int, param string) *PushReceptV2 {
	p.checkMsgNil()
	if p.Payload.Hps.Msg.Action == nil {
		p.Payload.Hps.Msg.Action = new(ActionV2)
	}
	switch ty {
	case MsgActionCustom:
		p.Payload.Hps.Msg.Action.Type = MsgTypeCustom
		p.Payload.Hps.Msg.Action.Param = &ParamV2{
			Intent: param,
		}
	case MsgActionUrl:
		p.Payload.Hps.Msg.Action.Type = MsgActionUrl
		p.Payload.Hps.Msg.Action.Param = &ParamV2{
			Url: param,
		}
	case MsgActionApp:
		p.Payload.Hps.Msg.Action.Type = MsgActionApp
		p.Payload.Hps.Msg.Action.Param = &ParamV2{
			AppPkgName: param,
		}
	default:
		p.Payload.Hps.Msg.Action.Type = MsgActionApp
		p.Payload.Hps.Msg.Action.Param = &ParamV2{
			AppPkgName: param,
		}
	}
	return p
}
