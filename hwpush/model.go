package hwpush

// PushReceptV2 推送结构v2版本
// https://developer.huawei.com/consumer/cn/service/hms/catalog/huaweipush_agent.html?page=hmssdk_huaweipush_api_reference_agent_s2
type PushReceptV2 struct {
	AccessToken     string     `json:"access_token"`
	NspSvc          string     `json:"nsp_svc"`
	NspTs           string     `json:"nsp_ts"`
	DeviceTokenList []string   `json:"device_token_list"`
	ExpireTime      string     `json:"expire_time"`
	Payload         *PayloadV2 `json:"payload"`
}

type PayloadV2 struct {
	Hps *HpsV2 `json:"hps"`
}

type HpsV2 struct {
	Msg *MsgV2 `json:"msg"`
	Ext *ExtV2 `json:"ext"`
}

type MsgV2 struct {
	Type   int       `json:"type"`
	Body   *BodyV2   `json:"body"`
	Action *ActionV2 `json:"action"`
}

type BodyV2 struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ActionV2 struct {
	Type  int      `json:"type"`
	Param *ParamV2 `json:"param"`
}

type ParamV2 struct {
	Intent     string `json:"intent"`
	Url        string `json:"url"`
	AppPkgName string `json:"appPkgName"`
}

type ExtV2 struct {
	BiTag       string `json:"biTag"`
	Customize   string `json:"customize"`
	BadgeAddNum string `json:"badgeAddNum"`
	BadgeClass  string `json:"badgeClass"`
}

// PushReceptV3 推送结构v3版本
type PushReceptV3 struct {
	ValidateOnly bool     `json:"validate_only"`
	Message      *Message `json:"message"`
}

type PushResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type AndroidConfig struct {
	CollapseKey   int                  `json:"collapse_key"`
	Priority      string               `json:"priority"`
	TTL           string               `json:"ttl"`
	BiTag         string               `json:"bi_tag"`
	FastAppTarget int                  `json:"fast_app_target"`
	Notification  *AndroidNotification `json:"notification"`
}

type AndroidNotification struct {
	Title         string             `json:"title"`
	Body          string             `json:"body"`
	Icon          string             `json:"icon"`
	Color         string             `json:"color"`
	Sound         string             `json:"sound"`
	Tag           string             `json:"tag"`
	ClickAction   *ClickAction       `json:"click_action"`
	BodyLocKey    string             `json:"body_loc_key"`
	BodyLocArgs   []string           `json:"body_loc_args"`
	TitleLocKey   string             `json:"title_loc_key"`
	TitleLocArgs  []string           `json:"title_loc_args"`
	ChannelId     string             `json:"channel_id"`
	NotifySummary string             `json:"notify_summary"`
	NotifyIcon    string             `json:"notify_icon"`
	Style         int                `json:"style"`
	BigTitle      string             `json:"big_title"`
	BigBody       string             `json:"big_body"`
	BigPicture    string             `json:"big_picture"`
	AutoClear     int                `json:"auto_clear"`
	NotifyId      int                `json:"notify_id"`
	Group         string             `json:"group"`
	Badge         *BadgeNotification `json:"badge"`
}

type ClickAction struct {
	Type         int    `json:"type"`
	Intent       string `json:"intent"`
	Url          string `json:"url"`
	RichResource string `json:"rich_resource"`
}

type BadgeNotification struct {
	Num   int    `json:"num"`
	Class string `json:"class"`
}

type Message struct {
	Data         string         `json:"data"`
	Notification *Notification  `json:"notification"`
	Android      *AndroidConfig `json:"android"`
	Token        []string       `json:"token"`
	Topic        string         `json:"topic"`
	Condition    string         `json:"condition"`
}

// https://developer.huawei.com/consumer/cn/service/hms/catalog/huaweipush_agent.html?page=hmssdk_huaweipush_api_reference_agent_s1
type TokenRes struct {
	Access_token     string `json:"access_token"`
	Expires_in       int    `json:"expires_in"`
	Token_type       string `json:"token_type"`
	Error            int    `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type Res struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}
