package jpush

// Push 极光推送结构
type PushRecept struct {
	Platform     *Platform
	Audience     *Audience
	Notification *Notification
	Message      *Message
	Options      *Options
}

// Platform 推送平台设置
type Platform struct {
	Value interface{}
}

// Audience 推送设备指定
type Audience struct {
	IsAll          bool
	Tag            []string `audience:"tag"`
	TagAnd         []string `audience:"tag_and"`
	TagNot         []string `audience:"tag_not"`
	Alias          []string `audience:"tag_alias"`
	RegistrationId []string `audience:"registration_id"`
	Segment        []string `audience:"segment"`
	Abtest         []string `audience:"abtest"`
}

// Message 消息内容体。是被推送到客户端的内容。与 notification 一起二者必须有其一，可以二者并存
type Message struct {
	MsgContent  string                 `json:"msg_content,omitempty"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

// Options 推送参数
type Options struct {
	SendNo          int64  `json:"sendno,omitempty"`
	TimeToLive      int64  `json:"time_to_live,omitempty"`
	OverrideMsgId   int64  `json:"override_msg_id,omitempty"`
	ApnsProduction  *bool  `json:"apns_production,omitempty"`
	ApnsCollapseId  string `json:"apns_collapse_id,omitempty"`
	BigPushDuration int64  `json:"big_push_duration,omitempty"`
}

// Notification 通知内容体。是被推送到客户端的内容。与 message 一起二者必须有其一，可以二者并存
type Notification struct {
	Alert   string               `json:"alert,omitempty"`
	Android *NotificationAndroid `json:"android,omitempty"`
	Apple   *NotificationApple   `json:"ios,omitempty"`
}

// NotificationAndroid Android消息
type NotificationAndroid struct {
	Alert      string                 `json:"alert"`
	Title      string                 `json:"title,omitempty"`
	BuilderId  int64                  `json:"builder_id,omitempty"`
	ChannelId  string                 `json:"channel_id,omitempty"`
	Priority   int64                  `json:"priority,omitempty"`
	Category   string                 `json:"category,omitempty"`
	Style      int64                  `json:"style,omitempty"`
	AlertType  int64                  `json:"alert_type,omitempty"`
	BigText    string                 `json:"big_text,omitempty"`
	Inbox      map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath string                 `json:"big_pic_path,omitempty"`
	Extras     map[string]interface{} `json:"extras,omitempty"`
	LargeIcon  string                 `json:"large_icon,omitempty"`
	Intent     map[string]interface{} `json:"intent,omitempty"`
}

// NotificationApple iOS 消息
type NotificationApple struct {
	Alert            interface{}            `json:"alert"`
	Sound            interface{}            `json:"sound,omitempty"`
	Badge            string                 `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	ThreadId         string                 `json:"thread-id,omitempty"`
}

// PushSend 推送发送
type PushSend struct {
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Options      *Options    `json:"options,omitempty"`
}
