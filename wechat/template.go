package wechat

const (
	UrlTemplateMessageSend   = "https://api.weixin.qq.com/cgi-bin/message/template/send"             // 发送模板信息
)

func NewTemplate()*Template{
	return &Template{}
}

func (t *Template) checkTemplateNil() {
	if t == nil {
		t = &Template{}
	}
}

func (t *Template) SetToUser(openid string) *Template {
	t.checkTemplateNil()
	t.ToUser = openid
	return t
}

func (t *Template) SetTemplateId(tid string) *Template {
	t.checkTemplateNil()
	t.TemplateId = tid
	return t
}

func (t *Template) SetUrl(url string) *Template {
	t.checkTemplateNil()
	t.Url = url
	return t
}

func (t *Template) SetMiniProgram(appId, pagePath string) *Template {
	t.checkTemplateNil()
	if appId != "" {
		t.MiniProgram = &MiniProgram{
			AppId:    appId,
			PagePath: pagePath,
		}
	}
	return t
}

func (t *Template) SetData(data interface{}) *Template {
	t.checkTemplateNil()
	t.Data = data
	return t
}
