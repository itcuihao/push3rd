package smss

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/itcuihao/gopush/utils"
)

var (
	cpid     = "123"
	password = "132"
	appID    = "123"
	appKey   = "123"
)

func TestTp(t *testing.T) {
	RegistrySmsTp(cpid, password, appID, appKey)
	err := NewSmsTp(utils.GetRandStr(18)).
		SetMobile("170999").
		SetTpId("0019120001").
		SetSms("", "hello").
		Send()
	t.Log(err)
}

func TestTpEmoji(t *testing.T) {
	RegistrySmsTp(cpid, password, appID, appKey)
	err := NewSmsTp(utils.GetRandStr(18)).
		SetMobile("170999").
		// SetTpId("0019120001").
		SetSms("", unquoteCodePoint("\\U0001F467")).
		Send()
	t.Log(err)
}

func unquoteCodePoint(s string) string {
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	fmt.Println(err)
	return string(r)
}

func TestSign(t *testing.T) {
	nonce := "5ERBE9IT8ZWHCUE6ZI"
	sign := signSmsTp(cpid, password, appID, appKey, nonce)
	t.Log(sign)
	sl := strings.ToLower(sign)
	t.Log(sl == "e4c0b350b1ee5781ffd61aa5539d440b")
}
