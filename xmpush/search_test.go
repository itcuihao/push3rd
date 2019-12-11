package xmpush

import (
	"testing"
)

func TestSearch(t *testing.T) {
	r, err := NewPushRecptor(Version01, false).
		SetSearchMsgId("sdm02076565848800018uY").
		SearchStatus(secret)
	t.Log(err)
	t.Log(r)
}
