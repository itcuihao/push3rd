package xmpush

import "testing"

var (
	packageName = "1"
	secret      = "2"
	rid         = "/5LjNgF5WrIeFtAFb05oIB5QFr/wNS3l08fmQC+WzgW51og2qWiArnCfCgLfkpfT"
	mrid        = "m49UwCPVY8h8TfLe6+pFsI+56MOmy+IXV9ow16nYWOneTgyAV1TX4bAes0hBPgHl"
)

func TestPush(t *testing.T) {
	r, err := NewPushRecptor(Version03, false).
		SetRestrictedPackageName(packageName).
		SetPassThrough(0).
		SetPayload("hellomao").
		SetTitle("081403").
		SetDescription("read").
		SetRegistrationId(rid).
		Push(secret)
	t.Log(err)
	t.Log(r)
}
