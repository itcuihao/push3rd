package smss

import (
	"strconv"
	"testing"
)

func TestTx(t *testing.T) {
	appId := "123"
	appKey := "1231"
	RegistrySmsTx(appId, appKey)
	err := NewSmsTx().SetTplId(428386).
		SetTel("", "170999").
		SetParams("1234", strconv.Itoa(smsLoginDuration)).
		Send()
	t.Log(err)
}
