package apple

import "testing"

func TestPush(t *testing.T) {
	if err := InitApns("push_apple.p12"); err != nil {
		t.Log(err)
		return
	}
	topic := "package.name"
	apns := NewApns(topic, false)
	title := "1"
	body := "1"
	sound := ""
	link := ""
	payload := Payload{
		Aps:  Aps{Alert: Alert{Title: title, Body: body}, Badge: 1, Sound: sound},
		Link: link,
	}
	t.Logf("%+v", payload)
	resp, err := apns.QuickPush("f9a4040ecd89b0f08511fd479de65ba016f0ac15b25ae57d585b972a49a8e242", payload)
	t.Log(err)
	t.Logf("%+v", resp)
}

func TestMassPush(t *testing.T) {
	if err := InitApns("push_apple.p12"); err != nil {
		t.Log(err)
		return
	}
	topic := "package.name"
	apns := NewApns(topic, false)

	title := "1"
	body := "1"
	sound := ""
	link := ""
	payload := Payload{
		Aps:  Aps{Alert: Alert{Title: title, Body: body}, Badge: 1, Sound: sound},
		Link: link,
	}
	t.Logf("%+v", payload)
	devices := make([]string, 0)
	for i := 0; i < 1; i++ {
		devices = append(devices, "3442dbd56caca5d53f8ae265941f33adc3c2c10b839d6be14dbbd59e92add96c")
	}
	err := apns.SendMassDevice(payload, devices...)
	t.Log(err)
}
