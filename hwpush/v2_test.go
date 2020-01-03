package hwpush

import "testing"

func TestDevice(t *testing.T) {
	p := PushReceptV2{}
	p.SetDeviceList("1")
	t.Log(p.DeviceTokenList)
}
