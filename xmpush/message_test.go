package xmpush

import (
	"fmt"
	"testing"
)

func TestPkgn(t *testing.T) {
	vs := []string{Version02, Version03}
	for _, v := range vs {
		fmt.Println(v)
		p := NewPushRecptor(v, false).
			SetRestrictedPackageName("1234", "2345")
		fmt.Printf("p:%+v\n", p)
		fmt.Printf("pm:%+v\n", p.Message)
	}
}
