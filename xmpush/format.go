package xmpush

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func FormatPush(p PushRecept) (result string, err error) {
	if p.Message == nil {
		err = fmt.Errorf("message is empty")
		return
	}

	result = fmtUrlForm(*p.Message)
	return
}

func fmtUrlForm(m Message) string {
	form := url.Values{}
	rt := reflect.TypeOf(m)
	rv := reflect.ValueOf(m)

	for i := 0; i < rt.NumField(); i++ {
		key := rt.Field(i).Tag.Get("json")
		v := rv.Field(i).Interface()
		vtype := reflect.TypeOf(v)

		switch vtype.Kind() {
		case reflect.Map:
			valueM := v.(map[string]string)
			for k, value := range valueM {
				form.Add(key+"."+k, value)
			}
		case reflect.String:
			value := v.(string)
			if value != "" {
				form.Add(key, value)
			}
		case reflect.Int64:
			value := strconv.FormatInt(v.(int64), 10)
			if value != "0" {
				form.Add(key, value)
			}
			if key == "pass_through" {
				form.Add(key, value)
			}
		}
	}
	return form.Encode()
}

func fmtMessage(m Message) map[string]interface{} {
	result := make(map[string]interface{})
	rt := reflect.TypeOf(m)
	rv := reflect.ValueOf(m)
	for i := 0; i < rt.NumField(); i++ {
		switch rt.Kind() {
		case reflect.Slice, reflect.Array:
			n := rt.Field(i).Name
			v := rv.Field(i).Interface()
			result[n] = v
		case reflect.Struct:
			key := rt.Field(i).Tag.Get("json")
			v := rv.Field(i).Interface()
			vv, ok := v.([]string)
			if key != "" && ok && len(vv) > 0 {
				result[key] = vv
			}
		}
	}
	return result
}

func FormatSearch(p PushRecept) (result string, err error) {
	if p.Search == nil {
		err = fmt.Errorf("message is empty")
		return
	}
	search := p.Search
	form := url.Values{}
	if search.MsgId != "" {
		form.Add("msg_id", search.MsgId)
	}
	if search.JobKey != "" {
		form.Add("job_key", search.JobKey)
	}
	result = fmt.Sprintf("?%s", form.Encode())

	return
}
