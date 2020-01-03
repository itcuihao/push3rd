package gopush

import (
	"sync"
	"testing"
)

func TestSimple(t *testing.T){

	var (
		aDevices []string // 安卓
		iDevices []string // 苹果
		hDevices []string // 华为
		mDevices []string // 小米
		wg       = &sync.WaitGroup{}

		title, alert string
	)

	wg.Add(4)
	go sendJAndroid(wg, title, alert, aDevices...)
	go sendJIos(wg, title, alert, true, iDevices...)
	go sendHuaWei(wg, title, alert, hDevices...)
	go sendXiaoMi(wg, title, alert, mDevices...)
	wg.Wait()
}
