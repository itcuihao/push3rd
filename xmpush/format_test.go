package xmpush

import (
	"fmt"
	"testing"
)

func TestMess(t *testing.T) {
	// vs := []string{Version02, Version03}
	vs := []string{Version02}
	for _, v := range vs {
		fmt.Println(v)
		p := NewPushRecptor(v, false).
			SetRestrictedPackageName("1234", "2345").
			SetRegistrationId("123")
		fmt.Printf("p:%+v\n", p)
		fmt.Printf("pm:%+v\n", p.Message)
		m := fmtMessage(*p.Message)
		fmt.Println(m)
	}
}

func TestUrlf(t *testing.T) {
	vs := []string{Version03}
	for _, v := range vs {
		fmt.Println(v)
		p := NewPushRecptor(v, false).
			SetRestrictedPackageName("1234", "2345").
			SetRegistrationId("123").
			SetExtra("Link", "1234")
		fmt.Printf("p:%+v\n", p)
		fmt.Printf("pm:%+v\n", p.Message)
		m := fmtUrlForm(*p.Message)
		fmt.Println(m)
	}
}
