package hwpush

import "testing"

func TestUrl(t *testing.T) {
	u := pushUrl("11234")
	t.Log(u)
}

func TestPush(t *testing.T) {
	aid := "1234"
	secret := "1234"
	p, err := NewPusherV2().
		SetDeviceList("0868793037585404300004109700CN01").
		SetMsgType(MsgTypeNotify).
		SetMsgTitle("hello").
		SetMsgContent("20200102").
		SetAction(MsgActionApp, "pkgname").
		Push(aid, secret)
	t.Log(p)
	t.Log(err)
}
