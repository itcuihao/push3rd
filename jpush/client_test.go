package jpush

import "testing"

func TestNewPush(t *testing.T) {
	s := "81fc7c99b4acc2affa653ec7"
	a := "af24c3ea89be9396b5a66262"
	pc := NewPushClient(s, a)
	t.Log(pc)
}
