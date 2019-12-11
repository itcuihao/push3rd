package jpush

func (p *PushRecept) SetNotifyAll(n string) *PushRecept {
	if p.Notification == nil {
		p.Notification = new(Notification)
	}
	p.Notification.Alert = n
	return p
}

func (p *PushRecept) SetNotifyAndroid(n *NotificationAndroid) *PushRecept {
	if p.Notification == nil {
		p.Notification = new(Notification)
	}
	p.Notification.Android = n
	return p
}

func (p *PushRecept) SetNotifyApple(n *NotificationApple) *PushRecept {
	if p.Notification == nil {
		p.Notification = new(Notification)
	}
	p.Notification.Apple = n
	return p
}
