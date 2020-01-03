package hwpush

import "testing"

func TestGetToken(t *testing.T) {
	aid := "1234"
	secret := "1234"
	a, err := GetToken(aid, secret)
	t.Log(a)
	t.Log(err)
}
