package jpush

func (p *PushRecept) SetOptionSendNo(o int64) *PushRecept {
	if p.Options == nil {
		p.Options = new(Options)
	}
	p.Options.SendNo = o
	return p
}

func (p *PushRecept) SetOptionTimelive(o int64) *PushRecept {
	if p.Options == nil {
		p.Options = new(Options)
	}
	p.Options.TimeToLive = o
	return p
}

func (p *PushRecept) SetOptionOverrideMsgId(o int64) *PushRecept {
	if p.Options == nil {
		p.Options = new(Options)
	}
	p.Options.OverrideMsgId = o
	return p
}

func (p *PushRecept) SetOptionApns(o bool) *PushRecept {
	if p.Options == nil {
		p.Options = new(Options)
	}
	p.Options.ApnsProduction = &o
	return p
}

func (p *PushRecept) SetOptionApnsCollapseId(o string) *PushRecept {
	if p.Options == nil {
		p.Options = new(Options)
	}
	p.Options.ApnsCollapseId = o
	return p
}

func (p *PushRecept) SetOptionBigPushDuration(o int64) *PushRecept {
	if p.Options == nil {
		p.Options = new(Options)
	}
	p.Options.BigPushDuration = o
	return p
}
