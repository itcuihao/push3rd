package gopush

import (
	"encoding/json"
	"io/ioutil"
)

var conf *PushConfig

type PushConfig struct {
	JPush     *JPush     `json:"jpush"`
	HwPush    *HwPush    `json:"hwpush"`
	MiPush    *MiPush    `json:"mipush"`
	ApplePush *ApplePush `json:"apple_push"`
}

type JPush struct {
	AppKey string `json:"app_key"`
	Secret string `json:"secret"`
}

type HwPush struct {
	AppId       string `json:"app_id"`
	Secret      string `json:"secret"`
	PackageName string `json:"package_name"`
}

type MiPush struct {
	Secret      string `json:"secret"`
	PackageName string `json:"package_name"`
}

type ApplePush struct {
	PackageName string `json:"package_name"`
	P12Path     string `json:"p12_path"`
}

func GetJPush() JPush {
	if conf.JPush == nil {
		panic("push not conf")
	}
	return *conf.JPush
}

func GetHwPush() HwPush {
	if conf.HwPush == nil {
		panic("push not conf")
	}
	return *conf.HwPush
}

func GetMiPush() MiPush {
	if conf.MiPush == nil {
		panic("push not conf")
	}
	return *conf.MiPush
}

func GetApplePush() ApplePush {
	if conf.ApplePush == nil {
		panic("push not conf")
	}
	return *conf.ApplePush
}

func InitConfig(cfgFile string) error {
	var err error
	conf, err = readConf(cfgFile)
	if err != nil {
		return err
	}
	return nil
}

func readConf(cfgFile string) (*PushConfig, error) {
	f, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}
	data := &PushConfig{}
	err = json.Unmarshal(f, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
