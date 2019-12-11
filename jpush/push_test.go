package jpush

import "testing"

func TestPush(t *testing.T) {
	s := "1"
	a := "2"
	na := &NotificationAndroid{
		Alert: "081501",
		Title: "hello world",
	}
	result, err := PushReceptor().
		SetPlatform(PlatformA).
		SetRegistrationId("100d8559093c67ceb36").
		SetNotifyAndroid(na).
		Push(s, a)
	t.Log(err)
	t.Log(result)
}
