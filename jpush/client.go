package jpush

import (
	"encoding/base64"
)

type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthB64      string
	URL          string
}

func NewPushClient(secret, appKey string) *PushClient {
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(appKey+":"+secret))
	return &PushClient{secret, appKey, auth, PushURL}
}
