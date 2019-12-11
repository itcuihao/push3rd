package jpush

func (p *PushRecept) SetMsgContent(m string) *PushRecept {
	if p.Message == nil {
		p.Message = new(Message)
	}
	p.Message.MsgContent = m
	return p
}
func (p *PushRecept) SetMsgTitle(m string) *PushRecept {
	if p.Message == nil {
		p.Message = new(Message)
	}
	p.Message.Title = m
	return p
}

func (p *PushRecept) SetMsgContentType(m string) *PushRecept {
	if p.Message == nil {
		p.Message = new(Message)
	}
	p.Message.ContentType = m
	return p
}

func (p *PushRecept) SetMsgExtras(key string, value interface{}) *PushRecept {
	if p.Message == nil {
		p.Message = new(Message)
	}
	if p.Message.Extras == nil {
		p.Message.Extras = make(map[string]interface{})
	}
	p.Message.Extras[key] = value
	return p
}
