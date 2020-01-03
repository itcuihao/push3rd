package jpush

import (
	"encoding/json"
	"fmt"
)

// FormatSendByte 格式化推送数据
func FormatSendByte(pr *PushRecept) (b []byte, err error) {
	if pr == nil {
		err = fmt.Errorf("push recept is empty")
		return
	}
	if pr.Platform == nil {
		err = fmt.Errorf("push platform is empty")
		return
	}
	pf := pr.Platform.Value

	if pr.Audience == nil {
		err = fmt.Errorf("push audience is empty")
		return
	}
	au := pr.Audience.formatAudience()

	no := pr.Notification
	me := pr.Message
	if no == nil && me == nil {
		err = fmt.Errorf("push notification and message both cannot be empty")
		return
	}

	ps := &PushSend{
		Platform: pf,
		Audience: au,
	}

	if no != nil {
		ps.Notification = no
		if no.Apple != nil {
			// 如果目标平台为 iOS 平台，推送 Notification 时需要在 options 中通过 apns_production 字段来设定推送环境。True 表示推送生产环境，False 表示要推送开发环境； 如果不指定则为推送生产环境；一次只能推送给一个环境。
			// pr.SetOptionApns(false)
		}
	}
	if me != nil {
		ps.Message = me
	}

	ps.Options = pr.Options

	return json.Marshal(ps)
}
