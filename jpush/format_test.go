package jpush

import "testing"

func TestFmt(t *testing.T) {
	a := &Audience{
		Tag:    []string{"A", "B"},
		TagAnd: []string{"a", "b"},
	}
	p := &PushRecept{
		Platform: &Platform{"all"},
		Audience: a,
	}
	p.SetNotifyAll("hi")
	b, err := FormatSendByte(p)
	t.Log(err)
	t.Log(string(b))
}
