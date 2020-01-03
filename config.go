package gopush

import (
	"encoding/json"
	"io/ioutil"
)

var conf *Config

type Config struct {
	Push *Push `json:"push"`
}

type Push struct {
	JPush  JPush  `json:"jpush"`
	HwPush HwPush `json:"hwpush"`
	MiPush MiPush `json:"mipush"`
}

type JPush struct {
	AppKey string `json:"app_key"`
	Secret string `json:"secret"`
}

type HwPush struct {
	AppId  string `json:"app_id"`
	Secret string `json:"secret"`
}

type MiPush struct {
	Secret string `json:"secret"`
}

func GetJPush() JPush {
	p := conf.Push
	if p == nil {
		panic("push not conf")
	}
	return p.JPush
}

func GetHwPush() HwPush {
	p := conf.Push
	if p == nil {
		panic("push not conf")
	}
	return p.HwPush
}

func GetMiPush() MiPush {
	p := conf.Push
	if p == nil {
		panic("push not conf")
	}
	return p.MiPush
}

func InitConfig(cfgFile string) error {
	var err error
	conf, err = readConf(cfgFile)
	if err != nil {
		return err
	}

	return nil
}

func readConf(cfgFile string) (*Config, error) {
	f, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	data := &Config{}
	err = json.Unmarshal(f, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
