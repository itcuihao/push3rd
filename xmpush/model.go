package xmpush

// PushRecept push 接收
type PushRecept struct {
	version   string
	uri       string
	isSandbox bool
	Message   *Message
	Search    *SearchRecept
}

// Message 消息
type Message struct {
	Payload               string            `json:"payload"`                 // 消息内容payload
	RestrictedPackageName string            `json:"restricted_package_name"` // 设置app的多包名packageNames（多包名发送广播消息）。p
	PassThrough           int64             `json:"pass_through"`            // 是否通过透传的方式送给app，1表示透传消息，0表示通知栏消息。
	Title                 string            `json:"title"`                   // 通知栏展示的通知的标题
	Description           string            `json:"description"`             // 通知栏展示的通知的描述
	NotifyType            int64             `json:"notify_type"`             // DEFAULT_ALL = -1; DEFAULT_SOUND  = 1;   // 使用默认提示音提示 DEFAULT_VIBRATE = 2;   // 使用默认震动提示 DEFAULT_LIGHTS = 4;    // 使用默认led灯光提示
	TimeToLive            int64             `json:"time_to_live"`            // 可选项。如果用户离线，设置消息在服务器保存的时间，单位：ms。服务器默认最长保留两周。
	TimeToSend            int64             `json:"time_to_send"`            // 可选项。定时发送消息。timeToSend是以毫秒为单位的时间戳。注：仅支持七天内的定时消息。
	NotifyID              int64             `json:"notify_id"`               // 可选项。默认情况下，通知栏只显示一条推送消息。如果通知栏要显示多条推送消息，需要针对不同的消息设置不同的notify_id（相同notify_id的通知栏消息会覆盖之前的）。
	Extra                 map[string]string `json:"extra"`                   // 可选项，对app提供一些扩展的功能，请参考2.2。除了这些扩展功能，开发者还可以定义一些key和value来控制客户端的行为。注：key和value的字符数不能超过1024，至多可以设置10个key-value键值对。

	RegistrationId string `json:"registration_id"` // 根据registration_id，发送消息到指定设备上。可以提供多个registration_id，发送给一组设备，不同的registration_id之间用“,”分割。参数仅适用于“/message/regid”HTTP API。
	Alias          string `json:"alias"`           // 根据alias，发送消息到指定的设备。可以提供多个alias，发送给一组设备，不同的alias之间用“,”分割。参数仅适用于“/message/alias”HTTP API。
	Topic          string `json:"topic"`           // 根据topic，发送消息给订阅了该topic的所有设备。注：不同提供多个topic。参数仅适用于“/message/topic”HTTP API。
	Topics         string `json:"topics"`          // topic列表，使用;$;分割。注: topics参数需要和topic_op参数配合使用。参数仅适用于“/message/multi_topic”HTTP API。
	TopicOp        string `json:"topic_op"`        // topic之间的操作关系。支持以下三种：UNION并集 INTERSECTION交集 EXCEPT差集
}

// GetVersion version
func (p PushRecept) GetVersion() string {
	return p.version
}

// GetURI uri
func (p PushRecept) GetURI() string {
	return p.uri
}

// SetURI uri
func (p *PushRecept) SetURI(u string) {
	p.uri = u
}

// IsSandbox is sandbox
func (p PushRecept) IsSandbox() bool {
	return p.isSandbox
}

type SearchRecept struct {
	MsgId  string `json:"msg_id,omitempty"`
	JobKey string `json:"job_key,omitempty"`
}

type Result struct {
	Result      string `json:"result"`
	MessageID   string `json:"trace_id"`
	Code        int64  `json:"code"`
	Description string `json:"description,omitempty"`
	Info        string `json:"info,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

type SendResult struct {
	Result
	Data struct {
		ID string `json:"id,omitempty"`
	} `json:"data,omitempty"`
}

type SingleStatusResult struct {
	Result
	Data struct {
		Data struct {
			CreateTime      string `json:"create_time"`
			CreateTimestamp int64  `json:"create_timestamp"`
			TimeToLive      string `json:"time_to_live"`
			ClickRate       string `json:"click_rate"`
			MsgType         string `json:"msg_type"`
			DeliveryRate    string `json:"delivery_rate"`
			Delivered       int32  `json:"delivered"`
			ID              string `json:"id"`
			Click           int32  `json:"click"`
			Resolved        int32  `json:"resolved"`
		} `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
