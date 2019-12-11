package jpush

import (
	"reflect"
)

func (p *PushRecept) SetAll() *PushRecept {
	if p.Audience == nil {
		p.Audience = new(Audience)
	}
	p.Audience.IsAll = true
	return p
}

func (p *PushRecept) SetTag(t ...string) *PushRecept {
	if p.Audience == nil {
		p.Audience = new(Audience)
	}
	p.Audience.Tag = t
	return p
}

func (p *PushRecept) SetTagAnd(t ...string) *PushRecept {
	if p.Audience == nil {
		p.Audience = new(Audience)
	}
	p.Audience.TagAnd = t
	return p
}

func (p *PushRecept) SetTagNot(t ...string) *PushRecept {
	if p.Audience == nil {
		p.Audience = new(Audience)
	}
	p.Audience.TagNot = t
	return p

}

func (p *PushRecept) SetAlias(t ...string) *PushRecept {
	if p.Audience == nil {
		p.Audience = new(Audience)
	}
	p.Audience.Alias = t
	return p
}

func (p *PushRecept) SetRegistrationId(t ...string) *PushRecept {
	if p.Audience == nil {
		p.Audience = new(Audience)
	}
	p.Audience.RegistrationId = t
	return p
}

func (a *Audience) formatAudience() (result interface{}) {
	if a.IsAll {
		result = "all"
		return
	}
	return a.getTag()
}

func (a Audience) getTag() map[string]interface{} {
	result := make(map[string]interface{})
	rt := reflect.TypeOf(a)
	rv := reflect.ValueOf(a)
	for i := 0; i < rt.NumField(); i++ {
		switch rt.Kind() {
		case reflect.Slice, reflect.Array:
			n := rt.Field(i).Name
			v := rv.Field(i).Interface()
			result[n] = v
		case reflect.Struct:
			key := rt.Field(i).Tag.Get("audience")
			v := rv.Field(i).Interface()
			vv, ok := v.([]string)
			if ok && len(vv) > 0 {
				result[key] = vv
			}
		}
	}
	return result
}
