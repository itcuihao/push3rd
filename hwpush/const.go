package hwpush

const (
	// url
	TokenUrl = "https://login.cloud.huawei.com/oauth2/v2/token"
	PushUrl  = "https://api.push.hicloud.com/pushsend.do"

	// config
	GrantType = "client_credentials"
	NspSvc    = "openpush.message.api.send"

	MsgTypeCustom = 1
	MsgTypeNotify = 3

	MsgActionCustom = 1
	MsgActionUrl    = 2
	MsgActionApp    = 3

	SuccessCode = "80000000"
)
