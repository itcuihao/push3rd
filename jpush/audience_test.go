package jpush

import "testing"

func TestTag(t *testing.T) {
	a := Audience{
		Tag:    []string{"A", "B"},
		TagAnd: []string{"a", "b"},
	}
	tag := a.getTag()
	t.Log(tag)
}
