package wechat

const (
	UrlTemplateSetIndustry   = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry"         // 设置所属行业
	UrlTemplateGetIndustry   = "https://api.weixin.qq.com/cgi-bin/template/get_industry"             // 获取设置的行业信息
	UrlTemplateAdd           = "https://api.weixin.qq.com/cgi-bin/template/api_add_template"         // 获取模板ID
	UrlTemplateGetAllPrivate = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template" // 获取模板列表
	UrlTemplateDel           = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"     // 删除模板
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
